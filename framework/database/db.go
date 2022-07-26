package database

import (
	"log"

	"github.com/MaxHariel/encoder/domain"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db          *gorm.DB
	Dsn         string
	DsnTest     string
	DbType      string
	DbTypeTest  string
	Debug       bool
	AutoMigrate bool
	Env         string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstace := NewDb()
	dbInstace.Env = "Test"
	dbInstace.DbTypeTest = "sqlite3"
	dbInstace.DsnTest = ":memory:"
	dbInstace.AutoMigrate = true
	dbInstace.Debug = true

	connection, err := dbInstace.Connect()

	if err != nil {
		log.Fatalf("Test db error: %s", err.Error())
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "Test" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn))
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.Debug()
	}

	if d.AutoMigrate {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, err
}
