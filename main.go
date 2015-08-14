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
    
   
	deviceIn := adm.AddResource(&db.DeviceIn{}, &admin.Config{Menu: []string{"Device Management"}})
	deviceOut := adm.AddResource(&db.DeviceOut{}, &admin.Config{Menu: []string{"Device Management"}})
	
	var deviceOuts []*db.DeviceOut
	var inNumbers []string
	db.DB.Find(&deviceOuts).Pluck("number", &inNumbers)
	deviceIn.Meta(&admin.Meta{Name: "Number", Type: "select_one", Collection: inNumbers})

	var devices []*db.Device
	var outNumbers []string
	db.DB.Find(&devices).Where("available_amount > ?", 0).Pluck("number", &outNumbers)
	deviceOut.NewAttrs("-LendedAt")
	deviceOut.Meta(&admin.Meta{Name: "Number", Type: "select_one", Collection: outNumbers})
	




	adm.MountTo("/admin", http.DefaultServeMux)

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
