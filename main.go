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

func main() {

}
