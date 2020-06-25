package pkg

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Server struct {
	Http        http.Handler
	HttpAddress string
	DB          *gorm.DB
}
