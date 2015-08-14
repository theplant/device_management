package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DeviceOut struct {
	gorm.Model
	Number     string
	Amount     int
	LendedBy   string
	BorrowedBy string
	LendedAt   time.Time
	ExperiedAt time.Time
}
