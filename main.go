package main

import (
	"fmt"
	"os"
	"strings"
	"xversioner/help"
	"xversioner/manipulator"
)

type Options struct {
	Type     string
	Filepath string
	Version  string
}

func main() {

	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help" {
		fmt.Println(help.Default())
	}

	options := new(Options)

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

	oldVersionTmp := ""
	if opts.Filepath != "" {
		oldVersionTmp = manipulator.GetVersion(opts.Filepath)
	} else if opts.Version != "" {
		oldVersionTmp = opts.Version
		if oldVersionTmp != "" {
			if strings.Contains(oldVersionTmp, "v") {
				oldVersionTmp = strings.ReplaceAll(oldVersionTmp, "v", "")
			} else if strings.Contains(oldVersionTmp, "V") {
				oldVersionTmp = strings.ReplaceAll(oldVersionTmp, "V", "")
			}
		}
	} else {
		fmt.Println(help.Default())
		os.Exit(1)
	}

	oldVersion := "'" + oldVersionTmp + "'"
	newVersion := manipulator.IncrementVersion(oldVersion, opts.Type)

	if opts.Filepath != "" {
		manipulator.WriteVersionOnFile(opts.Filepath, oldVersion, newVersion)
		fmt.Println("Application Version:", oldVersion)
		fmt.Println("New Version:", newVersion)
	} else if opts.Version != "" {
		//os.Setenv("newVersionApp", newVersion)
		//fmt.Println("newVersionApp:", os.Getenv("newVersionApp"))
		fmt.Println(newVersion)
	}

}

func (opts Options) get() {
	fmt.Println(manipulator.GetVersion(opts.Filepath))
}
