package db

import (
	"fmt"
	"github.com/qor/qor"
)

var DeviceCategories = [][]string{
	{"1", "自有设备"},
	{"2", "消耗品"},
	// {"3", "客户设备"},
}

func WarehouseCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*Warehouse{}
	if err := DB.Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), wh.Name})
	}
	return
}

func DeviceCheckCompanyCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*DeviceCheckCompany{}
	if err := DB.Order("name ASC").Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), wh.Name})
	}
	return
}

func CurrentWarehouseDeviceCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*ReportItem{}
	if err := DB.Where("count > 0 AND who_has_them_type = 'Warehouse' AND device_category_id = 1 AND client_device_in_id = 0").Order("device_name ASC").Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), fmt.Sprintf("[%s] %s - 剩余数量: %d - 仓库: %s", wh.DeviceCode, wh.DeviceName, wh.Count, wh.WhoHasThemName)})
	}
	return
}

func CurrentDeviceCheckCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*ReportItem{}
	if err := DB.Where("count > 0 AND who_has_them_type = 'DeviceCheckCompany'").Order("client_name ASC, device_name ASC").Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), fmt.Sprintf("[%s] %s - 数量: %d - 送检公司: %s", wh.ClientName, wh.DeviceName, wh.Count, wh.WhoHasThemName)})
	}
	return
}

func CurrentClientDeviceCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*ReportItem{}
	if err := DB.Where("count > 0 AND client_device_in_id > 0 AND who_has_them_type = 'Warehouse'").Order("client_name ASC, device_name").Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), fmt.Sprintf("[%s] %s - 剩余数量: %d - 仓库: %s", wh.ClientName, wh.DeviceName, wh.Count, wh.WhoHasThemName)})
	}
	return
}

func CurrentEmployeeDeviceCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*ReportItem{}
	if err := DB.Where("count > 0 AND who_has_them_type = 'Employee' AND client_device_in_id = 0").Order("device_name ASC").Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), fmt.Sprintf("[%s] %s - 带出数量: %d - 员工: %s", wh.DeviceCode, wh.DeviceName, wh.Count, wh.WhoHasThemName)})
	}
	return
}

func CurrentConsumableCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*ReportItem{}
	if err := DB.Where("count > 0 AND who_has_them_type = 'Warehouse' AND device_category_id = 2 AND client_device_in_id = 0").Order("device_name ASC").Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), fmt.Sprintf("[%s] %s - 库存数量: %d - 仓库: %s", wh.DeviceCode, wh.DeviceName, wh.Count, wh.WhoHasThemName)})
	}
	return
}

// func ClientDeviceOutValuer(resource interface{}, ctx *qor.Context) interface{} {
// 	id := resource.(*ClientDeviceOut).ClientDeviceInID
// 	cdIn := &ClientDeviceIn{}
// 	DB.Find(cdIn, id)
// 	return fmt.Sprintf("%s - 数量：%d - %s", cdIn.DeviceName, cdIn.Quantity, cdIn.ClientName)
// }

func EmployeeCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var employees = []*Employee{}
	if err := DB.Order("name ASC").Find(&employees).Error; err != nil {
		panic(err)
	}
	for _, e := range employees {
		r = append(r, []string{fmt.Sprintf("%d", e.ID), e.Name})
	}
	return
}
