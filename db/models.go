package db

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model
	Name     string
	Number   string
	Total    int
	Category string
}
