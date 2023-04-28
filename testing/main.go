package main

import (
	"errors"
	"fmt"
)

func DiscountedPrice(price float64, discountRate float64) (float64, error) {
	switch {
	case discountRate == 0:
		return price, nil
	case discountRate < 0:
		return 0, errors.New("invalid discount rate")
	case discountRate > 1:
		return 0, errors.New("invalid discount rate")
	}
	discount := price * discountRate
	newPrice := price - discount
	return newPrice, nil

}

func main() {

	result, err := DiscountedPrice(1000, 0.6)
	fmt.Println(result, err)
}
