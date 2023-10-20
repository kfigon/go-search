package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	docs := map[FileName]string {
		"fileA":"The car is driven on the road",
		"fileB": "The truck is driven on the highway",
	}

	s := NewService(docs)
	fmt.Println("car ->", s.Results("car"))
	fmt.Println("truck ->", s.Results("truck"))
	fmt.Println("the ->", s.Results("the"))
	
	// res := Parse(docs)
	// for r, v := range res {
	// 	fmt.Println(r,"->",v)
	// }
}
