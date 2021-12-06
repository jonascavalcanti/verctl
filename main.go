package main

import (
	"fmt"
	"os"
	"xversioner/utils"
)

func main() {
	fmt.Println("xVersioner Commands (CLI)")

	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "help" {
		fmt.Println(utils.Help())
	}
}
