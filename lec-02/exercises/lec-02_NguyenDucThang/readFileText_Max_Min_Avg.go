package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile() {
	file, err := os.Open("C:\\Users\\nguye\\OneDrive\\Desktop\\Test Golang\\lec-02_NguyenDucThang/Test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := scanner.Text()
		arr := strings.Fields(a)
		arr2 := []int{}
		for _, i := range arr {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			arr2 = append(arr2, j)
		}
		fmt.Println(arr2)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		max_min_avg(arr2)
		checkIsPrime(arr2)
		checkExist(arr2)

	}
}
func max_min_avg(arr []int) {
	max := arr[0]
	min := arr[0]
	sum := 0
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
		sum += arr[i]
	}
	avg := sum / len(arr)
	fmt.Printf("\nMax element is : %d", max)
	fmt.Printf("\nMin element is : %d", min)
	fmt.Printf("\nAverage element is : %d", avg)
	fmt.Println()
	fmt.Println("--------------------------------")
}
func checkIsPrime(arr []int) {
	check := false
	for i := 0; i < len(arr); i++ {
		x := int(math.Sqrt(float64(arr[i])))
		for j := 2; j <= x; j++ {
			if arr[i]%j == 0 {
				check = false
			} else {
				check = true
			}
		}
	}
	if check == true {
		fmt.Println("There are prime numbers")
	} else {
		fmt.Println("There are no prime numbers")
	}
}
func checkExist(arr []int) {
	var num int
	for {
		fmt.Println("Enter the number to check: ")
		fmt.Scanln(&num)
		check := false
		for i := 0; i < len(arr); i++ {
			if arr[i] == num {
				check = true
				break
			}
		}
		if check == true {
			fmt.Println("Exist!")
		} else {
			fmt.Println("Does not exist!!")
		}

		fmt.Println("Do you want to check again?")
		fmt.Println("1. Enter 1 to continue checking.\n2.Enter 0 to finish!")
		var cont int
		fmt.Scanln(&cont)
		if cont == 0 {
			break
		}
	}
}
func main() {
	readFile()
}
