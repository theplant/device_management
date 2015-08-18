package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/validations"
	"log"
	"strings"
)

func (device Device) Validate(db *gorm.DB) {
	var deviceInDb Device
	db.Where("number = ?", device.Code).First(&deviceInDb)

	if deviceInDb.ID != device.ID {
		db.AddError(validations.NewError(device, "Number", "Number already taken"))
	}

	if device.Name == "" {
		db.AddError(validations.NewError(device, "Name", "Name could not be blank"))
	}
}

func (cdIn *ClientDeviceIn) Validate(db *gorm.DB) {
	if len(strings.TrimSpace(cdIn.DeviceName)) == 0 {
		db.AddError(validations.NewError(cdIn, "DeviceName", "收入设备名不能为空"))
	}
	if len(strings.TrimSpace(cdIn.ClientName)) == 0 {
		db.AddError(validations.NewError(cdIn, "ClientName", "收入客户名不能为空"))
	}
	if cdIn.Warehouse.ID == 0 {
		db.AddError(validations.NewError(cdIn, "Wharehouse", "收入到的仓库不能为空"))
	}
	if cdIn.Quantity <= 0 {
		db.AddError(validations.NewError(cdIn, "Quantity", "收入设备的数量要大于0"))
	}
	if cdIn.ByWhom.ID == 0 {
		db.AddError(validations.NewError(cdIn, "ByWhom", "请选择操作员"))
	}
}

func (cdOut *ClientDeviceOut) Validate(db *gorm.DB) {

	if cdOut.ClientDeviceInID == 0 {
		db.AddError(validations.NewError(cdOut, "ClientDeviceInID", "选择收入过的客户设备"))
	}
	if cdOut.ByWhom.ID == 0 {
		db.AddError(validations.NewError(cdOut, "ByWhom", "请选择操作员"))
	}
	// panic(fmt.(cdOut.ClientDeviceInID))
	if cdOut.ClientDeviceInID > 0 {
		cdIn := &ClientDeviceIn{}
		DB.Preload("Warehouse").Find(cdIn, cdOut.ClientDeviceInID)
		// log.Println(cdIn)
		// panic("stop")

		cdOut.DeviceName = cdIn.DeviceName
		cdOut.ClientName = cdIn.ClientName
		cdOut.Quantity = cdIn.Quantity
		cdOut.WarehouseName = cdIn.Warehouse.Name
	}
}

func (cdOut ClientDeviceOut) AfterCreate(db *gorm.DB) (err error) {
	err = DB.Where(&ReportItem{ClientDeviceInID: cdOut.ClientDeviceInID}).Delete(&ReportItem{}).Error
	return
}

func (cdOut ClientDeviceOut) BeforeDelete(db *gorm.DB) (err error) {
	err = DB.Model(&ReportItem{}).Unscoped().Where(&ReportItem{ClientDeviceInID: cdOut.ClientDeviceInID}).UpdateColumn("deleted_at", nil).Error
	return
}

func (cdIn ClientDeviceIn) AfterCreate(db *gorm.DB) (err error) {
	err = DB.Unscoped().Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Delete(&ReportItem{}).Error
	if err != nil {
		return
	}

	err = createOrUpdateReportItem(
		cdIn.ID,
		cdIn.WarehouseID,
		cdIn.DeviceName,
		cdIn.ClientName,
		cdIn.ByWhomID,
		cdIn.Quantity)
	return
}

func (cdIn ClientDeviceIn) BeforeDelete(db *gorm.DB) (err error) {
	var cdOut ClientDeviceOut
	err = DB.Where(&ClientDeviceOut{ClientDeviceInID: cdIn.ID}).Find(&cdOut).Error
	if cdOut.ID > 0 {
		err = errors.New("设备已经还回，不能删除。")
		return
	}

	err = DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Delete(&ReportItem{}).Error
	return
}

// func (cdOut ClientDeviceOut) AfterCreate(db *gorm.DB) error {
// 	createOrUpdateReportItem(cdOut.Device, "", 0, false, cdOut.Client, cdOut.Quantity)
// 	return nil
// }

func createOrUpdateReportItem(clientDeviceInID uint, warehouseId uint, deviceName string, clientName string, operatedByWhomId uint, quantity int) (err error) {

	var reportItem ReportItem

	var warehouse Warehouse
	err = DB.Find(&warehouse, warehouseId).Error
	if err != nil {
		log.Println(err)
		return
	}

	var bywhom Employee
	err = DB.Find(&bywhom, operatedByWhomId).Error
	if err != nil {
		log.Println(err)
		return
	}

	if DB.Where(&ReportItem{DeviceName: deviceName, ClientName: clientName}).Find(&reportItem).RecordNotFound() {
		reportItem := ReportItem{
			WhoHasThem:       warehouse.Name,
			WhoHasThemId:     warehouse.ID,
			WhoHasThemType:   "Warehouse",
			ClientName:       clientName,
			DeviceName:       deviceName,
			Count:            uint(quantity),
			OperatedByWhomId: bywhom.ID,
			OperatedByWhom:   bywhom.Name,
			ClientDeviceInID: clientDeviceInID,
		}
		err = DB.Create(&reportItem).Error
		if err != nil {
			log.Println(err)
		}
		return
	}

	if quantity > 0 {
		err = DB.Model(&reportItem).UpdateColumn("count", gorm.Expr("count + ?", quantity)).Error
	} else {
		err = DB.Model(&reportItem).UpdateColumn("count", gorm.Expr("count - ?", quantity)).Error
	}
	if err != nil {
		log.Println(err)
	}

	return
}
