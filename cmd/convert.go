package cmd

import (
	"log"
	"mojo/logic"

	"github.com/spf13/cobra"
)

func newConvertCommand() *cobra.Command {
	var from, to, contract, file string

	var convertCmd = &cobra.Command{
		Use:   "convert",
		Short: "Convert data from some file to another",
		Run: func(cmd *cobra.Command, args []string) {
			validator := logic.NewValidator()

			err := validator.ValidateConverterArgs(from, to, contract)
			if err != nil {
				log.Fatalln(err)
			}

		},
	}

	convertCmd.Flags().StringVarP(&from, "from", "fr", "", "from indicates the file to read the data")
	convertCmd.Flags().StringVarP(&to, "to", "t", "t", "to what kind of file the convertion is needed")
	convertCmd.Flags().StringVarP(&contract, "contract", "c", "", "the contract to be used during convertion")
	convertCmd.Flags().StringVarP(&file, "file", "f", "", "file to convert")

	convertCmd.MarkFlagRequired("from")
	convertCmd.MarkFlagRequired("to")
	convertCmd.MarkFlagRequired("contract")
	convertCmd.MarkFlagRequired("file")

	return convertCmd
}
