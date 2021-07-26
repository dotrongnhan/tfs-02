package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(10)
	fmt.Println("Unsorted: ", slice)
	fmt.Println("After mergeSort: ", mergeSort(slice))
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) - 0
	}
	return slice
}

func mergeSort(arr []int) []int {
	var num = len(arr)

	if num == 1 {
		return arr
	}

	mid := int(num / 2)
	var (
		left  = make([]int, mid)
		right = make([]int, num-mid)
	)
	for i := 0; i < num; i++ {
		if i < mid {
			left[i] = arr[i]
		} else {
			right[i-mid] = arr[i]
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}
