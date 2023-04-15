// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// Print each argument on a separate line
	for idx, arg := range os.Args {
		// combine idx and arg into a single string separated with a space
		s += sep + fmt.Sprintf("%d %s", idx, arg)
		sep = "\n"
	}
	fmt.Println(s)
}

//!-
