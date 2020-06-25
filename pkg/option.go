package pkg

import "github.com/jinzhu/gorm"

type ServiceOptions struct {
	Db          *gorm.DB
	HttpAddress string
}

type TodoOptions struct {
	ServiceOptions
}
