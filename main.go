package main

import (
	"fmt"
	"os"
	"xversioner/file"
	"xversioner/utils"
)

func main() {
	fmt.Println("xVersioner Commands (CLI)")

	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "help" {
		fmt.Println(utils.Help())
	}

	version := file.GetVersion("app.properties")

	fmt.Println("App Version:", version[1])
	file.IncrementVersion(version[1], ".")

}
