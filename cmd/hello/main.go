package main

import (
	"fmt"
	"os"
)

var (
	version string
)

func main() {
	switch {
	case len(os.Args) == 2 && os.Args[1] == "version":
		fmt.Printf("hello %s\n", version)
	default:
		fmt.Println("Hello world!")
	}
}
