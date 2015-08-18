package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

// master data
type Device struct {
	gorm.Model
	Name          string
	Code          string `sql:"unique"`
	TotalQuantity uint
	WarehouseID   uint
	CategoryID    uint
}

type Warehouse struct {
	gorm.Model
	Name    string
	Address string
}

type Employee struct {
	gorm.Model
	Name   string
	Mobile string
}

// operations data
type DeviceIn struct {
	gorm.Model
	DeviceOutID  uint
	DeviceName   string
	FromWhomName string
	Quantity     int
	WarehouseID  uint
	Warehouse    Warehouse
	Date         time.Time
	ByWhomID     uint
	ByWhom       Employee
}

type DeviceOut struct {
	gorm.Model
	DeviceID    uint
	Device      Device
	Quantity    uint
	ToWhomID    uint
	ToWhom      Employee
	WarehouseID uint
	Warehouse   Warehouse
	ByWhomID    uint
	ByWhom      Employee
	Date        time.Time
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
	WhoHasThemName     string
	WhoHasThemID       uint
	WhoHasThemType     string
	ClientName         string
	DeviceID           uint
	DeviceName         string
	DeviceCode         string
	OperatedByWhomID   uint
	OperatedByWhomName string
	Count              uint
	ClientDeviceInID   uint
}
