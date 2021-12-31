package semver

import (
	"fmt"
	"os"
	re "regexp"
	"strconv"
	"strings"
	"time"
	"xversioner/file"
)

//Read and return the content as string
func GetVersion(filepath string) string {
	count := 0
	rgex := re.MustCompile("=|:")

	var v []string
	var version string

	for _, line := range file.LinesInFile(filepath) {
		if strings.Contains(line, "version") || strings.Contains(line, "VERSION") {
			//fmt.Printf("Line Number = %v, line = %v\n", index, line)
			if count < 1 {
				v = rgex.Split(line, -1)
				//fmt.Println(version)
				version = v[1]
			}
			count++
		}

	}

	if version == "" {
		fmt.Println("File have no version tag or file does not exist on", filepath)
		os.Exit(1)
	}

	return strings.ReplaceAll(version, "'", "")
}

type SemVer struct {
	Major int
	Minor int
	Patch int
}

func IncrementVersion(oldVersion, typeInc string) string {

	version := ""

	if typeInc == "major" || typeInc == "minor" || typeInc == "patch" {
		version = GenerateSemVer(oldVersion, typeInc)
	} else if typeInc == "date" {
		version = GenerateDateVer(oldVersion)
	} else {
		fmt.Println("Type", typeInc, "increment unavailable")
	}

	return version
}

func GenerateSemVer(oldVersion, typeInc string) string {
	arr := strings.Split(oldVersion, ".")

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

	version := "'" + strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch) + "'"

	return version
}

func GenerateDateVer(oldVersion string) string {
	arr := strings.Split(strings.ReplaceAll(oldVersion, "'", ""), ".")

	layout := "2006.01.02"
	t := time.Now()

	dayInc, _ := (strconv.Atoi(arr[len(arr)-1]))
	dayInc++

	date := t.Format(layout)

	version := "'" + date + "." + strconv.Itoa(dayInc) + "'"

	return version

}

func WriteVersionOnFile(filepath, oldVersion, newVersion string) {

	file.ReplaceOnFile(filepath, oldVersion, newVersion)
}
