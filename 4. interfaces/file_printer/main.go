package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name must be passed as an argument!")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file =>", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
