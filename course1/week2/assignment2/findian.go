package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string

	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = strings.ToLower(scanner.Text())

	if strings.HasPrefix(line, "i") && strings.HasSuffix(line, "n") && strings.Contains(line, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
