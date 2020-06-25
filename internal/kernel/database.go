package kernel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
	"todo/internal/automigrate"
	"todo/pkg/database"
	"todo/pkg/logger"
	"todo/pkg/ptr"
)

func (c *DatabaseConfig) makeDBConnection() *gorm.DB {
	return database.NewDefaultPostgresDialOrPanic(c.GetConnectionString()).
		LogMode(c.GetGormDebug())
}

func (c *DatabaseConfig) recreateDatabaseIfExistOrPanic() {
	conn := c.makeDBConnection().Debug()
	defer func() { _ = conn.Close() }()

	err := conn.Exec(fmt.Sprintf(
		`DROP DATABASE IF EXISTS "%s"`,
		c.GetDatabaseName())).Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec(fmt.Sprintf(
		`CREATE DATABASE  "%s"`,
		c.GetDatabaseName())).Error
	if err != nil {
		panic(err)
	}
}

func (c *DatabaseConfig) createDatabaseOrPanic() {
	conn := c.makeDBConnection().Debug()
	defer func() { _ = conn.Close() }()

	dbExist := &struct {
		Exists bool `gorm:"column:exists"`
	}{}

	statement := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = ?) as exists`
	err := conn.Raw(statement, c.GetDatabaseName()).Scan(dbExist).Error

	if err != nil {
		panic(err)
	}

	if dbExist.Exists {
		return
	}

	err = conn.Exec(fmt.Sprintf(
		`CREATE DATABASE  "%s"`,
		c.GetDatabaseName())).Error
	if err != nil {
		panic(err)
	}
}

func (c *DatabaseConfig) ensureDatabase(log *logger.Logger) {
	log = log.WithServiceInfo("ensureDatabase")
	if c.GetDatabaseName() == "" {
		log.Println("no db name specified, do nothing")
		return
	}
	if c.GetRecreateDB() {
		log.Println("try to recreate if exist")
		c.recreateDatabaseIfExistOrPanic()
	} else if c.GetCreateDbIfNotExist() {
		log.Println("try to create db if not exist")
		c.createDatabaseOrPanic()
	}
}

func (c *DatabaseConfig) switchDBConnectionOrPanic() {
	if dbUrl, err := url.Parse(c.GetConnectionString()); err == nil {
		q := dbUrl.Query()
		q.Set("database", c.GetDatabaseName())
		dbUrl.RawQuery = q.Encode()
		c.ConnectionString = ptr.String(dbUrl.String())
	} else {
		panic(err)
	}
}

func (c *DatabaseConfig) mustMigrateDB(logger *logger.Logger) {
	logger.Info("database migration start...")

	defer func() {
		logger.Info("database migration done...")
	}()

	if !c.GetMigrateDB() {
		logger.Info("MigrateDB options is false, do nothing")
	}

	db := database.NewDefaultPostgresDialOrPanic(c.GetConnectionString()).
		LogMode(c.GetGormDebug())

	logger.Info("perform migration")
	automigrate.MigrateDB(db)

	logger.Info("closing db connection")
	err := db.Close()
	if err != nil {
		logger.Panic("database migration", err)
	}
}
