package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	return nil
}
