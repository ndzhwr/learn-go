package main

import (
	"fmt"
)

func printRange(nums ...int) ([]int) {
	params := make([]int,0)
	var i int = 0
	for _, num := range nums {
		if i < 10 {
			params = append(params,num);
			i++
		}
		fmt.Printf("%d\t", num)
	}

	return params ;
}

func main() {
	var  params []int = printRange(1, 2, 3, 4, 5)
	fmt.Print(params)

}