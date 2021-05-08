package cmd

import (
	"fmt"
	"os"

	"mojo/mmi"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mojo",
	Short: "Mojo is a helper cli, perfect for you lazy programmer",
	Long: `Mojo is a general purpose helper, perfect for you lazy programmer, written in go, a simple and beautiful language. 
                Complete documentation is available at http://TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mojo --help, in the meanwhile here's a fact:")
		fmt.Println(mmi.GetRandomSentence())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
