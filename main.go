package main

import "fmt"

//	func intSeq() func() int {
//		i := 0
//		return func() int {
//			i++
//			return i
//		}
//	}
func main() {
	s := "Sanzhar"
	defer fmt.Println(s)
	defer func() {
		fmt.Println(s)
	}()
	s = "Abylay"
}
