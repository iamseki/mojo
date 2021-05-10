package main

import (
	"log"
	"mojo/cmd"
	"mojo/helper"
)

func init() {
	err := helper.SetupMojoBrain()
	if err != nil {
		log.Printf(`
		Something went wrong on Mojo's brain setup: %v\n
		but dont worry, mojo still working but without some memory cache features\n`, err)
	}
}

func main() {
	cmd.Execute()
}
