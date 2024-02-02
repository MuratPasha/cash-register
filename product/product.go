package product

import "fmt"

// Item is a struct that represents a product
type Item struct {
	Name     string
	Price    float64
	Discount float64
}

// Description returns a string with the name and price of the item
func (i Item) Description() string {
	if i.Discount != 0 {
		return fmt.Sprintf("%s - %.2f (%.2f with discount)", i.Name, i.Price, i.Price*i.Discount)
	}
	return fmt.Sprintf("%s - %.2f", i.Name, i.Price)
}

// CalculatePrice returns the price of an item after applying the discount
func CalculatePrice(item Item) float64 {
	if item.Discount != 0 {
		return item.Price * item.Discount
	}
	return item.Price
}

// TotalPrice returns the total price of a list of items
func TotalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price
	}
	return total
}

// Descripable is an interface that requires a Description method
type Descripable interface {
	Description() string
}
