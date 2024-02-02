package main

import (
	"fmt"
	"temp/product"
)

func main() {
	product1 := product.Item{Name: "Laptop", Price: 1000, Discount: 0.10}
	product2 := product.Item{Name: "Mouse", Price: 10, Discount: 0.05}
	product3 := product.Item{Name: "Keyboard", Price: 20, Discount: 0.10}
	products := []product.Item{product1, product2, product3}

	fmt.Println(product.TotalPrice(products))
	fmt.Println(product.CalculatePrice(product1))
	fmt.Println(product1.Description())

}
