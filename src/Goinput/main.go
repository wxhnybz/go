package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s\n", name)
}
