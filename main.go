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

	cdIn := adm.AddResource(&db.ClientDeviceIn{}, &admin.Config{Menu: []string{"日常操作"}})
	cdIn.Meta(&admin.Meta{Name: "Client", Type: "select_one"})
	cdIn.Meta(&admin.Meta{Name: "Device", Type: "select_one"})
	cdIn.Meta(&admin.Meta{Name: "Warehouse", Type: "select_one", Collection: db.WarehouseCollection})
	cdIn.Scope(&admin.Scope{
		Default: true,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Preload("Device").Preload("Client")
		},
	})
	cdIn.IndexAttrs("Client", "Device", "Warehouse", "Quantity", "Date")
	cdIn.EditAttrs(cdIn.IndexAttrs()...)
	cdIn.NewAttrs(cdIn.IndexAttrs()...)

	cdOut := adm.AddResource(&db.ClientDeviceOut{}, &admin.Config{Menu: []string{"日常操作"}})
	cdOut.Meta(&admin.Meta{Name: "Client", Type: "select_one"})
	cdOut.Meta(&admin.Meta{Name: "Device", Type: "select_one"})
	cdOut.Meta(&admin.Meta{Name: "Warehouse", Type: "select_one", Collection: db.WarehouseCollection})
	cdOut.Scope(&admin.Scope{
		Default: true,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Preload("Device").Preload("Client")
		},
	})
	cdOut.IndexAttrs("Client", "Device", "Warehouse", "Quantity", "Date")
	cdOut.EditAttrs(cdOut.IndexAttrs()...)
	cdOut.NewAttrs(cdOut.EditAttrs()...)

	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"数据维护"}})
	device.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: db.DeviceCategories})

	deviceIn := adm.AddResource(&db.DeviceIn{}, &admin.Config{Menu: []string{"日常操作"}})
	deviceOut := adm.AddResource(&db.DeviceOut{}, &admin.Config{Menu: []string{"日常操作"}})

	var deviceOuts []*db.DeviceOut
	var inNumbers []string
	db.DB.Find(&deviceOuts).Pluck("number", &inNumbers)
	deviceIn.Meta(&admin.Meta{Name: "Number", Type: "select_one", Collection: inNumbers})

	var devices []*db.Device
	var outNumbers []string
	db.DB.Find(&devices).Where("available_amount > ?", 0).Pluck("number", &outNumbers)
	deviceOut.NewAttrs("-LendedAt")
	deviceOut.Meta(&admin.Meta{Name: "Number", Type: "select_one", Collection: outNumbers})

	consumableIn := adm.AddResource(&db.ConsumableIn{}, &admin.Config{Menu: []string{"日常操作"}})
	consumableIn.Meta(&admin.Meta{Name: "Name", Type: "string"})
	consumableIn.Meta(&admin.Meta{Name: "Code", Type: "string"})
	consumableIn.Meta(&admin.Meta{Name: "Count", Type: "int"})
	consumableIn.EditAttrs("Name", "Code", "Count")
	consumableIn.NewAttrs(consumableIn.EditAttrs()...)

	consumableOut := adm.AddResource(&db.ConsumableOut{}, &admin.Config{Menu: []string{"日常操作"}})
	consumableOut.Meta(&admin.Meta{Name: "Name", Type: "string"})
	consumableOut.Meta(&admin.Meta{Name: "Code", Type: "string"})
	consumableOut.Meta(&admin.Meta{Name: "Count", Type: "int"})
	consumableOut.EditAttrs("Name", "Code", "Count")
	consumableOut.NewAttrs(consumableOut.EditAttrs()...)

	adm.AddResource(&db.Client{}, &admin.Config{Menu: []string{"数据维护"}})
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
