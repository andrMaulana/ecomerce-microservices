package routes

import (
	"net/http"

	"github.com/andrMaulana/ecomerce-microservices/api/products/controllers"
	"github.com/gorilla/mux"
)

func RegisterProductsRoutes(router *mux.Router, productsController *controllers.ProductsController) {
	router.HandleFunc("/products", productsController.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", productsController.GetProductByID).Methods(http.MethodGet)
	router.HandleFunc("/products", productsController.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", productsController.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/products/{id}", productsController.DeleteProduct).Methods(http.MethodDelete)
}
