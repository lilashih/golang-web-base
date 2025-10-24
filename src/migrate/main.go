package migrate

import (
	"gbase/src/core/db"
	"gbase/src/core/migrate"
	"gbase/src/migrate/migration"
)

func Run() {
	// 所有 migration 都登記在這裡
	list := []migrate.MigrateSql{
		migration.CreateUsersTable,
		migration.CreatePermissionsTable,
		migration.CreateSettingsTable,

		migration.InsertPermission,
		migration.InsertSetting,
	}

	db.DisableDebug()

	migrate.Run(db.DB, list)
}
