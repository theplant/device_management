package main

import (
	"fmt"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/i18n"
	"github.com/qor/qor/i18n/backends/database"
	"github.com/theplant/device_management/db"
	"log"
	"net/http"
)

func main() {
	adm := admin.New(&qor.Config{DB: &db.DB})
	adm.SetAuth(&Auth{})

	reportItem := adm.AddResource(&db.ReportItem{}, &admin.Config{Menu: []string{"查询"}})
	_ = reportItem

	customerDeviceIncoming := adm.AddResource(&db.CustomerDeviceIncoming{}, &admin.Config{Menu: []string{"日常操作"}})
	customerDeviceIncoming.Meta(&admin.Meta{Name: "CustomerName", Type: "string", Label: "客户名"})
	customerDeviceIncoming.Meta(&admin.Meta{Name: "DeviceId", Type: "select_one", Collection: allDevices, Label: "设备名", Valuer: formatedDeviceName})
	customerDeviceIncoming.EditAttrs("CustomerName", "DeviceId")
	customerDeviceIncoming.NewAttrs(customerDeviceIncoming.EditAttrs()...)
	customerDeviceOutcoming := adm.AddResource(&db.CustomerDeviceOutcoming{}, &admin.Config{Menu: []string{"日常操作"}})
	customerDeviceOutcoming.Meta(&admin.Meta{Name: "CustomerName", Type: "string", Label: "客户名"})
	customerDeviceOutcoming.Meta(&admin.Meta{Name: "DeviceId", Type: "select_one", Collection: allDevices, Label: "设备名", Valuer: formatedDeviceName})
	customerDeviceOutcoming.EditAttrs("CustomerName", "DeviceId")
	customerDeviceOutcoming.NewAttrs(customerDeviceIncoming.EditAttrs()...)

	adm.AddResource(&db.Consumable{}, &admin.Config{Menu: []string{"日常操作"}})

	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"数据维护"}})
	device.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"自有设备", "消耗品", "客户设备"}})
	adm.AddResource(&db.WareHouse{}, &admin.Config{Menu: []string{"数据维护"}})
	adm.AddResource(&db.Client{}, &admin.Config{Menu: []string{"数据维护"}})
	adm.AddResource(&db.Employee{}, &admin.Config{Menu: []string{"数据维护"}})

	I18nBackend := database.New(&db.DB)
	// config.I18n = i18n.New(I18nBackend)
	adm.AddResource(i18n.New(I18nBackend), &admin.Config{Menu: []string{"系统设置"}})

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

func formatedDeviceName(resource interface{}, ctx *qor.Context) interface{} {
	var text string
	switch model := resource.(type) {
	case *db.CustomerDeviceIncoming:
		text = model.Device.Name
	case *db.CustomerDeviceOutcoming:
		text = model.Device.Name
	}

	return text
}

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {

	return &User{}
}

type User struct {
}

func (u User) AvailableLocales() []string {
	return []string{"zh_CN"}
}

func (u User) DisplayName() string {
	return "管理员"
}
