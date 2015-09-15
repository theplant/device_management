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
	Code          string    `sql:"unique_index:idx_code_deleted_at"`
	TypeSize      string    //型号规格
	UnitName      string    //计量单位
	MakerName     string    //生产厂家
	MakeDate      time.Time //生产日期
	FromSource    string    // 资产来源
	Note          string    // 备注
	TotalQuantity uint
	WarehouseID   uint
	CategoryID    uint
	CategoryName  string
}

type Warehouse struct {
	gorm.Model
	Name    string
	Address string
}

type DeviceCheckCompany struct {
	gorm.Model
	Name    string
	Address string
}

type Employee struct {
	gorm.Model
	Name   string
	Mobile string
	Title  string // 职位
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
	DeviceName    string
	ClientName    string
	Quantity      uint
	Date          time.Time
	WarehouseID   uint
	WarehouseName string
	ByWhomID      uint
	ByWhomName    string
}

type ClientDeviceOut struct {
	gorm.Model
	ReportItemID  uint
	DeviceName    string
	ClientName    string
	Quantity      uint
	WarehouseName string
	Date          time.Time
	ByWhomID      uint
	ByWhomName    string
}

type ClientDeviceCheckIn struct {
	gorm.Model
	FromReportItemID           uint
	FromDeviceCheckCompanyName string
	DeviceName                 string
	ClientName                 string
	Quantity                   int
	ToWarehouseID              uint
	ToWarehouseName            string
	Date                       time.Time
	ByWhomID                   uint
	ByWhomName                 string
}

type ClientDeviceCheckOut struct {
	gorm.Model
	FromReportItemID         uint
	FromWarehouseName        string
	DeviceName               string
	ClientName               string
	Quantity                 int
	ToDeviceCheckCompanyID   uint
	ToDeviceCheckCompanyName string
	Date                     time.Time
	ByWhomID                 uint
	ByWhomName               string
}

type ConsumableIn struct {
	gorm.Model
	ReportItemID  uint
	DeviceName    string
	Quantity      int
	WarehouseName string
	ByWhomID      uint
	ByWhomName    string
	Date          time.Time
}

type ConsumableOut struct {
	gorm.Model
	ReportItemID  uint
	DeviceName    string
	Quantity      int
	WarehouseName string
	ToWhomID      uint
	ToWhomName    string
	ByWhomID      uint
	ByWhomName    string
	Date          time.Time
}

// report data
type ReportItem struct {
	ID                 uint `gorm:"primary_key"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
	WhoHasThemName     string
	WhoHasThemID       uint
	WhoHasThemType     string
	DeviceID           uint
	DeviceName         string
	DeviceCode         string
	DeviceCategoryID   uint
	OperatedByWhomID   uint
	OperatedByWhomName string
	Count              int
	ClientDeviceInID   uint
	ClientName         string
}
