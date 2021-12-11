package main

import (
	"fmt"
	"os"
	"xversioner/file"
	"xversioner/help"
)

type Options struct {
	Type     string
	Filepath string
}

func main() {

	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help" {
		fmt.Println(help.Default())
	}

	options := new(Options)

	for index, opt := range os.Args {

		switch opt {
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
				options.update()
			}
		case "get":
			if len(os.Args[index:]) <= 1 || os.Args[index+1] == "--help" || os.Args[index+1] == "help" {
				fmt.Println(help.Get())
				os.Exit(1)
			} else {
				options.get()
			}
		}
	}

}

func (opts Options) update() {

	oldVersion, fileByteArray := file.GetVersion(opts.Filepath)
	fmt.Println("Application Version:", oldVersion)

	newVersion := file.IncrementVersion(oldVersion, ".", opts.Type)
	fmt.Println("New Version:", newVersion)

	file.WriteVerionOnFile(oldVersion, newVersion, fileByteArray)

}

func (opts Options) get() {
	version, _ := file.GetVersion(opts.Filepath)
	fmt.Println(version)
}
