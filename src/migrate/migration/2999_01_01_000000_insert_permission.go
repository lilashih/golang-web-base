package migration

import (
	"fmt"

	"gorm.io/gorm"

	"gbase/src/core/helper"
	"gbase/src/core/migrate"
	"gbase/src/model"
)

var InsertPermission = migrate.MigrateSql{
	Name: helper.GetCurrentFileNameNoExt(),
	Up: func(db *gorm.DB) error {
		g2Order := map[string]int{
			"設定":   0,
			"項目設定": 1,
		}

		items := []model.Permission{
			newPermission("settingSystem", "app", "菜單", "設定", 1, "系統設定", "settings/groups/system", "", true),
			newPermission("settingAppearance", "app", "菜單", "設定", 2, "外觀設定", "settings/groups/appearance", "", true),

			newPermission("user", "app", "菜單", "項目設定", 1, "使用者", "users", "", true),
		}

		if err := seedPermission(db, items, g2Order); err != nil {
			return err
		}

		return nil
	},
}

func newPermission(id string, system string, group1 string, group2 string, order int, name string, path string, description string, isActive bool) model.Permission {
	return model.Permission{
		Id:          id,
		Order:       order,
		System:      system,
		Group1:      group1,
		Group2:      group2,
		Name:        name,
		Path:        path,
		Description: description,
		IsActive:    isActive,
	}
}

func seedPermission(db *gorm.DB, items []model.Permission, g2Order map[string]int) error {
	for i, item := range items {
		// INSERT OR IGNORE
		if err := db.Exec(`INSERT OR IGNORE INTO permissions(id) VALUES (?)`, item.Id).Error; err != nil {
			return fmt.Errorf("insert error on id %s: %w", item.Id, err)
		}

		if order, ok := g2Order[items[i].Group2]; ok {
			items[i].GroupOrder = order
		}

		// UPDATE 欄位
		if err := db.Exec(`
			UPDATE permissions
			SET groupOrder = ?, "order" = ?, system = ?, group1 = ?, group2 = ?, name = ?, description = ?, path = ?, IsActive = ?
			WHERE id = ?`,
			items[i].GroupOrder, item.Order, item.System, item.Group1, item.Group2, item.Name, item.Description, item.Path, item.IsActive, item.Id,
		).Error; err != nil {
			return fmt.Errorf("update error on id %s: %w", item.Id, err)
		}
	}
	return nil
}
