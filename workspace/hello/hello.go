package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse" // if we modify this file, go work use /example/hello
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
