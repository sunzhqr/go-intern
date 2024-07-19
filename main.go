package main

import "fmt"

func main() {
	i := 5
	switch {
	case i < 5:
		fmt.Println("i is less than 5")
		fallthrough
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 15:
		fmt.Println("i is less than 15")
		fallthrough
	default:
		fmt.Printf("i is equal to %d\n", i)
	}
}
