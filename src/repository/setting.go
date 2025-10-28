package repository

import (
	"fmt"
	"gbase/src/core/logger"
	"gbase/src/http/resource"
	"gbase/src/model"
	"strings"

	"gorm.io/gorm"
)

var _ IRepository = (*SettingRepository)(nil)

type SettingRepository struct {
	model *model.Setting
}

func NewSettingRepository() *SettingRepository {
	return &SettingRepository{model: &model.Setting{}}
}

func (r SettingRepository) NewInput() interface{} {
	return model.SettingInputs{}
}

func (r SettingRepository) NewModel() model.IModel {
	return r.model.NewModel()
}

func (r SettingRepository) FindAll(search map[string]string) (interface{}, resource.Pagination, error) {
	model := r.NewModel()
	db := model.DB().Model(model)
	out := model.NewModels()

	for key, val := range search {
		if strings.TrimSpace(val) != "" {
			db.Where(fmt.Sprintf("`%s` = ?", key), val)
		}
	}

	err := db.Order("`group`, `order`").Find(&out).Error

	return out, resource.Pagination{}, err
}

func (r SettingRepository) Find(key any) (interface{}, error) {
	return nil, nil
}

func (r SettingRepository) Create(data model.IModel) (model.IModel, error) {
	return nil, nil
}

func (r SettingRepository) UpdateSettings(inputs model.SettingInputs, group string) error {
	db := r.NewModel().DB().Session(&gorm.Session{})

	for _, input := range inputs {
		existing := r.NewModel().(*model.Setting)
		err := db.Where("`group` = ? AND `id` = ?", group, input.Id).First(existing).Error
		if err != nil {
			// 查無資料
			continue
		}

		// 更新
		existing.Value = input.Value

		// Updates不會自動更新時戳，需手動設置
		existing.SetUpdatedAt()

		// 更新資料
		err = db.Model(existing).Updates(existing).Error
		if err != nil {
			logger.Log.Printf("編輯錯誤: %v", err)
			return err
		}
	}

	return nil
}

func (r SettingRepository) Update(data model.IModel, key any) (model.IModel, error) {
	return nil, nil
}

func (r SettingRepository) Delete(key any) error {
	return nil
}

func (r SettingRepository) UpdateOrder(id1, id2 any) error {
	return nil
}
