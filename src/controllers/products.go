package controllers

import (
	"localhost/alura-GoWebProducts/src/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

var path, _ = os.Getwd()
var templates = template.Must(template.ParseGlob(filepath.Join(path, "src", "templates", "*.html")))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro on converting price:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro on converting quantity:", err)
		}

		models.CreateProduct(name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	product := models.EditProduct(id)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro on converting id:", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro on converting price:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro on converting quantity:", err)
		}

		models.UpdateProduct(convertedId, convertedQuantity, name, description, convertedPrice)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
