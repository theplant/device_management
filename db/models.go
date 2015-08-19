package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

// master data
type Device struct {
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"unique_index:idx_code_deleted_at"`
	Name          string
	Code          string `sql:"unique_index:idx_code_deleted_at"`
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
	FromReportItemID uint
	FromWhomName     string
	DeviceName       string
	Quantity         int
	ToWarehouseID    uint
	ToWarehouseName  string
	ByWhomID         uint
	ByWhomName       string
	Date             time.Time
}

type DeviceOut struct {
	gorm.Model
	FromReportItemID  uint
	FromWarehouseName string
	DeviceName        string
	Quantity          uint
	ToWhomID          uint
	ToWhomName        string
	ByWhomID          uint
	ByWhomName        string
	Date              time.Time
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
	Count              int
	ClientDeviceInID   uint
}
