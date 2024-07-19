package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/andrMaulana/ecomerce-microservices/api/products/controllers"
	"github.com/andrMaulana/ecomerce-microservices/api/products/middleware"
	"github.com/andrMaulana/ecomerce-microservices/api/products/models"
	"github.com/andrMaulana/ecomerce-microservices/api/products/routes"
)

func main() {
	// Connect to PostgreSQL database
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		"localhost", 5432, "postgres", "postgres", "microservice_ecommerce_db")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate database schema
	db.AutoMigrate(&models.Product{})

	// Create products controller
	productsController := controllers.NewProductsController(db)

	// Create router
	router := mux.NewRouter()
	// Register middleware
	router.Use(middleware.LogRequestMiddleware)
	router.Use(middleware.AuthMiddleware)

	// Register products routes
	routes.RegisterProductsRoutes(router, productsController)
	// Start API server
	log.Println("Starting API server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
