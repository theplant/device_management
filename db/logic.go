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

func (e DeviceCheckCompany) HolderName() string {
	return e.Name
}
func (e DeviceCheckCompany) HolderID() uint {
	return e.ID
}
func (e DeviceCheckCompany) HolderType() string {
	return "DeviceCheckCompany"
}

func (dOut DeviceOut) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(dOut.FromReportItemID, dOut.ToWhomID, "Employee", int(dOut.Quantity))
	return
}

func (dOut DeviceOut) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(dOut.FromReportItemID, dOut.ToWhomID, "Employee", -1*int(dOut.Quantity))
	return
}

func (dIn DeviceIn) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse", dIn.Quantity)
	return
}

func (dIn DeviceIn) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse", -1*dIn.Quantity)
	return
}

func (dIn ClientDeviceCheckIn) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse", dIn.Quantity)
	return
}

func (dIn ClientDeviceCheckIn) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(dIn.FromReportItemID, dIn.ToWarehouseID, "Warehouse", -1*dIn.Quantity)
	return
}

func (dOut ClientDeviceCheckOut) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(dOut.FromReportItemID, dOut.ToDeviceCheckCompanyID, "DeviceCheckCompany", dOut.Quantity)
	return
}

func (dOut ClientDeviceCheckOut) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(dOut.FromReportItemID, dOut.ToDeviceCheckCompanyID, "DeviceCheckCompany", -1*dOut.Quantity)
	return
}

func (cIn ConsumableIn) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(cIn.ReportItemID, 0, "Warehouse", -1*cIn.Quantity)
	return
}

func (cIn ConsumableIn) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(cIn.ReportItemID, 0, "Warehouse", cIn.Quantity)
	return
}

func (cOut ConsumableOut) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(cOut.ReportItemID, 0, "Warehouse", cOut.Quantity)
	return
}

func (cOut ConsumableOut) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(cOut.ReportItemID, 0, "Warehouse", -1*cOut.Quantity)
	return
}

func (cdOut ClientDeviceOut) AfterCreate(db *gorm.DB) (err error) {
	err = moveDeviceByID(cdOut.ReportItemID, 0, "Warehouse", int(cdOut.Quantity))
	return
}

func (cdOut ClientDeviceOut) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(cdOut.ReportItemID, 0, "Warehouse", -1*int(cdOut.Quantity))
	return
}

func (cdIn ClientDeviceIn) AfterCreate(db *gorm.DB) (err error) {
	var holder DeviceHolder
	holder, err = holderByIDType(cdIn.WarehouseID, "Warehouse")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = getOrCreateReportItem(holder, nil, &cdIn, cdIn.Quantity)
	return
}

func (cdIn ClientDeviceIn) BeforeDelete(db *gorm.DB) (err error) {
	err = moveDeviceByID(0, cdIn.WarehouseID, "Warehouse", -1*int(cdIn.Quantity))
	return
}

