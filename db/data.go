package db

import (
	"fmt"
	"github.com/qor/qor"
)

var DeviceCategories = [][]string{
	{"1", "自有设备"},
	{"2", "消耗品"},
	{"3", "客户设备"},
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

func CurrentClientDeviceIns(prop interface{}, c *qor.Context) (r [][]string) {
	var ris = []*ReportItem{}
	if err := DB.Where("client_device_in_id <> 0 OR client_device_in_id IS NOT NULL").Find(&ris).Error; err != nil {
		panic(err)
	}
	for _, ri := range ris {
		r = append(r, []string{fmt.Sprintf("%d", ri.ClientDeviceInID), fmt.Sprintf("%s - 数量：%d - %s", ri.DeviceName, ri.Count, ri.ClientName)})
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
	if err := DB.Find(&employees).Error; err != nil {
		panic(err)
	}
	for _, e := range employees {
		r = append(r, []string{fmt.Sprintf("%d", e.ID), e.Name})
	}
	return
}
