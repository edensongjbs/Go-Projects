package main

import (
	"fmt"
	"os"
)

func main() {
	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error :", err)
	}
	bs := make([]byte, 100000)
	_, err = f.Read(bs)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println(string(bs))
}