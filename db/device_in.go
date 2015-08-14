package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DeviceIn struct {
	gorm.Model
	Number      string
	Amount      int
	ReceivedBy  string
	ReturnedBy  string
	ReceivedAt  time.Time
	WareHouseId int
}
