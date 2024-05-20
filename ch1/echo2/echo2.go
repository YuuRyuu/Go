// Echo 2 is the same as echo1 but it iterates over a range of values instead
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// The blank identifier is used to ignore the index value of the range
	// arg is the value of each element in the slice os.Args
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
