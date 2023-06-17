package main

import (
	"fmt"
	"os"
	"github.com/jonascavalcantineto/versctl/help"
	model "github.com/jonascavalcantineto/versctl/model"
)

func main() {

	if len(os.Args) < 2 || string.Contains(os.Args[1], "h") {
		fmt.Println(help.Default())
	}

	options := new(model.Options)

	for index, opt := range os.Args {

		switch opt {
		case "version", "--version":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set the file path of application that contain the application version")
				os.Exit(1)
			} else {
				options.Version = os.Args[index+1]
			}
		case "filepath", "-f":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to set the file path of application that contain the application version")
				os.Exit(1)
			} else {
				options.Filepath = os.Args[index+1]
			}
		case "increment", "-i":
			if len(os.Args[index:]) <= 1 {
				fmt.Println("You need to define a increment type - (major | minor | path")
				os.Exit(1)
			} else {
				options.Type = os.Args[index+1]
			}
		}
	}

	for index, opt := range os.Args {

		switch opt {
		case "update":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(help.Update())
				os.Exit(1)
			} else {
				options.Update()
			}
		case "get":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(help.Get())
				os.Exit(1)
			} else {
				options.Get()
			}
		}
	}
}