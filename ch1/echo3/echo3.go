package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // simpler and more efficient approach for large data

	fmt.Println(os.Args[1:]) // prints the elements without format (useful for quick debugging)
}
