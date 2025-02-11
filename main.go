package main

import (
	"errors"
	"time"
)

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

func (p Product) GetPrice() float64 {
	return p.Price
}

func (cart *ShoppingCart) AddItem(product Product) {
	cart.Items = append(cart.Items, product)
}

func (cart *ShoppingCart) RemoverItem(productID int) error {
	for i, item := range cart.Items {
		if item.ID == productID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("produto não encontrado no carrinho")
}

func main() {

}
