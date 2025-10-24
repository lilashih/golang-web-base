package model

import (
	"gbase/src/core/db"
	"gbase/src/core/helper"
	core "gbase/src/core/model"
	"strconv"

	"gorm.io/gorm"
)

/**
 * Model
 */

var _ core.ISoftDelete = (*SoftDeleteModel)(nil)

type Model struct {
	TimestampModel
	Id    int `gorm:"primaryKey" json:"id"` // ID
	Order int `json:"order"`                // 排序
}

func (m *Model) DB() *gorm.DB {
	return db.DB
}

func (m *Model) GetKeyName() string {
	return "id"
}

func (m *Model) GetKey() any {
	return m.Id
}

func (m *Model) SetKey(id any) {
	switch val := id.(type) {
	case int:
		m.Id = val
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			m.Id = i
		}
	}
}

func (m *Model) GetOrder() int {
	return m.Order
}

/**
 * 時間戳
 */

type TimestampModel struct {
	CreatedAt string `gorm:"column:createdAt" json:"createdAt"` // 新增時間
	UpdatedAt string `gorm:"column:updatedAt" json:"updatedAt"` // 編輯時間
}

func (m *TimestampModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := helper.GetNow()
	m.CreatedAt = now
	m.UpdatedAt = now
	return
}

func (m *TimestampModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = helper.GetNow()
	return
}

func (m *TimestampModel) BeforeSave(tx *gorm.DB) (err error) {
	m.UpdatedAt = helper.GetNow()
	return
}

// BeforeUpdate、BeforeSave無法觸發時直接呼叫此方法
func (m *TimestampModel) SetUpdatedAt() {
	m.UpdatedAt = helper.GetNow()
}

/**
 * 軟刪
 */

type SoftDeleteModel struct {
	Model
	IsDeleted bool `gorm:"column:isDeleted" json:"isDeleted" example:"false"` // 是否被軟刪，0：未刪除、1：已刪除
}

func (m *SoftDeleteModel) IsTrashed() bool {
	return m.IsDeleted
}

func (m *SoftDeleteModel) WithoutTrashed(db *gorm.DB) *gorm.DB {
	return db.Where("isDeleted = ?", false)
}

func (m *SoftDeleteModel) OnlyTrashed(db *gorm.DB) *gorm.DB {
	return db.Where("isDeleted = ?", true)
}

func (m *SoftDeleteModel) QueryTrashed(isTrashed bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("isDeleted = ?", isTrashed)
	}
}

/**
 * interface
 */

type IModel interface {
	DB() *gorm.DB
	NewModel() IModel
	NewModels() interface{}
	GetKeyName() string
	GetKey() any
	SetKey(id any)
	SetUpdatedAt()
}

type IOrder interface {
	GetOrder() int
}

/**
 * 排序表單
 */

type OrderInput struct {
	Id1 any `json:"id1" validate:"required" example:"1" swaggertype:"string"` // 第一筆資料id
	Id2 any `json:"id2" validate:"required" example:"2" swaggertype:"string"` // 第二筆資料id
} //@name OrderInput
