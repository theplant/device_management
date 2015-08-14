package db

import (
// "time"
)

type Device struct {
	Login string `gorm:"primary_key" sql:"size:128;not null;unique"`
	// AvatarURL string `sql:"size:255;"`
	// Name      string `sql:"size:255;"`
	// Company   string `sql:"size:255;"`
	// Location  string `sql:"size:255;"`
	// Email     string `sql:"size:255;"`
	// Blog      string `sql:"size:255;"`
	// Bio       string `sql:"type:longtext;"`
}
