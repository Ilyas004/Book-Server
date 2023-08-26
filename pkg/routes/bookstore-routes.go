package routes

import (
	"main/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(route *mux.Router) {
	route.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	route.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	route.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	route.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	route.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
