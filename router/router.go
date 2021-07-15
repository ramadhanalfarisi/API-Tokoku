package router

import (
	"net/http"
	"restapi-basic/controller"
	"github.com/gorilla/mux"
)

func Router() {
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/category",controller.InsertCategory).Methods("POST")
	mainRouter.HandleFunc("/category",controller.SelectAllCategory).Methods("GET")
	mainRouter.HandleFunc("/category",controller.UpdateCategory).Methods("PUT")
	mainRouter.HandleFunc("/category/{id}",controller.SelectOneCategory).Methods("GET")
	mainRouter.HandleFunc("/category/{id}",controller.DeleteCategory).Methods("DELETE")
	mainRouter.HandleFunc("/menu",controller.SelectAllMenu).Methods("GET")
	http.ListenAndServe(":8000", mainRouter)

}