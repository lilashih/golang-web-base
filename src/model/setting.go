package model

import (
	"gbase/src/core/db"

	"gorm.io/gorm"
)

var _ IModel = (*Setting)(nil)

type SettingInput struct {
	Id    string `json:"id" example:"ip" validate:"required"` // 設定ID
	Value string `json:"value" example:"192.1.1.1"`           // 設定值
}

type SettingInputs []SettingInput //@name SettingInputs

type Setting struct {
	Model
	Id         string `json:"id"`
	Value      string `json:"value"`                                           // 設定值
	Name       string `json:"name"`                                            // 名稱
	Type       string `json:"type"`                                            // 表單類型
	Group      string `json:"group"`                                           // 群組
	Option     any    `gorm:"null;default:null;serializer:json" json:"option"` // 表單選項，為陣列或object
	IsViewable bool   `gorm:"column:isViewable" json:"isViewable"`
	IsEditable bool   `gorm:"column:isEditable" json:"isEditable"`
} //@name Setting

func (Setting) NewModel() IModel {
	return &Setting{}
}

func (Setting) NewModels() interface{} {
	return []Setting{}
}

func (m *Setting) DB() *gorm.DB {
	return db.DB
}

func (m *Setting) SetKey(id any) {
	if val, ok := id.(string); ok {
		m.Id = val
	}
}
