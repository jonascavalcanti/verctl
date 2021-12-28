package file

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
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

type SemVer struct {
	Major int
	Minor int
	Patch int
}

func IncrementVersion(oldVersion, separator, typeInc string) string {

	version := ""

	if typeInc == "major" || typeInc == "minor" || typeInc == "patch" {
		version = generateSemVer(oldVersion, separator, typeInc)
	} else if typeInc == "date" {
		version = generateDateVer(oldVersion, separator)
	} else {
		fmt.Println("Type", typeInc, "increment unavailable")
	}

	return version
}

func generateSemVer(oldVersion, separator, typeInc string) string {
	arr := strings.Split(oldVersion, separator)

	v := new(SemVer)

	v.Major, _ = strconv.Atoi(arr[0])
	v.Minor, _ = strconv.Atoi(arr[1])
	v.Patch, _ = strconv.Atoi(arr[2])

	if typeInc == "major" {
		v.Major++
	} else if typeInc == "minor" {
		v.Minor++
	} else if typeInc == "patch" {
		v.Patch++
	} else {
		fmt.Println("Increment Type", typeInc, "unavailable ")
	}

	version := strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)

	return version
}

/**/
func generateDateVer(oldVersion, separator string) string {
	arr := strings.Split(oldVersion, separator)

	fmt.Println(oldVersion)

	layout := "2006.01.02"

	t := time.Now()

	dayInc, _ := (strconv.Atoi(arr[len(arr)-1]))

	dayInc++

	date := t.Format(layout)

	version := date + "." + strconv.Itoa(dayInc)

	return version

}

/**/
func WriteVerionOnFile(filepath, oldVersion, newVersion string, file []byte) {
	output := bytes.Replace(file, []byte(oldVersion), []byte(newVersion), -1)

	var err error

	if err = ioutil.WriteFile(filepath, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
