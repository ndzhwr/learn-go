package main

import (
	"fmt"
	"sync"
	"time"
)

var ii int64 ;

func print3(from string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ii++ ;
		fmt.Println(ii);
		fmt.Println(from, " : ", i)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	go print3("Routine 1", &wg)
	go print3("Routine 2", &wg)
	time.Sleep(time.Second)
	go print3("Routine 3", &wg)

	wg.Wait()

	// go func(name string) {
	// 	fmt.Println(name)
	// }("regis")

}
