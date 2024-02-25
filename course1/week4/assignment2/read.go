package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX_LEN = 20

type Person struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter a name to the file: ")
	reader := bufio.NewReader(os.Stdin)
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var people []Person

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), " ", 2)
		fname := line[0]
		lname := line[1]

		if len(fname) > MAX_LEN {
			fname = fname[:MAX_LEN]
		}
		if len(lname) > MAX_LEN {
			lname = lname[:MAX_LEN]
		}

		p := Person{fname, lname}
		people = append(people, p)
	}

	for _, e := range people {
		fmt.Printf("First name: %s, Last name: %s\n", e.fname, e.lname)
	}
}
