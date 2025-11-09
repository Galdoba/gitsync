package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Galdoba/gitsync/internal/application/command"
)

func main() {
	program, err := command.GitSync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "initialization failed: %v", err)
		os.Exit(1)
	}
	if err := program.Run(context.TODO(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "gitsync error: %v\n", err)
		os.Exit(1)
	}

}
