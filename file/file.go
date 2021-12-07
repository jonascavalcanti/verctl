package file

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Read and return the content as string
func GetVersion(name string) []string {
	file, error := ioutil.ReadFile(name)
	if error != nil {
		panic(error)
	}
	strVersion := string(file)
	version := strings.Split(strVersion, "=")

	return version
}

type ver struct {
	major   string `example:"major"`
	minor 	string `example:"minor"`
	patch   string `example:"patch"`
}

func IncrementVersion(ver string, separator string) string {
	versionSplited := strings.Split(ver, separator)

	for 
	fmt.Println(versionSplited)

	return ""
}
