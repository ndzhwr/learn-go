package main

import "fmt"

func calculateStats(nums []int) (int, int, int) {
	var sum int
	var max int = nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] >= max {
			max = nums[i]
		}
	}

	avg := sum / len(nums)

	return sum, avg, max
}

func main(){
	sum, _, _ := calculateStats([]int{1, 2, 3, 5, 6})
	_ , avg ,_ := calculateStats([]int{1,2, 3, 5, 6})
	_ , _ , max := calculateStats([]int{1,2, 3, 5, 6})

	fmt.Printf("The sum is %d , The average is %d , The maximum is %d", sum,avg,max)


}