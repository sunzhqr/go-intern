package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	sum(1, 2, 3)
	nums := []int{7, 8, 9}
	sum(nums...)
}
