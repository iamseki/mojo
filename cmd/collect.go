package cmd

import (
	"fmt"
	"mojo/mmi"

	"github.com/spf13/cobra"
)

func newCollectCommand() *cobra.Command {
	var collectCmd = &cobra.Command{
		Use:   "collect",
		Short: "Collect data from some source",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`Must provide one of available sources:
			s3
			b3
			anbima
	In the meanwhile, here's some fact:`, mmi.GetRandomSentence())
		},
	}

	// Adding child commands
	collectCmd.AddCommand(
		s3Command(),
	)

	return collectCmd
}
