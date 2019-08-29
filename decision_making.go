package main

import "fmt"

func main() {
	age := 15

	if age >= 16 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not an adult")
	}

	if age >= 16 {
		fmt.Println("in school")
	} else if age >= 18 {
		fmt.Println("in collage")
	} else {
		fmt.Println("probably dead")
	}

	switch age {
	case 16:
		fmt.Println("Go Drive")
	case 18:
		fmt.Println("Go Vote")
	default:
		fmt.Println("Go Have Fun")
	}
}
