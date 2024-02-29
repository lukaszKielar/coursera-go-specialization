package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(numbers *[]int) {
	for i := 0; i < len(*numbers)-1; i++ {
		for j := 0; j < len(*numbers)-1; j++ {
			if (*numbers)[j] > (*numbers)[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func Swap(numbers *[]int, index int) {
	(*numbers)[index], (*numbers)[index+1] = (*numbers)[index+1], (*numbers)[index]
}

func main() {
	fmt.Println("Enter up to 10 integers separated with a space, e.g. 12 132 1")

	var elems []string
	numbers := make([]int, 0, 10)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	elems = strings.Split(strings.TrimSpace(input), " ")

	if len(elems) > 10 {
		panic("More than 10 integers provided")
	}

	for _, el := range elems {
		var n int
		n, err := strconv.Atoi(el)
		if err != nil {
			panic(fmt.Sprintf("Value %s is not an integer", el))
		}
		numbers = append(numbers, n)
	}

	BubbleSort(&numbers)

	fmt.Println("Bubble sorted integers:", numbers)
}
