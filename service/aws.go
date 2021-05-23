package service

import (
	"bytes"
	"errors"
	"log"
	"mojo/helper"
	"os/exec"
)

type AWS interface {
	DownloadFileFromS3(DownloadS3Options) error
	ListS3FoldersFromObject(string) []string
}

type awsCommandService struct {
	profile string
}

func NewAWSCommand(profile string) AWS {
	return &awsCommandService{
		profile: profile,
	}
}

// ListS3FoldersFromObject
// list all available folders/objects from a given folder
// folder must be in the format as follow: s3://some-bucket/some-folder/
// this works either: s3://some-bucket/
func (aws *awsCommandService) ListS3FoldersFromObject(folder string) []string {
	cmd := exec.Command("aws",
		"s3",
		"ls",
		folder,
		"--profile",
		aws.profile,
		"--page-size",
		"1000",
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatalln(err)
	}

	folders := helper.SliceFromAWSUnformattedString(out.String())
	return folders
}

type DownloadS3Options struct {
	filepath     string
	targetFolder string
}

// DownloadFileFromS3
func (aws *awsCommandService) DownloadFileFromS3(o DownloadS3Options) error {
	if o.filepath == "" {
		return errors.New("filepath must be provided to download S3 file")
	}

	if o.targetFolder == "" {
		return errors.New("targetFolder must be provided to download S3 file")
	}

	cmd := exec.Command(
		"aws",
		"s3",
		"cp",
		o.filepath,
		o.targetFolder,
		"--profile",
		aws.profile,
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
