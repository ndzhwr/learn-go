package main

import "fmt"



func main() {
	user := map[string]any{"name": "regis", "age": 0}
	fmt.Print("What is your name : ")
	var name string
	fmt.Scanln(&name)
	user["name"] = name

	fmt.Print("\n")

	fmt.Print("How old are you : ")
	var age int
	fmt.Scanln(&age)
	user["age"] = age

	thename := user["name"]

	fmt.Printf("\n\n\n Thankk you for your data\n\n\n\n==============================\n\n\n")

	// delete an element from the  map
	delete(user, "name")
	fmt.Print(user, "Map size = ", len(user), thename)

}
