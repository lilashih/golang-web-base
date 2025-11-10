package db

import (
	"fmt"
	"gbase/src/core/config"
	"gbase/src/core/helper"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // 主連線

func init() {
	basePath, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("讀取專案根目錄失敗: %v", err))
	}

	// -------- 1. 連到主資料庫 --------
	mainPath := filepath.Join(basePath, config.DB.Main.Path)
	DB, err = gorm.Open(sqlite.Open(mainPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("連線 db.sqlite 失敗: %v", err))
	}

	// -------- 2. 把第2個db掛進來 --------
	secondPath := filepath.Join(basePath, config.DB.Second.Path)
	if err := DB.Exec("ATTACH DATABASE ? AS second", secondPath).Error; err != nil {
		panic(fmt.Errorf("ATTACH %s 失敗: %v", secondPath, err))
	}

	// -------- 3. 依需求開啟 Debug --------
	if helper.IsDebug() {
		DB = DB.Debug()
	}
}

// 關閉 Debug 日誌 (保持與舊 API 兼容)
func DisableDebug() {
	DB = DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
}
