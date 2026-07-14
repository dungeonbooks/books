package main

import (
	"fmt"
	"os"

	"github.com/dungeonbooks/books/internal/cli"
	"github.com/dungeonbooks/books/internal/clierr"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "books:", err)
		os.Exit(clierr.Code(err))
	}
}
