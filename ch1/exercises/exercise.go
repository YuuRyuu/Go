package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Exercise 1.1
	// This displays the path of the command that was run and the arguments.
	fmt.Println(strings.Join(os.Args, " "))

	// Exercise 1.2
	// This displays the index and value of each argument.
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}

}
