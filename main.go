package main

import "fmt"

func main() {
	arr := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	slc := []int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10}
	fmt.Println(len(arr), cap(arr))
	fmt.Printf("array: %T, slice: %T\n", arr, slc)
	fmt.Printf("array[1]: %d, slice[1]: %d\n", arr[1], slc[1])

	// arr = append(arr, 123)  // wrong
	slc = append(slc, 321)

	slc = append(slc[len(slc)-1:], slc[:len(slc)-1]...)
	fmt.Println(slc)

}
