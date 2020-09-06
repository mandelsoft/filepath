package main

import (
	"fmt"
	"os"

	"github.com/mandelsoft/filepath/pkg/filepath"
)

func main() {
	if len(os.Args) < 2 {
		handle(".")
	} else {
		for _, p := range os.Args[1:] {
			handle(p)
		}
	}
}

func handle(p string) {
	a, err := filepath.EvalSymlinks(p)
	if err != nil {
		fmt.Printf("%s: %s\n", p, err)
	} else {
		fmt.Printf("%s -> %s\n", p, a)
	}
}
