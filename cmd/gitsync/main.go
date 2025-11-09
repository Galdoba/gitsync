package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/gitsync/internal/application/command"
)

func main() {
	cmd, err := command.GitSync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "initialization failed: %v", err)
		os.Exit(1)
	}
}
