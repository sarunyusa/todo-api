package kernel

import (
	"github.com/jinzhu/gorm"
	"time"
	"todo/pkg/logger"
	"todo/pkg/ptr"
)

func loadDevConfigFile() *ConfigFile {
	config := &ConfigFile{
		Database: &DatabaseConfig{
			ConnectionString:   ptr.String("postgresql://postgres:P@ssw0rd@localhost:9990?sslmode=disable"),
			DatabaseName:       ptr.String("todo-local"),
			GormDebug:          ptr.Bool(false),
			CreateDbIfNotExist: ptr.Bool(true),
			MigrateDB:          ptr.Bool(true),
			RecreateDb:         ptr.Bool(false),
		},
		Todo: &TodoConfig{
			Host: ptr.String("localhost"),
			Port: ptr.Int(8080),
		},
	}

	return config
}

func (c *ConfigFile) ConnectDB() {
	log := logger.New("ConfigFile").WithServiceInfo("WaitForDB")
	doCheck := func() error {
		log.Infoln("type connecting to db", c.Database.GetConnectionString())
		conn, err := gorm.Open("postgres", c.Database.GetConnectionString())
		if err != nil {
			log.Infoln("failed to connect to db", err)
			return err
		} else {
			log.Infoln("db is now reachable", err)
		}

		defer func() { _ = conn.Close() }()

		err = conn.Raw("select 1").Error

		if err != nil {
			log.Infoln("Select 1 with err =", err)
			return err
		}

		return nil
	}

	for doCheck() != nil {
		log.Info("wait for db")
		time.Sleep(3 * time.Second)
	}
}

func (c *ConfigFile) MustPrepareDatabase() {
	log := logger.New("PrepareDatabase").WithField("database", c.Database.DatabaseName)
	c.Database.ensureDatabase(log)
	c.Database.switchDBConnectionOrPanic()
	c.Database.mustMigrateDB(log)
}
