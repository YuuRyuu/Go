// Convention to add a description of the package here.
// hello prints a quote from a function from the quote package.
// It also prints the command line arguments.
package main

// This is how you import packages.
import (
	"fmt"
	"os"

	"rsc.io/quote/v4"
)

// This is the main function
func main() {
	fmt.Println(quote.Go())

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