func moveDeviceByID(fromReportItemID uint, toHolderId uint, toHolderType string, quantity int) (err error) {
	var d Device
	var cdIn ClientDeviceIn
	var from, to DeviceHolder
	from, to, d, cdIn, _, _ = fromToDevice(fromReportItemID, toHolderId, toHolderType)

	var fromRi, toRi *ReportItem

	var fcount = 0
	if from != nil {
		fromRi, err = getOrCreateReportItem(from, &d, &cdIn, 0)
		if err != nil {
			log.Println(err)
			return
		}
		fcount = fromRi.Count - quantity
	}

	tcount := 0
	if to != nil {
		toRi, err = getOrCreateReportItem(to, &d, &cdIn, 0)
		if err != nil {
			log.Println(err)
			return
		}
		tcount = toRi.Count + quantity
	}

	if fcount < 0 || tcount < 0 {
		err = errors.New(fmt.Sprintf("数量输入有误"))
		return
	}

	if from != nil {
		err = DB.Model(&fromRi).Where("id = ?", fromRi.ID).UpdateColumn("count", gorm.Expr("count - ?", quantity)).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	if to != nil {
		err = DB.Model(&toRi).Where("id = ?", toRi.ID).UpdateColumn("count", gorm.Expr("count + ?", quantity)).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func fromToDevice(fromReportItemID uint, toHolderId uint, toHolderType string) (from DeviceHolder, to DeviceHolder, d Device, cdIn ClientDeviceIn, fromRi ReportItem, err error) {

	if fromReportItemID > 0 {
		err = DB.Find(&fromRi, fromReportItemID).Error

		if err != nil {
			log.Println(err)
			return
		}
	}

	if fromRi.DeviceID > 0 {
		err = DB.Find(&d, fromRi.DeviceID).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	if fromRi.ClientDeviceInID > 0 {
		err = DB.Find(&cdIn, fromRi.ClientDeviceInID).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	from, err = holderByIDType(fromRi.WhoHasThemID, fromRi.WhoHasThemType)
	if err != nil {
		log.Println(err)
		return
	}

	if toHolderId > 0 {
		to, err = holderByIDType(toHolderId, toHolderType)

		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func holderByIDType(id uint, t string) (h DeviceHolder, err error) {
	switch t {
	case "Employee":
		employee := Employee{}
		err = DB.Find(&employee, id).Error
		h = employee
	case "Warehouse":
		warehouse := Warehouse{}
		err = DB.Find(&warehouse, id).Error
		h = warehouse
	case "DeviceCheckCompany":
		dcc := DeviceCheckCompany{}
		err = DB.Find(&dcc, id).Error
		h = dcc
	}
	return
}

func (d Device) AfterCreate(db *gorm.DB) (err error) {

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, d.WarehouseID).Error
	if err != nil {
		return
	}

	_, err = getOrCreateReportItem(warehouse, &d, nil, d.TotalQuantity)
	return
}

func (d Device) BeforeUpdate(db *gorm.DB) (err error) {

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, d.WarehouseID).Error
	if err != nil {
		return
	}

	var ri *ReportItem
	ri, err = getOrCreateReportItem(warehouse, &d, nil, d.TotalQuantity)
	if err != nil {
		return
	}

	var oldDev Device
	err = DB.Find(&oldDev, d.ID).Error
	if err != nil {
		return
	}

	inc := d.TotalQuantity - oldDev.TotalQuantity

	if ri.Count+int(inc) < 0 {
		err = errors.New(fmt.Sprintf("更新后的库存数量不能小于零，在库%d", ri.Count))
		return
	}

	err = DB.Model(&ReportItem{}).Where(&ReportItem{ID: ri.ID}).UpdateColumns(&ReportItem{Count: ri.Count + int(inc)}).Error
	if err != nil {
		return
	}

	err = DB.Model(&ReportItem{}).Where(&ReportItem{DeviceID: d.ID}).UpdateColumns(&ReportItem{DeviceCode: d.Code, DeviceName: d.Name}).Error
	return
}

func (d Device) BeforeDelete(db *gorm.DB) (err error) {

	warehouse := Warehouse{}
	err = DB.Find(&warehouse, d.WarehouseID).Error
	if err != nil {
		return
	}

	var ri *ReportItem
	ri, err = getOrCreateReportItem(warehouse, &d, nil, 0)
	if uint(ri.Count) != d.TotalQuantity {
		err = errors.New(fmt.Sprintf("有人带出设备%s，不能删除，当前库存数量%d，总数量%d", d.Name, ri.Count, d.TotalQuantity))
		return
	}

	err = DB.Delete(&ri).Error
	return
}

func getOrCreateReportItem(holder DeviceHolder, device *Device, cdIn *ClientDeviceIn, count uint) (r *ReportItem, err error) {
	var reportItem ReportItem

	if cdIn != nil && cdIn.ID > 0 {
		DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID, WhoHasThemID: holder.HolderID(), WhoHasThemType: holder.HolderType()}).Find(&reportItem)
	} else {
		fmt.Println(device, holder)
		DB.Where(&ReportItem{DeviceID: device.ID, WhoHasThemID: holder.HolderID(), WhoHasThemType: holder.HolderType()}).Find(&reportItem)
	}

	if reportItem.ID > 0 {
		r = &reportItem
		return
	}

	reportItem = ReportItem{
		WhoHasThemName: holder.HolderName(),
		WhoHasThemID:   holder.HolderID(),
		WhoHasThemType: holder.HolderType(),
		Count:          int(count),
	}

	if device != nil {
		reportItem.DeviceID = device.ID
		reportItem.DeviceCode = device.Code
		reportItem.DeviceName = device.Name
		reportItem.DeviceCategoryID = device.CategoryID
	}

	if cdIn != nil {
		reportItem.ClientDeviceInID = cdIn.ID
		reportItem.ClientName = cdIn.ClientName
		reportItem.DeviceName = cdIn.DeviceName
	}

	err = DB.Create(&reportItem).Error
	if err != nil {
		log.Println(err)
	}
	r = &reportItem
	return
}
