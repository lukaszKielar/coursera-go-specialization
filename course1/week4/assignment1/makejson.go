package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

func main() {
	var m = make(map[string]string)

	var name string
	var address string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name: ")
	name, _ = reader.ReadString('\n')
	m["name"] = strings.TrimSpace(name)

	fmt.Print("Enter an address: ")
	address, _ = reader.ReadString('\n')
	m["address"] = strings.TrimSpace(address)

	j, _ := json.Marshal(m)
	fmt.Printf("JSON: %s", j)
	fmt.Println()
}
