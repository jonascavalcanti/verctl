package model

import (
	"os"
	"fmt"
	"strings"
	"github.com/jonascavalcantineto/versctl/help"
	controller "github.com/jonascavalcantineto/versctl/controller"
)

type Options struct {
	Type     string
	Filepath string
	Version  string
}

func (opts Options) Update() {

	oldVersionTmp := ""
	if opts.Filepath != "" {
		oldVersionTmp = controller.GetVersion(opts.Filepath)
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
	newVersion := controller.IncrementVersion(oldVersion, opts.Type)

	if opts.Filepath != "" {
		controller.WriteVersionOnFile(opts.Filepath, oldVersion, newVersion)
		fmt.Println("Application Version:", oldVersion)
		fmt.Println("New Version:", newVersion)
	} else if opts.Version != "" {
		fmt.Println(strings.ReplaceAll(newVersion, "'", ""))
	}

}

func (opts Options) Get() {
	fmt.Println(controller.GetVersion(opts.Filepath))
}
