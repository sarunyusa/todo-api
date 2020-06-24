package automigrate

import "github.com/jinzhu/gorm"

var tableList = []interface{}{}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(tableList...)
}
