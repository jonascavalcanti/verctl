package manipulator

import (
	"fmt"
	"os"
	re "regexp"
	"strconv"
	"strings"
	"time"
)

//Read and return the content as string
func GetVersion(filepath string) string {
	count := 0
	rgex := re.MustCompile("=|:")

	var v []string
	var version string

	for _, line := range ReadLinesInFile(filepath) {
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
		version = generateSemVer(oldVersion, typeInc)
	} else if typeInc == "date" {
		version = generateDateVer(oldVersion)
	} else if typeInc == "rc" || strings.Contains(typeInc, ":") {
		if strings.Contains(typeInc, "major") || strings.Contains(typeInc, "minor") || strings.Contains(typeInc, "patch") {
			semver := strings.Split(typeInc, ":")
			version = generateRCVer(oldVersion, semver[1])
		} else {
			version = generateRCVer(oldVersion, "")
		}

	} else if typeInc == "staging" {
		version = generateStagingVer(oldVersion)
	} else {
		fmt.Println("Type", typeInc, "increment unavailable")
	}

	return version
}

func generateSemVer(oldVersion, typeInc string) string {
	arr := strings.Split(strings.ReplaceAll(oldVersion, "'", ""), ".")

	arr[2] = strings.ReplaceAll(arr[2], "-rc", "")
	arr[2] = strings.ReplaceAll(arr[2], "-staging", "")

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

func generateDateVer(oldVersion string) string {
	arr := strings.Split(strings.ReplaceAll(oldVersion, "'", ""), ".")

	layout := "2006.01.02"
	t := time.Now()

	dayInc, _ := (strconv.Atoi(arr[len(arr)-1]))
	dayInc++

	date := t.Format(layout)

	version := "'" + date + "." + strconv.Itoa(dayInc) + "'"

	return version

}

func generateRCVer(oldVersion, semver string) string {

	var sb strings.Builder
	var rcInc int

	oldVersionRemovedStaging := strings.ReplaceAll(oldVersion, "-staging", "")

	arr := strings.Split(strings.ReplaceAll(oldVersionRemovedStaging, "'", ""), ".")

	if !strings.Contains(oldVersion, "-rc") {
		arr = append(arr, "-rc")
		rcInc = 0
	}

	if semver == "" {
		rcInc, _ = strconv.Atoi(arr[len(arr)-1])
		rcInc++
	} else if semver == "major" {
		major, _ := strconv.Atoi(arr[0])
		major++
		arr[0] = strconv.Itoa(major)
	} else if semver == "minor" {
		minor, _ := strconv.Atoi(arr[1])
		minor++
		arr[1] = strconv.Itoa(minor)
	} else if semver == "patch" {
		var patchNum int
		if strings.Contains(arr[2], "-rc") {
			patchArr := strings.Split(arr[2], "-")
			patchNum, _ = strconv.Atoi(patchArr[0])

		} else {
			patchNum, _ = strconv.Atoi(arr[2])
		}
		patchNum++
		arr[2] = strconv.Itoa(patchNum) + "-rc"
	}

	sb.WriteString("'")

	for i, value := range arr {
		if i < (len(arr) - 1) {
			if arr[i+1] == "-rc" {
				sb.WriteString(value + "-rc.")
			} else {
				sb.WriteString(value + ".")
			}
		}
	}

	sb.WriteString(strconv.Itoa(rcInc) + "'")

	version := sb.String()

	return version
}

func generateStagingVer(oldVersion string) string {
	arr := strings.Split(strings.ReplaceAll(oldVersion, "'", ""), ".")

	v := new(SemVer)

	arr[2] = strings.ReplaceAll(arr[2], "-rc", "")
	arr[2] = strings.ReplaceAll(arr[2], "-staging", "")

	v.Major, _ = strconv.Atoi(arr[0])
	v.Minor, _ = strconv.Atoi(arr[1])
	v.Patch, _ = strconv.Atoi(arr[2])

	version := "'" + strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch) + "-staging" + "'"

	return version
}

func WriteVersionOnFile(filepath, oldVersion, newVersion string) {

	ReplaceInFile(filepath, oldVersion, newVersion)
}
