package cli

import (
	"github.com/spf13/cobra"
)

var jsonOut bool

func Execute() error {
	root := &cobra.Command{
		Use:           "books",
		Short:         "Book lookup and ISBN resolution for Dungeon Books",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().BoolVar(&jsonOut, "json", false, "output JSON")
	root.AddCommand(bookCmd())
	root.AddCommand(trendingCmd())
	root.AddCommand(resolveCmd())
	root.AddCommand(mcpCmd())
	return root.Execute()
}
