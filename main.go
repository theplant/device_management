package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/i18n"
	"github.com/qor/qor/i18n/backends/database"
	"github.com/qor/qor/roles"
	"github.com/theplant/device_management/db"
	"html/template"
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
	reportItem.IndexAttrs("WhoHasThemName", "ClientName", "DeviceName", "DeviceCode", "Count")
	reportItem.Scope(&admin.Scope{
		Default: true,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Where("count > 0").Order("updated_at DESC")
		},
	})

	reportItem.Meta(&admin.Meta{
		Name: "WhoHasThemName",
		Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
			ri := resource.(*db.ReportItem)
			name := ri.WhoHasThemName
			if ri.WhoHasThemType == "Employee" {
				name = fmt.Sprintf(`<strong style="color:red">%s</string>`, name)
			}
			return template.HTML(name)
		},
	})

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

	deviceOut := adm.AddResource(&db.DeviceOut{}, &admin.Config{
		Menu:       []string{"日常操作"},
		Permission: noUpdatePermission,
	})

	deviceOut.IndexAttrs("ToWhomName", "FromWarehouseName", "DeviceName", "Quantity", "ByWhomName", "Date")
	deviceOut.NewAttrs("FromReportItemID", "Quantity", "ToWhomID", "ByWhomID", "Date")
	deviceOut.Meta(&admin.Meta{Name: "FromReportItemID", Type: "select_one", Collection: db.CurrentWarehouseDeviceCollection})
	deviceOut.Meta(&admin.Meta{Name: "ToWhomID", Type: "select_one", Collection: db.EmployeeCollection})
	deviceOut.Meta(&admin.Meta{Name: "ByWhomID", Type: "select_one", Collection: db.EmployeeCollection})
	deviceOut.Meta(&admin.Meta{Name: "Date", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		date := resource.(*db.DeviceOut).Date
		if date.IsZero() {
			date = time.Now()
		}
		return date
	}})

	deviceIn := adm.AddResource(&db.DeviceIn{}, &admin.Config{
		Menu:       []string{"日常操作"},
		Permission: noUpdatePermission,
	})
	deviceIn.IndexAttrs("FromWhomName", "ToWarehouseName", "DeviceName", "Quantity", "ByWhomName", "Date")
	deviceIn.NewAttrs("FromReportItemID", "Quantity", "ToWarehouseID", "ByWhomID", "Date")
	deviceIn.Meta(&admin.Meta{Name: "FromReportItemID", Type: "select_one", Collection: db.CurrentEmployeeDeviceCollection})
	deviceIn.Meta(&admin.Meta{Name: "ToWarehouseID", Type: "select_one", Collection: db.WarehouseCollection})
	deviceIn.Meta(&admin.Meta{Name: "ByWhomID", Type: "select_one", Collection: db.EmployeeCollection})
	deviceIn.Meta(&admin.Meta{Name: "Date", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		date := resource.(*db.DeviceIn).Date
		if date.IsZero() {
			date = time.Now()
		}
		return date
	}})

	consumableOut := adm.AddResource(&db.ConsumableOut{}, &admin.Config{
		Menu:       []string{"日常操作"},
		Permission: noUpdatePermission,
	})
	consumableOut.IndexAttrs("DeviceName", "Quantity", "ToWhomName", "WarehouseName", "ByWhomName", "Date")
	consumableOut.NewAttrs("ReportItemID", "Quantity", "ToWhomID", "ByWhomID", "Date")
	consumableOut.Meta(&admin.Meta{Name: "ReportItemID", Type: "select_one", Collection: db.CurrentConsumableCollection})
	consumableOut.Meta(&admin.Meta{Name: "ToWhomID", Type: "select_one", Collection: db.EmployeeCollection})
	consumableOut.Meta(&admin.Meta{Name: "ByWhomID", Type: "select_one", Collection: db.EmployeeCollection})
	consumableOut.Meta(&admin.Meta{Name: "Date", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		date := resource.(*db.ConsumableOut).Date
		if date.IsZero() {
			date = time.Now()
		}
		return date
	}})

	consumableIn := adm.AddResource(&db.ConsumableIn{}, &admin.Config{
		Menu:       []string{"日常操作"},
		Permission: noUpdatePermission,
	})

	consumableIn.IndexAttrs("DeviceName", "Quantity", "WarehouseName", "ByWhomName", "Date")
	consumableIn.NewAttrs("ReportItemID", "Quantity", "ToWhomID", "ByWhomID", "Date")
	consumableIn.Meta(&admin.Meta{Name: "ReportItemID", Type: "select_one", Collection: db.CurrentConsumableCollection})
	consumableIn.Meta(&admin.Meta{Name: "ByWhomID", Type: "select_one", Collection: db.EmployeeCollection})
	consumableIn.Meta(&admin.Meta{Name: "Date", Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
		date := resource.(*db.ConsumableIn).Date
		if date.IsZero() {
			date = time.Now()
		}
		return date
	}})

	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"数据维护"}})
	device.Meta(&admin.Meta{Name: "CategoryID", Type: "select_one", Collection: db.DeviceCategories})
	device.Meta(&admin.Meta{Name: "WarehouseID", Type: "select_one", Collection: db.WarehouseCollection})
	device.EditAttrs("Name", "Code", "TotalQuantity")
	device.IndexAttrs("Name", "Code", "TotalQuantity")

	// _ = device

	employee := adm.AddResource(&db.Employee{}, &admin.Config{Menu: []string{"数据维护"}})
	employee.IndexAttrs("Name", "Mobile")

	warehouse := adm.AddResource(&db.Warehouse{}, &admin.Config{Menu: []string{"数据维护"}})
	warehouse.Meta(&admin.Meta{Name: "Name", Type: "string"})
	warehouse.Meta(&admin.Meta{Name: "Address", Type: "string"})
	warehouse.EditAttrs("Name", "Address")
	warehouse.NewAttrs(warehouse.EditAttrs()...)
	warehouse.IndexAttrs("Name", "Address")

	I18nBackend := database.New(&db.DB)
	// config.I18n = i18n.New(I18nBackend)
	adm.AddResource(i18n.New(I18nBackend), &admin.Config{Menu: []string{"系统设置"}, Invisible: true})

	adm.MountTo("/admin", http.DefaultServeMux)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/report_items", 302)
	})

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/admin/report_items"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/admin/report_items"
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
