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

	n, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid integer.\n", os.Args[1])
		return
	}

	if n <= 0 {
		fmt.Println("Error: Please enter a positive integer.")
		return
	}

	var sum = (n * (n + 1)) / 2

	fmt.Printf("The sum of the first %d natural numbers is: %d\n", n, sum)
}
