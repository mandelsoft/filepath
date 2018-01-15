package main

import (
	"fmt"
	"os"

	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("two argument expected\n")
	} else {
		fmt.Printf("%s | %s -> %s\n", os.Args[1], os.Args[2], filepath.Join(os.Args[1], os.Args[2]))
	}
}
