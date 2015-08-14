package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

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

// func AvailableDevices() {
// 	var reportItems *[]ReportItem
// 	var counts *[]int
// 	db.DB.Find(&reportItems).Where("is_ware_house = ? AND count > ?", true, 0)

// 	var devices []*db.Device
// 	var outNumbers []string
// 	db.DB.Find(&devices).Where("available_amount > ?", 0).Pluck("number", &outNumbers)
// }
