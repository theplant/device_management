package main

import (
	"github.com/jinzhu/gorm"
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
	customerDeviceIncoming.Meta(&admin.Meta{Name: "Client", Type: "select_one", Label: "客户名"})
	customerDeviceIncoming.Meta(&admin.Meta{Name: "Device", Type: "select_one", Label: "设备名"})
	customerDeviceIncoming.Scope(&admin.Scope{
		Default: true,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Preload("Device").Preload("Client")
		},
	})
	customerDeviceIncoming.IndexAttrs("Client", "Device", "Quantity", "Date")
	customerDeviceIncoming.EditAttrs(customerDeviceIncoming.IndexAttrs()...)
	customerDeviceIncoming.NewAttrs(customerDeviceIncoming.IndexAttrs()...)
	customerDeviceOutcoming := adm.AddResource(&db.CustomerDeviceOutcoming{}, &admin.Config{Menu: []string{"日常操作"}})
	customerDeviceOutcoming.Meta(&admin.Meta{Name: "Client", Type: "select_one", Label: "客户名"})
	customerDeviceOutcoming.Meta(&admin.Meta{Name: "Device", Type: "select_one", Label: "设备名"})
	customerDeviceOutcoming.Scope(&admin.Scope{
		Default: true,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Preload("Device").Preload("Client")
		},
	})
	customerDeviceOutcoming.IndexAttrs("Client", "Device", "Quantity", "Date")
	customerDeviceOutcoming.EditAttrs(customerDeviceOutcoming.IndexAttrs()...)
	customerDeviceOutcoming.NewAttrs(customerDeviceOutcoming.EditAttrs()...)

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
