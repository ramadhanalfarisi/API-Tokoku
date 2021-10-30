package router

import (
	"net/http"
	"restapi-basic/controller"
	"restapi-basic/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() {
	mainRouter := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	mainRouter.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	secure := mainRouter.PathPrefix("/api").Subrouter()
	secure.Use(middleware.MiddlewareJWT)

	// User Endpoint
	mainRouter.HandleFunc("/register", controller.Register).Methods("POST")
	mainRouter.HandleFunc("/login", controller.Login).Methods("POST")

	// Category Endpoint
	secure.HandleFunc("/category", controller.InsertCategory).Methods("POST")
	secure.HandleFunc("/category", controller.SelectAllCategory).Methods("GET")
	secure.HandleFunc("/category", controller.UpdateCategory).Methods("PUT")
	secure.HandleFunc("/category/{id}", controller.SelectOneCategory).Methods("GET")
	secure.HandleFunc("/category/{id}", controller.DeleteCategory).Methods("DELETE")

	// Product Endpoint
	secure.HandleFunc("/product", controller.InsertProduct).Methods("POST")
	secure.HandleFunc("/product", controller.SelectAllProduct).Methods("GET")
	secure.HandleFunc("/product", controller.UpdateProduct).Methods("PUT")
	secure.HandleFunc("/product/{id}", controller.SelectOneModifier).Methods("GET")
	secure.HandleFunc("/product/{id}", controller.DeleteProduct).Methods("DELETE")

	// Modifier Endpoint
	secure.HandleFunc("/modifier", controller.InsertModifier).Methods("POST")
	secure.HandleFunc("/modifier", controller.SelectAllModifier).Methods("GET")
	secure.HandleFunc("/modifier", controller.UpdateModifier).Methods("PUT")
	secure.HandleFunc("/modifier/{id}", controller.SelectOneModifier).Methods("GET")
	secure.HandleFunc("/modifier/{id}", controller.DeleteModifier).Methods("DELETE")

	// Location Endpoint
	secure.HandleFunc("/location", controller.InsertLocation).Methods("POST")
	secure.HandleFunc("/location", controller.GetAllLocation).Methods("GET")
	secure.HandleFunc("/location", controller.UpdateLocation).Methods("PUT")
	secure.HandleFunc("/location/{id}", controller.GetOneLocation).Methods("GET")
	secure.HandleFunc("/location/{id}", controller.DeleteProduct).Methods("DELETE")

	// Menu Endpoint
	secure.HandleFunc("/menu", controller.SelectAllMenu).Methods("GET")

	// Product Endpoint
	http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(mainRouter))
}
