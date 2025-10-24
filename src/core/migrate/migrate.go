package migrate

import (
	"gbase/src/core/logger"

	"gorm.io/gorm"
)

type Migration struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type MigrateSql struct {
	Name string
	Up   func(*gorm.DB) error
}

func Run(db *gorm.DB, list []MigrateSql) {
	logger.Migrate.Println("=========== 開始執行 Migration ===========")

	// 建立 migration 記錄表
	if err := db.AutoMigrate(&Migration{}); err != nil {
		logger.Migrate.Fatalf("建立 migration 記錄表失敗: %v", err)
	}

	var execCount int64
	for _, m := range list {
		var count int64
		db.Model(&Migration{}).Where("name = ?", m.Name).Count(&count)
		if count > 0 {
			continue
		}

		if err := m.Up(db); err != nil {
			logger.Migrate.Fatalf("✗ 執行失敗: %s : %v", m.Name, err)
		}

		logger.Migrate.Printf("執行成功: %s ", m.Name)
		db.Create(&Migration{Name: m.Name})
		execCount++
	}

	if execCount == 0 {
		logger.Migrate.Println("沒有需要執行的 migration")
	}

	logger.Migrate.Println("=========== 結束執行 Migration ===========")
}
