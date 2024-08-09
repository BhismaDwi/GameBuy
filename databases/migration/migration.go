package migration

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func Initiator(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	// Check if migrations have already been applied
	// records, err := migrate.GetMigrationRecords(dbParam, "postgres")
	// if err != nil {
	// 	panic(err)
	// }

	// if len(records) > 0 {
	// Migrate down
	// 	_, downErr := migrate.Exec(dbParam, "postgres", migrations, migrate.Down)
	// 	if downErr != nil {
	// 		panic(downErr)
	// 	}
	// 	fmt.Println("Migrations down successfully!")
	// }

	// Migrate up
	n, upErr := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if upErr != nil {
		panic(upErr)
	}

	fmt.Println("Migration success, applied", n, "migrations!")
}
