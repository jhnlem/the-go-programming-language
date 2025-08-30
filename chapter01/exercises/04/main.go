// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin of from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("", os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]int, filenames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		j, M := 0, 0
		for i := 0; i < len(filenames[input.Text()]); i++ {
			if j < len(arg) {
				if filenames[input.Text()][i] == arg[j] {
					j++
				} else {
					j = 0
				}
				if j > M {
					M = j
				}
			}
		}
		if M < len(arg) {
			filenames[input.Text()] += arg + " "
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
