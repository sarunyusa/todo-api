package kernel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"todo/pkg"
)

type ConfigFile struct {
	Database *DatabaseConfig
	Todo     *TodoConfig
}

type TodoConfig struct {
	Host *string
	Port *int
}

func (c *TodoConfig) GetHost() string {
	if c == nil || c.Host == nil {
		return "0.0.0.0"
	}

	return *c.Host
}

func (c *TodoConfig) GetPort() int {
	if c == nil || c.Port == nil {
		return 8080
	}

	return *c.Port
}

func (c *TodoConfig) GetHttpAddress() string {
	return fmt.Sprintf("%s:%d", c.GetHost(), c.GetPort())
}

func (c *ConfigFile) makeDBConnections() *gorm.DB {
	c.Database.switchDBConnectionOrPanic()
	db := c.Database.makeDBConnection()

	return db
}

type DatabaseConfig struct {
	ConnectionString   *string `yaml:"connection-string"`
	DatabaseName       *string `yaml:"database-name"`
	GormDebug          *bool   `yaml:"gorm-debug"`
	CreateDbIfNotExist *bool   `yaml:"create-db-if-not-exist"`
	RecreateDb         *bool   `yaml:"recreate-db"`
	MigrateDB          *bool   `yaml:"migrate-db"`
}

func (c *DatabaseConfig) GetConnectionString() string {
	if c == nil || c.ConnectionString == nil {
		panic("missing database connection string")
	}

	return *c.ConnectionString
}

func (c *DatabaseConfig) GetDatabaseName() string {
	if c == nil || c.DatabaseName == nil {
		return ""
	}

	return *c.DatabaseName
}

func (c *DatabaseConfig) GetGormDebug() bool {
	if c == nil || c.GormDebug == nil {
		return false
	}

	return *c.GormDebug
}

func (c *DatabaseConfig) GetRecreateDB() bool {
	if c == nil || c.RecreateDb == nil {
		return false
	}

	return *c.RecreateDb
}

func (c *DatabaseConfig) GetCreateDbIfNotExist() bool {
	if c == nil || c.CreateDbIfNotExist == nil {
		return false
	}

	return *c.CreateDbIfNotExist
}

func (c *DatabaseConfig) GetMigrateDB() bool {
	if c == nil || c.MigrateDB == nil {
		return false
	}

	return *c.MigrateDB
}

func (c *ConfigFile) GetTodoOptions() *pkg.TodoOptions {
	return &pkg.TodoOptions{
		ServiceOptions: pkg.ServiceOptions{
			Db:          c.makeDBConnections(),
			HttpAddress: c.Todo.GetHttpAddress(),
		},
	}
}
