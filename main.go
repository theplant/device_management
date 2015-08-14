package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/theplant/device_management/db"
)

func main() {
	adm := admin.New(&qor.Config{DB: &db.DB})

	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"设备管理"}})
	device.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"自有设备", "消耗品", "客户设备"}})

	reportItem := adm.AddResource(&db.ReportItem{}, &admin.Config{Menu: []string{"查询"}})
	_ = reportItem

	customerDeviceIncoming := adm.AddResource(&db.CustomerDeviceIncoming{}, &admin.Config{Menu: []string{"设备管理"}})
	customerDeviceIncoming.Meta(&admin.Meta{Name: "CustomerName", Type: "string"})
	customerDeviceIncoming.Meta(&admin.Meta{Name: "DeviceId", Type: "select_one", Collection: allDevices})
	customerDeviceOutcoming := adm.AddResource(&db.CustomerDeviceOutcoming{}, &admin.Config{Menu: []string{"设备管理"}})
	customerDeviceOutcoming.Meta(&admin.Meta{Name: "CustomerName", Type: "string"})
	customerDeviceOutcoming.Meta(&admin.Meta{Name: "DeviceId", Type: "select_one", Collection: allDevices})

	adm.AddResource(&db.Client{}, &admin.Config{Menu: []string{"人事管理"}})

	adm.AddResource(&db.Employee{}, &admin.Config{Menu: []string{"人事管理"}})

	adm.AddResource(&db.Consumable{}, &admin.Config{Menu: []string{"消耗品管理"}})
	adm.AddResource(&db.WareHouse{}, &admin.Config{Menu: []string{"设备管理"}})

	adm.AddResource(&db.Consumable{}, &admin.Config{Menu: []string{"设备管理"}})
	adm.MountTo("/admin", http.DefaultServeMux)

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func allDevices(resource interface{}, context *qor.Context) (results [][]string) {
	var devices []db.Device
	context.GetDB().Find(&devices)
	for _, device := range devices {
		results = append(results, []string{fmt.Sprint(device.ID), device.Name})
	}
	return
}
