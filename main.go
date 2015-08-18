package main

import (
	// "github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/i18n"
	"github.com/qor/qor/i18n/backends/database"
	"github.com/qor/qor/roles"
	"github.com/theplant/device_management/db"
	"log"
	"net/http"
	"time"
)

func main() {
	roles.Register("admin", func(req *http.Request, cu qor.CurrentUser) bool {
		return true
	})

	adm := admin.New(&qor.Config{DB: &db.DB})

	adm.SetAuth(&Auth{})

	noUpdatePermission := roles.Deny(roles.Update, "admin")

	reportItem := adm.AddResource(&db.ReportItem{}, &admin.Config{Menu: []string{"查询"},
		Permission: roles.Deny(roles.Update, "admin").Deny(roles.Delete, "admin").Deny(roles.Create, "admin"),
	})
	reportItem.IndexAttrs("WhoHasThem", "ClientName", "DeviceName", "DeviceCode", "Count")

	cdIn := adm.AddResource(&db.ClientDeviceIn{}, &admin.Config{
		Menu:       []string{"日常操作"},
		Permission: noUpdatePermission,
	})

	cdIn.Meta(&admin.Meta{Name: "Warehouse", Type: "select_one", Collection: db.WarehouseCollection})
	cdIn.Meta(&admin.Meta{Name: "ByWhom", Type: "select_one", Collection: db.EmployeeCollection})
	// cdIn.Scope(&admin.Scope{
	// 	Default: true,
	// 	Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
	// 		return db.Preload("Device").Preload("Client")
	// 	},
	// })
	cdIn.IndexAttrs("ClientName", "DeviceName", "Warehouse", "Quantity", "ByWhom", "Date")
	cdIn.NewAttrs(cdIn.IndexAttrs()...)
	cdIn.Meta(&admin.Meta{Name: "Date", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		date := resource.(*db.ClientDeviceIn).Date
		if date.IsZero() {
			date = time.Now()
		}
		return date
	}})

	cdOut := adm.AddResource(&db.ClientDeviceOut{}, &admin.Config{
		Menu:       []string{"日常操作"},
		Permission: noUpdatePermission,
	})

	cdOut.Meta(&admin.Meta{Name: "ClientDeviceInID", Type: "select_one", Collection: db.CurrentClientDeviceIns})
	cdOut.Meta(&admin.Meta{Name: "ByWhom", Type: "select_one", Collection: db.EmployeeCollection})
	cdOut.Meta(&admin.Meta{Name: "Date", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		date := resource.(*db.ClientDeviceOut).Date
		if date.IsZero() {
			date = time.Now()
		}
		return date
	}})

	// cdOut.Scope(&admin.Scope{
	// 	Default: true,
	// 	Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
	// 		return db.Preload("Device").Preload("Client")
	// 	},
	// })
	cdOut.IndexAttrs("ClientName", "DeviceName", "Quantity", "WarehouseName", "ByWhom", "Date")
	cdOut.NewAttrs("ClientDeviceInID", "ByWhom", "Date")

	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"数据维护"}})
	device.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: db.DeviceCategories})

	deviceOut := adm.AddResource(&db.DeviceOut{}, &admin.Config{Menu: []string{"日常操作"}})
	deviceIn := adm.AddResource(&db.DeviceIn{}, &admin.Config{Menu: []string{"日常操作"}})

	_ = deviceIn
	// deviceIn.Meta(&admin.Meta{Name: "Code", Type: "select_one", Collection: inNumbers})

	deviceOut.NewAttrs("-LendedAt")
	_ = deviceOut
	// deviceOut.Meta(&admin.Meta{Name: "Number", Type: "select_one", Collection: outNumbers})

	consumableOut := adm.AddResource(&db.ConsumableOut{}, &admin.Config{Menu: []string{"日常操作"}})
	consumableOut.Meta(&admin.Meta{Name: "Name", Type: "string"})
	consumableOut.Meta(&admin.Meta{Name: "Code", Type: "string"})
	consumableOut.Meta(&admin.Meta{Name: "Count", Type: "int"})
	consumableOut.EditAttrs("Name", "Code", "Count")
	consumableOut.NewAttrs(consumableOut.EditAttrs()...)

	consumableIn := adm.AddResource(&db.ConsumableIn{}, &admin.Config{Menu: []string{"日常操作"}})
	consumableIn.Meta(&admin.Meta{Name: "Name", Type: "string"})
	consumableIn.Meta(&admin.Meta{Name: "Code", Type: "string"})
	consumableIn.Meta(&admin.Meta{Name: "Count", Type: "int"})
	consumableIn.EditAttrs("Name", "Code", "Count")
	consumableIn.NewAttrs(consumableIn.EditAttrs()...)

	adm.AddResource(&db.Employee{}, &admin.Config{Menu: []string{"数据维护"}})

	warehouse := adm.AddResource(&db.Warehouse{}, &admin.Config{Menu: []string{"数据维护"}})
	warehouse.Meta(&admin.Meta{Name: "Name", Type: "string"})
	warehouse.Meta(&admin.Meta{Name: "Address", Type: "string"})
	warehouse.EditAttrs("Name", "Address")
	warehouse.NewAttrs(warehouse.EditAttrs()...)

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
