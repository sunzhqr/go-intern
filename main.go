package main

import (
	"fmt"
	"time"
)

var weekday string = time.Monday.String()

func init() {
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Printf("Today is %s", weekday)
}
