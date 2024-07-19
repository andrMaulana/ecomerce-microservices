package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/andrMaulana/ecomerce-microservices/api/products/models"
)

type ProductsController struct {
	DB *gorm.DB
}

func NewProductsController(db *gorm.DB) *ProductsController {
	return &ProductsController{
		DB: db,
	}
}

func (c *ProductsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	c.DB.Find(&products)

	json.NewEncoder(w).Encode(products)
}

func (c *ProductsController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var product models.Product
	c.DB.First(&product, productID)

	json.NewEncoder(w).Encode(product)
}

func (c *ProductsController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	c.DB.Create(&product)

	json.NewEncoder(w).Encode(product)
}

func (c *ProductsController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var product models.Product
	c.DB.First(&product, productID)

	json.NewDecoder(r.Body).Decode(&product)

	c.DB.Save(&product)

	json.NewEncoder(w).Encode(product)
}
func (c *ProductsController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var product models.Product
	c.DB.First(&product, productID)

	c.DB.Delete(&product)

	json.NewEncoder(w).Encode(map[string]bool{"message": true})
}
