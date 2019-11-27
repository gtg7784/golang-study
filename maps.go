package main

import "fmt"

func main(){
	presAge := make(map[string] int)

	presAge["Go Tae Geon"] = 17

	fmt.Println(presAge["Go Tae Geon"])

	presAge["Kim hyun Woo"] = 18

	fmt.Println([presAge])

	delete(presAge, "Go Tae Geon")
	fmt.Println(len(presAge))
}