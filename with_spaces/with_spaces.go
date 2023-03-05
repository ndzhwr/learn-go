package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your full names")
	scanner.Scan()
	in_string := scanner.Text()
	fmt.Printf("You're called %s", in_string)

}
