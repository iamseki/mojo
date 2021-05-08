package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Mojo",
	Long:  `All software has versions. This is Mojo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mojo v0.1")
	},
}
