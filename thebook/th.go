package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	// SCANNER FUNCTION

	// util vars
	var (
		names [10]string
		index int  = 1
		cont  bool = true
	)

	for cont {
		var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)
		fmt.Println("What is the first name ?")
		scanner.Scan()
		names[index] = strings.TrimSpace(scanner.Text())
		index += 1

		var contd byte
		fmt.Println("Add another? [y/n]")
		fmt.Scan(&contd)
		if contd == 'y' {
			cont = true
		} else {
			cont = false
		}
	}
	fmt.Print(names)

	const (
		MATH_PI float32 = math.Pi
		MY_NAME string  = "regis"
	)

	fmt.Printf("The declared variables are : %f and %s" , MATH_PI , MY_NAME);
}