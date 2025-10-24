package model

import "gorm.io/gorm"

type ISoftDelete interface {
	IsTrashed() bool
	WithoutTrashed(*gorm.DB) *gorm.DB
	OnlyTrashed(*gorm.DB) *gorm.DB
	QueryTrashed(isTrashed bool) func(*gorm.DB) *gorm.DB
}
