package connection

import (
	"fmt"
	"gigigarino/challengeMELI/internal/infraestructure/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	
	"github.com/jmoiron/sqlx"
)

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = 3306
	DB_NAME = "test-db"
	DB_USER = "root"
	DB_PASS = "secret"
)

var db *sqlx.DB //nolint:gochecknoglobals

func GetConnectionDB() (*sqlx.DB, error) {
	var err error

	if db == nil {
		db, err = sqlx.Connect("mysql", dbConnectionURL())
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
			return nil, fmt.Errorf("### DB ERROR: %w", err)
		}
	}

	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
		return fmt.Errorf("error instantiating migration: %w", err)
	}

	dbMigration, err := migrate.NewWithDatabaseInstance(
		"file://../../internal/infra/mysql/migration",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
		return fmt.Errorf("error instantiating migration: %w", err)
	}

	if err := dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error executing migration: %w", err)
	}

	return nil
}

func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}