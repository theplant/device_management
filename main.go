package main

import (
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"github.com/theplant/device_management/db"
	"log"
	"net/http"
)

func main() {
	adm := admin.New(&qor.Config{DB: &db.DB})

	pgr := adm.AddResource(&db.Device{}, &admin.Config{
		Name: "Github Repositories", Menu: []string{"Profile"}, PageCount: 100,
	})
	pgr.IndexAttrs("Login")

	adm.MountTo("/admin", http.DefaultServeMux)

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
