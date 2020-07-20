package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s! nice to meet you", name)
}
