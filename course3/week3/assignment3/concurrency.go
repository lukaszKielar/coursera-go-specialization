package main

// Write a program to sort an array of integers.
// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list.

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

const (
	numGoroutines = 4
	arrSize       = 10
)

var (
	arr []int
	wg  sync.WaitGroup
)

func main() {
	c := make(chan []int)

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go sortArr(i, &c, &wg)
	}

	wg.Wait()

	for i := 0; i < numGoroutines; i++ {
		partialArr := <-c
		arr = append(arr, partialArr...)
	}

	sort.Ints(arr)
	fmt.Println()
	fmt.Println("Final sorted array:", arr)
}

func generateRandomArray() []int {
	arr := make([]int, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func sortArr(i int, c *chan []int, wg *sync.WaitGroup) {
	arr := generateRandomArray()
	fmt.Printf("Sorting array: %d by goroutine %d", arr, i)
	fmt.Println()

	sort.Ints(arr)
	fmt.Printf("Sorted array: %d by goroutine %d", arr, i)
	fmt.Println()

	wg.Done()

	*c <- arr
}
