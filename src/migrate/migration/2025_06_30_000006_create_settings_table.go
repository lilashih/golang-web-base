package migration

import (
	"gorm.io/gorm"

	"gbase/src/core/helper"
	"gbase/src/core/migrate"
)

var CreateSettingsTable = migrate.MigrateSql{
	Name: helper.GetCurrentFileNameNoExt(),
	Up: func(db *gorm.DB) error {
		return db.Exec(`
			CREATE TABLE "settings" (
				"id" VARCHAR(255),
				"order" integer NOT NULL DEFAULT '0',
				"group" VARCHAR(50) NOT NULL DEFAULT '',
				"name" VARCHAR(500) NOT NULL DEFAULT '',
				"type" VARCHAR(50) NOT NULL DEFAULT '',
				"value" text NOT NULL DEFAULT '',
				"option" json,
				"isViewable" boolean NOT NULL DEFAULT '0',
				"isEditable" boolean NOT NULL DEFAULT '0',
				"createdAt" VARCHAR(19) NOT NULL DEFAULT '',
				"updatedAt" VARCHAR(19) NOT NULL DEFAULT '',
				PRIMARY KEY ("id")
			);
		`).Error
	},
}
