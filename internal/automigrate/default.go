package automigrate

import (
	"github.com/jinzhu/gorm"
	"todo/entity"
)

var tableList = []interface{}{
	&entity.Todo{},
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(tableList...)
}
