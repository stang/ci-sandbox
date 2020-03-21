package main

import (
	"fmt"
	"os"

	"github.com/stang/ci-sandbox/internal/lib"
)

var (
	version string
)

func main() {

	switch {
	case len(os.Args) == 2 && os.Args[1] == "--version":
		fmt.Printf("hello %s\n", version)
	case len(os.Args) == 2:
		fmt.Println(lib.Greetings(os.Args[1]))
	default:
		fmt.Println(`usage: hello <name>`)
	}

}
