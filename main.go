package main

import "time"

type Category string

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    Category
}

type ShoppingCart struct {
	Items []Product
}

type Order struct {
	Cart  ShoppingCart
	Order time.Time
	Total float64
}

// interfaces
type ItemPrice interface {
	GetPrice() float64
}

type Catalog interface {
	AddProduct(product Product)
	GetProductByID(id int) (Product, error)
	ListProductByCategory(category Category) []Product
}

func main() {

}
