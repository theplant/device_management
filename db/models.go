package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

// master data
type Device struct {
	gorm.Model
	Name     string
	Code     string `sql:"unique"`
	Total    int
	Category string
}

type Warehouse struct {
	gorm.Model
	Name    string
	Address string
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

// operations data
type DeviceIn struct {
	gorm.Model
	Code        string
	Amount      uint
	ReceivedBy  string
	ReturnedBy  string
	ReceivedAt  time.Time
	WareHouseId int
}

type DeviceOut struct {
	gorm.Model
	Code          string
	FromWareHouse uint
	Amount        uint
	LendedBy      string
	BorrowedBy    string
	LendedAt      time.Time
	ExperiedAt    time.Time
}

type ClientDeviceIn struct {
	gorm.Model
	DeviceName  string
	ClientName  string
	Quantity    int
	Date        time.Time
	WarehouseID uint
	Warehouse   Warehouse
	ByWhomID    uint
	ByWhom      Employee
}

type ClientDeviceOut struct {
	gorm.Model
	ClientDeviceInID uint
	DeviceName       string
	ClientName       string
	Quantity         int
	WarehouseName    string
	Date             time.Time
	ByWhomID         uint
	ByWhom           Employee
}

type ConsumableIn struct {
	gorm.Model
	Name  string
	Code  string
	Count int
}

type ConsumableOut struct {
	gorm.Model
	Name  string
	Code  string
	Count int
}

// report data
type ReportItem struct {
	gorm.Model
	WhoHasThem       string
	WhoHasThemId     uint
	WhoHasThemType   string
	ClientID         uint
	ClientName       string
	DeviceID         uint
	DeviceName       string
	DeviceCode       string
	OperatedByWhomId uint
	OperatedByWhom   string
	Count            uint
	ClientDeviceInID uint
}
