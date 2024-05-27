package newsletter

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	driver *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", "base.db")
	return &Database{driver: db}, err
}

func (db *Database) RunMigrations() {
	log.Println("Running migrations")
	m, err := migrate.New("file://migrations/", "sqlite3://base.db")

	if err != nil {
		log.Println("Error running migrations")
		log.Fatal(err)
	}

	m.Up()
}

func (db *Database) Close() { db.driver.Close() }
