package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <positive_integer>")
		fmt.Println("Example: go run main.go 10")
		return
	}

	n, err := strconv.ParseFloat(os.Args[1], 0)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid integer.\n", os.Args[1])
		return
	}

	if n <= 0 {
		fmt.Println("Error: Please enter a positive integer.")
		return
	}
}
