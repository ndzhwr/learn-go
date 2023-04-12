package main

import "fmt"

func main(){
	var name string ;
	fmt.Print("What is your name: ");
	fmt.Scanln(&name)
	for _ , c :=  range name {
		fmt.Println(string(c))
	}

	fmt.Println("The length of your name: " );
	fmt.Println(len(name))
}