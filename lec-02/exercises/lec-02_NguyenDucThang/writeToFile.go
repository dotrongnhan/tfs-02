package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("C:\\Users\\nguye\\OneDrive\\Desktop\\Test Golang\\lec-02_NguyenDucThang/testCreate.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter the string of words you want to write:")
	r := bufio.NewReader(os.Stdin)
	words, _ := r.ReadString('\n')
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
