package dbs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type logers struct {
	Username string
	Password string
}

var (
	UsersLog []logers
	Conn     *gorm.DB
)
