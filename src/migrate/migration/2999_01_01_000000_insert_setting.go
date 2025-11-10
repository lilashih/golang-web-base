package migration

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"

	"gbase/src/core/helper"
	"gbase/src/core/migrate"
	core "gbase/src/core/model"
	"gbase/src/model"
)

var InsertSetting = migrate.MigrateSql{
	Name: helper.GetCurrentFileNameNoExt(),
	Up: func(db *gorm.DB) error {
		items := []model.Setting{
			newSetting("name", "system", "系統名稱", "text", 1, nil, ""),
			newSetting("ip", "system", "IP", "text", 5, nil, ""),
			newSetting("color", "appearance", "主題", "select", 6, map[string]string{
				"green": "綠色", "blue": "藍色", "red": "紅色", "yellow": "黃色", "purple": "紫色", "orange": "橙色",
			}, ""),
		}

		if err := seedSetting(db, items); err != nil {
			return err
		}

		return nil
	},
}

func newSetting(id, group, name, typ string, order int, option any, value any) model.Setting {
	return model.Setting{
		Model: model.Model{
			Order: order,
		},
		Id:         id,
		Name:       name,
		Group:      group,
		Type:       typ,
		Option:     option,
		Value:      core.JsonString{Raw: value},
		IsEditable: true,
		IsViewable: true,
	}
}

func seedSetting(db *gorm.DB, items []model.Setting) error {
	for _, item := range items {
		// INSERT OR IGNORE
		if err := db.Exec(`INSERT OR IGNORE INTO settings(id, value) VALUES (?, ?)`, item.Id, item.Value).Error; err != nil {
			return fmt.Errorf("insert error on id %s: %w", item.Id, err)
		}

		// option 轉成 JSON 字串（nullable）
		var optionJSON *string
		if item.Option != nil {
			b, _ := json.Marshal(item.Option)
			s := string(b)
			optionJSON = &s
		}

		// UPDATE 欄位
		if err := db.Exec(`
			UPDATE settings
			SET "order" = ?, "type" = ?, "isViewable" = ?, "isEditable" = ?, "group" = ?, "name" = ?, "option" = ?
			WHERE id = ?`,
			item.Order, item.Type, item.IsViewable, item.IsEditable, item.Group, item.Name, optionJSON, item.Id,
		).Error; err != nil {
			return fmt.Errorf("update error on id %s: %w", item.Id, err)
		}
	}
	return nil
}
