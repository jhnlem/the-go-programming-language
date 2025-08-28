// Modify the echo program to print the index and value
// of each arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for id, arg := range os.Args[:] {
		fmt.Println(id, " = ", arg)
	}
}
