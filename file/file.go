package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	re "regexp"
	"strconv"
	"strings"
	"time"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

//Read and return the content as string
func GetVersion(filepath string) string {
	count := 0
	rgex := re.MustCompile("=|:")

	var v []string
	var version string

	for _, line := range LinesInFile(filepath) {
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
	} else if typeInc == "rc" {
		version = generateRCVer(oldVersion)
	} else {
		fmt.Println("Type", typeInc, "increment unavailable")
	}

	return version
}

func generateSemVer(oldVersion, typeInc string) string {
	arr := strings.Split(strings.ReplaceAll(oldVersion, "'", ""), ".")

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

func generateRCVer(oldVersion string) string {
	arr := strings.Split(strings.ReplaceAll(oldVersion, "'", ""), ".")

	rcInc, _ := strconv.Atoi(arr[len(arr)-1])
	rcInc++

	var sb strings.Builder

	sb.WriteString("'")
	for i, num := range arr {
		if i < (len(arr) - 1) {
			sb.WriteString(num + ".")
		}
	}
	sb.WriteString(strconv.Itoa(rcInc) + "'")

	version := sb.String()

	return version
}

func WriteVerionOnFile(filepath, oldVersion, newVersion string) {

	file, er := ioutil.ReadFile(filepath)
	if er != nil {
		panic(er)
	}

	output := bytes.Replace(file, []byte(strings.ReplaceAll(oldVersion, "'", "")), []byte(strings.ReplaceAll(newVersion, "'", "")), -1)

	var err error

	if err = ioutil.WriteFile(filepath, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
