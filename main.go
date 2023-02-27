package main

import (
	"fmt"
	"reflect"
	"time"
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

	// for loops
	i := 0 ;
	for i <= 9 {
		fmt.Printf("Loop 1 is on %d\n",i)
		i += 1
	}

	// real for loops
	for j := 0 ; j<= 20 ; j++ {
		fmt.Printf("Loop 2 is on %d\n", j)
	}

	// endless loop
	// for {
	// 	fmt.Println("And endless loop i guess \n")
	// }
	today := time.Now().Weekday();
	switch today {
	case time.Saturday :
	case time.Sunday :
		println("We're in weekend")
	default :
		println("We're in weekdays")
	}


	// arrays
	var a [12]int ;
	a[0] = 10
	fmt.Println(a)


	// slices
	myslice := make([]int ,5)
	myslice[0] = 45;
	for x := 0; x < len(myslice) ; x++ {
		myslice[x] = x + 11
		}
	theSLice := append(myslice, 4)
	fmt.Println(theSLice)
	fmt.Println(theSLice[1:3])
}

