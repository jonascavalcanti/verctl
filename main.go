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

	oldVersion, b := file.GetVersion("./app.properties")
	fmt.Println("App Version:", oldVersion)

	newVersion := file.IncrementVersion(oldVersion, ".", "minor")
	fmt.Println("New Version:", newVersion)

	file.WriteVerionOnFile(oldVersion, newVersion, b)

}
