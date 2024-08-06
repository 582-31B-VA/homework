package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		name  string
		birth int
	)
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your year of birth: ")
	fmt.Scanln(&birth)
	year := time.Now().Year()
	age := year - birth
	fmt.Printf("Hello, %v! You are currently %v years old.", name, age)
}
