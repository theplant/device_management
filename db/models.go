package db

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/validations"
)

type Device struct {
	gorm.Model
	Name     string
	Number   string `sql:"unique"`
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
	Name    string
	Address string
}

type Consumable struct {
	gorm.Model
	Name  string
	Code  string
	Count int
}

type Client struct {
	gorm.Model
	Name   string
	Mobile string
}

type Employee struct {
	gorm.Model
	Name    string
	Mobile  string
}

func (device Device) Validate(db *gorm.DB) {
	var deviceInDb Device
	db.Where("number = ?", device.Number).First(&deviceInDb)

	if deviceInDb.ID != device.ID {
		db.AddError(validations.NewError(device, "Number", "Number already taken"))
	}

	if device.Name == "" {
		db.AddError(validations.NewError(device, "Name", "Name could not be blank"))
	}
}

type ReportItem struct {
	gorm.Model
	WhoHasThem  string
	CompanyName string
	DeviceName  string
	DeviceCode  string
	Count       int
}
