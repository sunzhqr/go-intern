package main

import "fmt"

//	func sum(nums ...int) {
//		fmt.Print(nums, " ")
//		total := 0
//		for _, num := range nums {
//			total += num
//		}
//		fmt.Println(total)
//	}
func main() {
	func() {
		fmt.Print("Hello ")
	}()
	world := func() {
		fmt.Println("World")
	}
	world()
}
