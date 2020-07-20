package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name")
	inputs, _ := fmt.Scanf("%q", &name)
	fmt.Printf("Hello %s! nice to meet you. %d\n", name, inputs)
}
