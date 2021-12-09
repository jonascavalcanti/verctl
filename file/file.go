package file

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Read and return the content as string
func GetVersion(filepath string) (string, []byte) {
	file, error := ioutil.ReadFile(filepath)
	if error != nil {
		panic(error)
	}
	strVersion := string(file)
	version := strings.Split(strVersion, "=")

	return string(version[1]), file
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func IncrementVersion(oldVersion, separator, typeInc string) string {
	arr := strings.Split(oldVersion, separator)

	v := new(Version)

	v.Major, _ = strconv.Atoi(arr[0])
	v.Minor, _ = strconv.Atoi(arr[1])
	v.Patch, _ = strconv.Atoi(arr[2])

	if typeInc == "major" {
		v.Major++
	} else if typeInc == "minor" {
		v.Minor++
	} else {
		v.Patch++
	}
	newVersion := strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
	return newVersion
}

func WriteVerionOnFile(oldVersion, newVersion string, file []byte) {
	output := bytes.Replace(file, []byte(oldVersion), []byte(newVersion), -1)

	var err error

	if err = ioutil.WriteFile("./app.properties", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
