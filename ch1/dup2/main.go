// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
type line_info struct {
	count int
	locations []string
}

func main() {
	counts := make(map[string]*line_info)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// sort by lines before printing
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, line := range keys {
		info := counts[line]
		fmt.Printf("%d\t%s\t%s\n", info.count, line, info.locations)
	}
	// for line, info := range counts {
	// 	if info.count > 1 {
	// 		fmt.Printf("%d\t%s\t%s\n", info.count, line, info.locations)
	// 	}
	// }
}

func countLines(f *os.File, counts map[string]*line_info) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &line_info{
				count: 0,
				locations: []string{},
			}
		}
		counts[input.Text()].count++
		loc_slice := counts[input.Text()].locations
		if len(loc_slice) == 0 || loc_slice[len(loc_slice)-1] != f.Name() {
			counts[input.Text()].locations = append(loc_slice, f.Name())

		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
