package product

import "fmt"

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

func (i Item) Description() string {
	if i.Discount != 0 {
		return fmt.Sprintf("%s - %.2f (%.2f with discount)", i.Name, i.Price, i.Price*i.Discount)
	}
	return fmt.Sprintf("%s - %.2f", i.Name, i.Price)
}

func CalculatePrice(item Item) float64 {
	if item.Discount != 0 {
		return item.Price * item.Discount
	}
	return item.Price
}

func TotalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price
	}
	return total
}

type Descripable interface {
	Description() string
}
