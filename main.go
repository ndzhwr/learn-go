package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string = "Dax"
	var age, class int = 32, 2
	fmt.Println("Hello World! " + name)
	fmt.Println(name+" is ", age, " and is in Y", class)

	// shorthand for declaring variable
	sch := "RCA"
	fmt.Println(sch)

	// constants declatation
	const pi float32 = 3.14
	fmt.Println("The value of PI is ", pi)

	// exponentail types
	fmt.Println(3e45)

	var typePI reflect.Type = reflect.TypeOf(pi);
	fmt.Println(typePI.String() == "float32");
	fmt.Println(typePI.String() == "float64")
}
