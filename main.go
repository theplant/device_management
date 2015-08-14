package main

import (
	"log"
	"net/http"

	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/theplant/device_management/db"
)

func main() {
	adm := admin.New(&qor.Config{DB: &db.DB})

	pgr := adm.AddResource(&db.Device{}, &admin.Config{
		Name: "Github Repositories", Menu: []string{"Profile"}, PageCount: 100,
	})
	pgr.IndexAttrs("Login")

	device := adm.AddResource(&db.Device{}, &admin.Config{Menu: []string{"Device Management"}})
	device.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"自有设备", "消耗品", "客户设备"}})

	adm.MountTo("/admin", http.DefaultServeMux)

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
