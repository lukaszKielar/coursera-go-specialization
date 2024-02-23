package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fstr string

	fmt.Printf("Enter a floating point number to truncate: ")
	fmt.Scan(&fstr)

	f, err := strconv.ParseFloat(fstr, 64)
	if err != nil {
		fmt.Printf("Entered value '%s' cannot be parsed to float.", fstr)
		fmt.Println()
		fmt.Printf("Error: %s", err)
		fmt.Println()
		return
	}

	i := int64(f)
	fmt.Printf("Truncated number: %d", i)
	fmt.Println()
}
