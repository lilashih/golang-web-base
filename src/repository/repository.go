package repository

import (
	"fmt"
	"gbase/src/core/helper"
	"gbase/src/core/logger"
	core "gbase/src/core/model"
	"gbase/src/def"
	"gbase/src/http/resource"
	m "gbase/src/model"
	"math"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IRepository interface {
	NewModel() m.IModel
	NewInput() interface{}
	Find(id any) (interface{}, error)
	FindAll(map[string]string) (interface{}, resource.Pagination, error)
	Create(m.IModel) (m.IModel, error)
	Update(m.IModel, any) (m.IModel, error)
	Delete(any) error
	UpdateOrder(any, any) error
}

func FindAllHelper[T m.IModel](model T, search map[string]string) (interface{}, resource.Pagination, error) {
	db := model.DB().Model(model).Preload(clause.Associations)
	out := model.NewModels()

	if s, ok := any(model).(core.ISoftDelete); ok {
		db = db.Scopes(s.WithoutTrashed)
	}

	// 解析分頁
	perPage := 100
	page := 1
	if val, ok := search["perPage"]; ok && val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			perPage = i
		}
		delete(search, "perPage")
	}
	if val, ok := search["page"]; ok && val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			page = i
		}
		delete(search, "page")
	}
	offset := (page - 1) * perPage

	// 模糊查詢 (用 OR 包起來)
	orConditions := []string{}
	orValues := []interface{}{}
	for key, val := range search {
		if strings.TrimSpace(val) != "" {
			orConditions = append(orConditions, fmt.Sprintf("%s LIKE ?", key))
			orValues = append(orValues, "%"+val+"%")
		}
	}
	if len(orConditions) > 0 {
		db.Where(strings.Join(orConditions, " OR "), orValues...)
	}

	var total int64
	db.Count(&total)

	pagination := resource.Pagination{
		Total:       int(total),
		PerPage:     perPage,
		CurrentPage: page,
		LastPage:    int(math.Ceil(float64(total) / float64(perPage))),
	}

	if _, ok := any(model).(m.IOrder); ok {
		db.Order(fmt.Sprintf("`order`, `%s` DESC", model.GetKeyName()))
	} else {
		db.Order(fmt.Sprintf("`%s` DESC", model.GetKeyName()))
	}

	err := db.Offset(offset).Limit(perPage).Find(&out).Error

	return out, pagination, err
}

func FindHelper[T m.IModel](model T, key any, out interface{}) error {
	db := model.DB()

	if s, ok := any(model).(core.ISoftDelete); ok {
		db = db.Scopes(s.WithoutTrashed)
	}

	return db.First(out, fmt.Sprintf("%s = ?", model.GetKeyName()), key).Preload(clause.Associations).Error
}

func CreateHelper[T m.IModel](model T) (T, error) {
	err := model.DB().Create(model).Error
	if err != nil {
		logger.Log.Printf("新增錯誤: %v , data: %s", err, model)
		// 新增失敗: UNIQUE constraint failed: drugs.code ?????

		return model, err
	}

	// 更新 order
	if _, ok := any(model).(m.IOrder); ok {
		err := model.DB().Model(model).Updates(map[string]interface{}{
			"order": model.GetKey(),
		}).Error

		if err != nil {
			logger.Log.Printf("新增錯誤: 更新 order 失敗: %v , %s", err, model)
			return model, err
		}
	}

	// 以主鍵欄位查詢，避免非id時語法錯誤
	err = model.DB().Preload(clause.Associations).First(model, fmt.Sprintf("%s = ?", model.GetKeyName()), model.GetKey()).Error

	return model, err
}

func UpdateHelper[T m.IModel](model T) (m.IModel, error) {
	db := model.DB().Session(&gorm.Session{})

	existing := model.NewModel()
	err := db.First(existing, fmt.Sprintf("%s = ?", model.GetKeyName()), model.GetKey()).Error
	if err != nil {
		return nil, err
	}

	// 檢查軟刪除
	if s, ok := any(existing).(core.ISoftDelete); ok {
		if s.IsTrashed() {
			return nil, def.ErrRecordIsTrashed
		}
	}

	// Updates不會自動更新時戳，需手動設置
	model.SetUpdatedAt()

	// 更新資料
	err = db.Model(existing).Updates(model).Error
	if err != nil {
		logger.Log.Printf("編輯錯誤: %v", err)
		return nil, err
	}

	err = existing.DB().Preload(clause.Associations).First(existing, fmt.Sprintf("%s = ?", model.GetKeyName()), existing.GetKey()).Error
	return existing, err
}

func DeleteHelper[T m.IModel](model T, key any) error {
	db := model.DB()

	// 讀取資料
	err := db.First(model, fmt.Sprintf("%s = ?", model.GetKeyName()), key).Error
	if err != nil {
		return err
	}

	// 軟刪
	if s, ok := any(model).(core.ISoftDelete); ok {
		if s.IsTrashed() {
			return def.ErrRecordIsTrashed // 已經是刪除狀態就略過
		}

		return db.Model(model).Updates(map[string]interface{}{
			"isDeleted": true,
			"updatedAt": helper.GetNow(),
		}).Error
	}

	// 非軟刪
	return db.Delete(model).Error
}

func UpdateOrderHelper[T m.IModel](model T, id1, id2 any) error {
	db := model.DB()
	item1 := model.NewModel()
	item2 := model.NewModel()

	// 撈出兩筆未被軟刪除的資料
	if s, ok := any(model).(core.ISoftDelete); ok {
		db = db.Scopes(s.WithoutTrashed).Session(&gorm.Session{})
	}

	if err := db.First(&item1, fmt.Sprintf("%s = ?", model.GetKeyName()), id1).Error; err != nil {
		logger.Log.Printf("編輯排序錯誤: 找不到項目 id1: %v: %v", id1, err)
		return err
	}

	if err := db.First(&item2, fmt.Sprintf("%s = ?", model.GetKeyName()), id2).Error; err != nil {
		logger.Log.Printf("編輯排序錯誤: 找不到項目 id2: %v: %v", id2, err)
		return err
	}

	// 取得目前順序
	order1 := 0
	order2 := 0
	if s, ok := any(item1).(m.IOrder); ok {
		order1 = s.GetOrder()
	}
	if s, ok := any(item2).(m.IOrder); ok {
		order2 = s.GetOrder()
	}

	// 開始交換 order
	return db.Transaction(func(tx *gorm.DB) error {
		if err := db.Model(item1).Update("order", order2).Error; err != nil {
			return err
		}
		if err := db.Model(item2).Update("order", order1).Error; err != nil {
			return err
		}
		return nil
	})
}
