package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	ints := make([]int, 0, 3)
	fmt.Println("To exit the program type \"X\".")

	for {
		var input string
		fmt.Print("Enter an integer: ")
		fmt.Scanln(&input)

		if input == "X" {
			break
		}

		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not an integer. Exiting...")
			os.Exit(-1)
		}

		ints = append(ints, i)
	}

	slices.Sort(ints)
	fmt.Printf("Ints: %d\n", ints)
}
