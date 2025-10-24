package model

import (
	"gbase/src/core/db"
	"gbase/src/core/helper"

	"gorm.io/gorm"
)

var _ IModel = (*Permission)(nil)

type Menu struct {
	Group2     string       `json:"group" example:"菜單"` // 菜單主群組名稱
	GroupOrder int          `json:"order" example:"1"`  // 菜單主群組排序
	Children   []Permission `json:"children"`           // 子群組項目
} //@name Menu

type Permission struct {
	Id          string `gorm:"primaryKey" json:"id"`
	GroupOrder  int    `gorm:"column:groupOrder" json:"-"` // 主排序
	Order       int    `json:"order"`                      // 子群組排序
	System      string `json:"system" example:"app"`       // 權限所屬系統
	Group1      string `json:"-" example:"菜單"`             // 主群組名稱
	Group2      string `json:"-" example:"設定"`             // 子群組名稱
	Name        string `json:"name" example:"基本資料"`        // 名稱
	Description string `json:"description"`
	Path        string `json:"path"` // API路徑
	IsActive    bool   `gorm:"column:isActive" json:"-"`
	CreatedAt   string `gorm:"column:createdAt" json:"-"`
	UpdatedAt   string `gorm:"column:updatedAt" json:"-"`
} //@name Permission

func (Permission) NewModel() IModel {
	return &Permission{}
}

func (Permission) NewModels() interface{} {
	return []Permission{}
}

func (m *Permission) DB() *gorm.DB {
	return db.DB
}

func (m *Permission) GetKeyName() string {
	return "id"
}

func (m *Permission) GetKey() any {
	return m.Id
}

func (m *Permission) SetKey(id any) {
	if val, ok := id.(string); ok {
		m.Id = val
	}
}

func (m *Permission) SetUpdatedAt() {
	m.UpdatedAt = helper.GetNow()
}
