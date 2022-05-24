package models

import (
	"localhost/alura-GoWebProducts/src/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()
	productsDB, err := db.Query("Select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	products := []Product{}
	for productsDB.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsDB.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product := Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity}
		products = append(products, product)
	}

	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	createProduct, err := db.Prepare("Insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	createProduct.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("Delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	productDB, err := db.Query("Select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	var product Product
	for productDB.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDB.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product = Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity}
	}

	defer db.Close()
	return product
}

func UpdateProduct(id, quantity int, name, description string, price float64) {
	db := db.ConnectDB()

	updateProduct, err := db.Prepare("Update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
