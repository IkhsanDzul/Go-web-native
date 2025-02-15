package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//Homepage
	http.HandleFunc("/", homecontroller.Home)

	//catgories
	http.HandleFunc("/categories", categorycontroller.IndexCategory)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//products

	log.Println("Server running")
	http.ListenAndServe(":8080", nil)

}