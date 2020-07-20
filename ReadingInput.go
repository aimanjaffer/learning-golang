package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("you typed " + scanner.Text())
		if scanner.Text() == "/quit" {
			fmt.Println("See ya later")
			os.Exit(3)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
