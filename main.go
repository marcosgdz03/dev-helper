// Package main is the entry point for dev-helper.
package main

import (
	"fmt"
	"os"

	"github.com/marcosgdz03/dev-helper/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
