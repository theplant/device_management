package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

// master data
type Device struct {
	gorm.Model
	Name     string
	Number   string `sql:"unique"`
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
	Number      string
	Amount      int
	ReceivedBy  string
	ReturnedBy  string
	ReceivedAt  time.Time
	WareHouseId int
}

type DeviceOut struct {
	gorm.Model
	Number        string
	FromWareHouse int
	Amount        int
	LendedBy      string
	BorrowedBy    string
	LendedAt      time.Time
	ExperiedAt    time.Time
}

type ClientDeviceIn struct {
	gorm.Model
	DeviceID    int
	Device      Device
	ClientID    int
	Client      Client
	Quantity    int
	Date        time.Time
	WareHouseID int
	Warehouse   Warehouse
}

type ClientDeviceOut struct {
	gorm.Model
	DeviceID    int
	Device      Device
	ClientID    int
	Client      Client
	Quantity    int
	Date        time.Time
	WareHouseID int
	Warehouse   Warehouse
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
	WhoHasThem   string
	WhoHasThemId uint
	InWareHouse  bool
	ClientID     uint
	ClientName   string
	DeviceName   string
	DeviceCode   string
	DeviceID     uint
	Count        int
}
