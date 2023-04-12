package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your full names: ")
	scanner.Scan()
	in_string := scanner.Text()
	in_string  =  strings.TrimSpace(in_string)
	fmt.Printf("\n");
	fmt.Printf("You're called %s", in_string)

}
