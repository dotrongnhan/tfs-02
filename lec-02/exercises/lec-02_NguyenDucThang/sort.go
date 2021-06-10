package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice1 := createSlice(10)
	fmt.Println("Unsorted: ", slice1)
	bubbleSort(slice1)
	fmt.Println("After bubbleSort: ", slice1)
	fmt.Println("------------------------------------------------")
	slice2 := createSlice(10)
	fmt.Println("Unsorted: ", slice2)
	quickSort(slice2)
	fmt.Println("After quickSort: ", slice2)
}

func createSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) - 0
	}
	return slice
}

func bubbleSort(arr []int) {
	var (
		n      = len(arr)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func quickSort(arr2 []int) []int {
	if len(arr2) < 2 {
		return arr2
	}

	begin, end := 0, len(arr2)-1

	pivot := rand.Int() % len(arr2)

	arr2[pivot], arr2[end] = arr2[end], arr2[pivot]

	for i, _ := range arr2 {
		if arr2[i] < arr2[end] {
			arr2[begin], arr2[i] = arr2[i], arr2[begin]
			begin++
		}
	}
	arr2[begin], arr2[end] = arr2[end], arr2[begin]

	quickSort(arr2[:begin])
	quickSort(arr2[begin+1:])

	return arr2
}
