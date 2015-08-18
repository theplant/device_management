package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type DeviceHolder interface {
	HolderName() string
	HolderID() uint
	HolderType() string
}

func (e Employee) HolderName() string {
	return e.Name
}
func (e Employee) HolderID() uint {
	return e.ID
}
func (e Employee) HolderType() string {
	return "Employee"
}

func (e Warehouse) HolderName() string {
	return e.Name
}
func (e Warehouse) HolderID() uint {
	return e.ID
}
func (e Warehouse) HolderType() string {
	return "Warehouse"
}

func (dOut DeviceOut) AfterCreate(db *gorm.DB) (err error) {

	warehouse, _, device, err1 := wed(dOut.WarehouseID, 0, dOut.DeviceID)
	err = err1
	if err != nil {
		return
	}

	err = moveDevice(warehouse, dOut.ToWhom, &device, uint(dOut.Quantity))

	return
}

func (dOut DeviceOut) BeforeDelete(db *gorm.DB) (err error) {
	var dIn DeviceIn
	DB.Where(&DeviceIn{DeviceOutID: dOut.ID}).Find(&dIn)
	if dIn.ID > 0 {
		err = errors.New("设备已经还回，不能删除。")
		return
	}

	warehouse, _, device, err1 := wed(dOut.WarehouseID, 0, dOut.DeviceID)
	err = err1
	if err != nil {
		return
	}

	err = moveDevice(dOut.ToWhom, warehouse, &device, uint(dOut.Quantity))

	return
}

func (dIn DeviceIn) AfterCreate(db *gorm.DB) (err error) {

	warehouse, _, _, err1 := wed(dIn.WarehouseID, 0, 0)
	err = err1
	if err != nil {
		return
	}

	var dOut DeviceOut
	err = DB.Find(&dOut, dIn.DeviceOutID).Error
	if err != nil {
		log.Println(err)
		return
	}

	_, fromWhom, device, err2 := wed(0, dOut.ToWhomID, dOut.DeviceID)
	err = err2
	if err != nil {
		return
	}

	err = moveDevice(&fromWhom, &warehouse, &device, uint(dIn.Quantity))
	return
}

func moveDevice(from DeviceHolder, to DeviceHolder, device *Device, quantity uint) (err error) {
	var fromRi, toRi *ReportItem
	fromRi, err = getOrCreateReportItem(from, device, 0)
	if err != nil {
		log.Println(err)
		return
	}

	toRi, err = getOrCreateReportItem(to, device, 0)
	if err != nil {
		log.Println(err)
		return
	}

	if fromRi.Count < quantity {
		err = errors.New(fmt.Sprintf("数量输入有误，不能大于%d", fromRi.Count))
		return
	}

	err = DB.Model(&fromRi).Where("id = ?", fromRi.ID).UpdateColumn("count", gorm.Expr("count - ?", quantity)).Error
	if err != nil {
		log.Println(err)
		return
	}

	err = DB.Model(&toRi).Where("id = ?", toRi.ID).UpdateColumn("count", gorm.Expr("count + ?", quantity)).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (d Device) AfterCreate(db *gorm.DB) (err error) {

	warehouse, _, _, err1 := wed(d.WarehouseID, 0, 0)
	err = err1
	if err != nil {
		return
	}

	_, err = getOrCreateReportItem(warehouse, &d, d.TotalQuantity)
	return
}

func getOrCreateReportItem(holder DeviceHolder, device *Device, count uint) (r *ReportItem, err error) {
	var reportItem ReportItem

	DB.Where(&ReportItem{DeviceID: device.ID, WhoHasThemID: holder.HolderID(), WhoHasThemType: holder.HolderType()}).Find(&reportItem)

	if reportItem.ID > 0 {
		r = &reportItem
		return
	}

	reportItem = ReportItem{
		WhoHasThemName: holder.HolderName(),
		WhoHasThemID:   holder.HolderID(),
		WhoHasThemType: holder.HolderType(),
		DeviceID:       device.ID,
		DeviceCode:     device.Code,
		DeviceName:     device.Name,
		Count:          count,
	}

	err = DB.Create(&reportItem).Error
	if err != nil {
		log.Println(err)
	}
	r = &reportItem
	return
}

func (cdIn DeviceIn) BeforeDelete(db *gorm.DB) (err error) {
	var cdOut ClientDeviceOut
	err = DB.Where(&ClientDeviceOut{ClientDeviceInID: cdIn.ID}).Find(&cdOut).Error
	if cdOut.ID > 0 {
		err = errors.New("设备已经还回，不能删除。")
		return
	}

	err = DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Delete(&ReportItem{}).Error
	return
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

func wed(warehouseID uint, employeeID uint, deviceID uint) (warehouse Warehouse, employee Employee, device Device, err error) {

	if warehouseID > 0 {
		err = DB.Find(&warehouse, warehouseID).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	if employeeID > 0 {
		err = DB.Find(&employee, employeeID).Error
		if err != nil {
			log.Println(err)
		}

	}

	if deviceID > 0 {
		err = DB.Find(&device, deviceID).Error
		if err != nil {
			log.Println(err)
		}

	}

	return
}

func createOrUpdateReportItem(clientDeviceInID uint, warehouseId uint, deviceName string, clientName string, operatedByWhomId uint, quantity int) (err error) {

	var reportItem ReportItem

	warehouse, bywhom, _, err1 := wed(warehouseId, operatedByWhomId, 0)
	err = err1
	if err != nil {
		return
	}

	if DB.Where(&ReportItem{ClientDeviceInID: clientDeviceInID}).Find(&reportItem).RecordNotFound() {
		reportItem := ReportItem{
			WhoHasThemName:     warehouse.Name,
			WhoHasThemID:       warehouse.ID,
			WhoHasThemType:     "Warehouse",
			ClientName:         clientName,
			DeviceName:         deviceName,
			Count:              uint(quantity),
			OperatedByWhomID:   bywhom.ID,
			OperatedByWhomName: bywhom.Name,
			ClientDeviceInID:   clientDeviceInID,
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
