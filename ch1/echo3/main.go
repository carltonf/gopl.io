// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
// Review:
// Here is an example of output:
// $ go run main.go some   you know jij j llla jijiea a b c
// some you know jij j llla jijiea a b c
// echo1: 0.001706
// some you know jij j llla jijiea a b c
// echo3: 0.000240
// 
// After reading the section on benchmark, I found it would be not very useful to
// benchmark a program printing a lot to the stdout.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//!+
func main() {
	// count := 100
	beg_t := time.Now()
	echo1()
	fmt.Printf("echo1: %f\n", time.Since(beg_t).Seconds())

	beg_t = time.Now()
	echo3()
	fmt.Printf("echo3: %f\n", time.Since(beg_t).Seconds())
}

//!-
