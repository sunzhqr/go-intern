package main

import "fmt"

func main() {
	s := "lorem"

	// loop through bytes(uint8)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v (%T): %v (%T):\n", i, i, s[i], s[i])
	}

	// loop through runes(int32)
	for i, r := range s {
		fmt.Printf("%v (%T): %v (%T):\n", i, i, r, r)
	}

	fmt.Println(s + string('A'+32) + string('0'+9))
}
