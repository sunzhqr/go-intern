package main

import "fmt"

type User struct {
	Level  int
	Name   string
	Age    int
	IsKeks struct {
		Keks bool
	}
	friends []uint8
}

func (u User) IsSigma() {
	if u.IsKeks.Keks == false {
		fmt.Printf("User %v has sigma\n", u.Name)
		return
	}
	fmt.Printf("Obviosly, %v is gay\n", u.Name)
}

func main() {
	user1 := User{1, "Abylay", 19, struct{ Keks bool }{true}, []uint8{1, 2, 3}}
	user2 := User{
		Level:   100,
		Name:    "Sanzhar Unchained",
		Age:     19,
		IsKeks:  struct{ Keks bool }{Keks: false},
		friends: []uint8{1, 2, 3, 5, 6, 7, 8, 9, 10},
	}
	user1.IsSigma()
	user2.IsSigma()
}
