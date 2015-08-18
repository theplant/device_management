package db

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/validations"
	"strings"
)

func (device Device) Validate(db *gorm.DB) {
	var deviceInDb Device
	db.Where("code = ?", device.Code).First(&deviceInDb)

	if deviceInDb.ID != device.ID {
		db.AddError(validations.NewError(device, "Code", "代码已经存在了，不能重复"))
	}

	if device.Name == "" {
		db.AddError(validations.NewError(device, "Name", "设备名称不能为空"))
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
