package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configure s3 command
func s3Command() *cobra.Command {
	var awsProfile, bucket, filename string
	var recursive bool

	var s3Cmd = &cobra.Command{
		Use:   "s3",
		Short: "Using s3 as source",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("S3 command executed, with flag profile:", awsProfile)
		},
	}

	s3Cmd.Flags().StringVarP(&awsProfile, "profile", "p", "", "profile to use in aws commands")
	s3Cmd.Flags().StringVarP(&bucket, "bucket", "b", "", "bucket to query data in s3 service")
	s3Cmd.Flags().StringVarP(&filename, "file", "f", "", "file ")
	s3Cmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "if true will download every folder in")

	s3Cmd.MarkFlagRequired("profile")
	s3Cmd.MarkFlagRequired("bucket")

	return s3Cmd
}
