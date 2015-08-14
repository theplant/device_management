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

type CustomerDeviceIncoming struct {
	gorm.Model
	CustomerName string `sql:"size:255;"`
	DeviceId     int
}

type WareHouse struct {
	gorm.Model
	Name      string
	Address   string
}
