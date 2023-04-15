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
	for _, arg := range os.Args {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(s)
}

//!-
