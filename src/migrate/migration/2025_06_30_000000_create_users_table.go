package migration

import (
	"gorm.io/gorm"

	"gbase/src/core/helper"
	"gbase/src/core/migrate"
)

var CreateUsersTable = migrate.MigrateSql{
	Name: helper.GetCurrentFileNameNoExt(),
	Up: func(db *gorm.DB) error {
		return db.Exec(`
			CREATE TABLE "users" (
				"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				"order" INTEGER NOT NULL DEFAULT 0,
				"name" VARCHAR(100) NOT NULL DEFAULT '',
				
				"isDeleted" TINYINT NOT NULL DEFAULT 0,
				"createdAt" VARCHAR(19) NOT NULL DEFAULT '',
				"updatedAt" VARCHAR(19) NOT NULL DEFAULT ''
			);
		`).Error
	},
}
