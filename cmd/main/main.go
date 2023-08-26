package main

import (
	"fmt"
	"log"
	"main/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Start server at port; :9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
