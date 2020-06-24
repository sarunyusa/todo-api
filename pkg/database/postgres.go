package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"todo/pkg/logger"
)

type Options struct {
	maxOpenConnections    int
	maxIdleConnections    int
	ConnectionLifeTimeMin int
}

var DefaultDatabaseOptions = Options{
	50,
	25,
	5,
}

func NewPostgresDialOrPanic(connectionString string, options Options) *gorm.DB {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	configurationDB(db,
		options.maxOpenConnections,
		options.maxIdleConnections,
		options.ConnectionLifeTimeMin,
	)

	return db
}

func NewDefaultPostgresDialOrPanic(connectionString string) *gorm.DB {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	configurationDB(db,
		DefaultDatabaseOptions.maxOpenConnections,
		DefaultDatabaseOptions.maxIdleConnections,
		DefaultDatabaseOptions.ConnectionLifeTimeMin,
	)

	return db
}

func configurationDB(db *gorm.DB, maxOpenConn int, maxIdleConn int, maxLiftTimeMin int) {
	log := logger.New("configurationDB")
	if db != nil {
		log.Println("[Util.DB] configurationDB")
		dialect := db.Dialect().GetName()
		if dialect != "postgres" {
			log.Println("No need to configure for ", dialect)
		} else {
			db.DB().SetMaxIdleConns(maxIdleConn)
			db.DB().SetMaxOpenConns(maxOpenConn)
			db.DB().SetConnMaxLifetime(time.Minute * time.Duration(maxLiftTimeMin))
			log.Printf("Set Idle %d, Max %d and LifetimeConn %d minutes. for DB type %s", maxIdleConn, maxOpenConn, maxLiftTimeMin, dialect)
		}
	}
}
