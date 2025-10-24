package migration

import (
	"gorm.io/gorm"

	"gbase/src/core/helper"
	"gbase/src/core/migrate"
)

var CreatePermissionsTable = migrate.MigrateSql{
	Name: helper.GetCurrentFileNameNoExt(),
	Up: func(db *gorm.DB) error {
		return db.Exec(`
			CREATE TABLE "permissions" (
				"id" VARCHAR(255),
				"groupOrder" integer NOT NULL DEFAULT '0',
				"order" integer NOT NULL DEFAULT '0',
				"system" VARCHAR(100) NOT NULL DEFAULT '',
				"group1" VARCHAR(100) NOT NULL DEFAULT '',
				"group2" VARCHAR(100) NOT NULL DEFAULT '',
				"name" VARCHAR(100) NOT NULL DEFAULT '',
				"description" VARCHAR(50) NOT NULL DEFAULT '',
				"path" VARCHAR(300) NOT NULL DEFAULT '',
				"isActive" boolean NOT NULL DEFAULT '0',
				"createdAt" VARCHAR(19) NOT NULL DEFAULT '',
				"updatedAt" VARCHAR(19) NOT NULL DEFAULT '',
				PRIMARY KEY ("id")
			);
		`).Error
	},
}
