package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("C:\\Users\\nguye\\OneDrive\\Desktop\\Test Golang\\hw/testCreate.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var words string
	fmt.Println("Enter the string of words you want to write:")
	fmt.Scanln(&words)
	l, err := f.WriteString(words)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
