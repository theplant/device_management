package db

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/validations"
	"time"
)

type Device struct {
	gorm.Model
	Name     string
	Number   string `sql:"unique"`
	Total    int
	AvailableAmount int
	Category string
}

type CustomerDeviceIncoming struct {
	gorm.Model
	DeviceID    int
	Device      Device
	ClientID    int
	Client      Client
	Quantity    int
	Date        time.Time
	WareHouseID int
	WareHouse   WareHouse
}

type CustomerDeviceOutcoming struct {
	gorm.Model
	DeviceID    int
	Device      Device
	ClientID    int
	Client      Client
	Quantity    int
	Date        time.Time
	WareHouseID int
	WareHouse   WareHouse
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
	Name   string
	Mobile string
}

type ReportItem struct {
	gorm.Model
	WhoHasThem  string
	InWareHouse bool
	CompanyName string
	DeviceName  string
	DeviceCode  string
	Count       int
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

func (customerDeviceIncoming CustomerDeviceIncoming) Validate(db *gorm.DB) {
	if customerDeviceIncoming.Quantity > customerDeviceIncoming.Device.Total {
		db.AddError(validations.NewError(customerDeviceIncoming, "Quantity", "超过库存数量"))
	}
}

func (customerDeviceOutcoming CustomerDeviceOutcoming) Validate(db *gorm.DB) {
	if customerDeviceOutcoming.Quantity > customerDeviceOutcoming.Device.Total {
		db.AddError(validations.NewError(customerDeviceOutcoming, "Quantity", "超过库存数量"))
	}
}
