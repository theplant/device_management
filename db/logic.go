package db

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/validations"
)

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

func (cdIn ClientDeviceIn) Validate(db *gorm.DB) {
	if cdIn.Quantity > cdIn.Device.Total {
		db.AddError(validations.NewError(cdIn, "Quantity", "超过库存数量"))
	}
}

func (cdOut ClientDeviceOut) Validate(db *gorm.DB) {
	if cdOut.Quantity > cdOut.Device.Total {
		db.AddError(validations.NewError(cdOut, "Quantity", "超过库存数量"))
	}
}

func (cdIn ClientDeviceIn) AfterCreate(db *gorm.DB) error {
	createOrUpdateReportItem(cdIn.Device,
		cdIn.Warehouse.Name,
		cdIn.Warehouse.ID,
		true,
		cdIn.Client,
		cdIn.Quantity)
	return nil
}

func (cdOut ClientDeviceOut) AfterCreate(db *gorm.DB) error {
	createOrUpdateReportItem(cdOut.Device, "", 0, false, cdOut.Client, cdOut.Quantity)
	return nil
}

func createOrUpdateReportItem(device Device, whoHasThem string, whoHasThemId uint, inWareHouse bool, client Client, quantity int) {
	var reportItem ReportItem
	if DB.Where("device_id = ? AND client_id = ?", device.ID).Find(&reportItem).RecordNotFound() {
		reportItem := ReportItem{
			WhoHasThem:   whoHasThem,
			WhoHasThemId: whoHasThemId,
			InWareHouse:  inWareHouse,
			ClientID:     client.ID,
			ClientName:   client.Name,
			Count:        quantity,
			DeviceID:     device.ID,
			DeviceName:   device.Name,
			DeviceCode:   device.Number,
		}
		if err := DB.Save(&reportItem).Error; err != nil {
			panic(err)
		}
	} else {
		if quantity > 0 {
			DB.Model(&reportItem).UpdateColumn("count", gorm.Expr("count + ?", quantity))
		} else {
			DB.Model(&reportItem).UpdateColumn("count", gorm.Expr("count - ?", quantity))
		}
	}
}
