package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name")
	inputs, _ := fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s! nice to meet you. %d\n", name, inputs)
}
