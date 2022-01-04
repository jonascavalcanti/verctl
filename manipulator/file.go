package manipulator

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadLinesInFile(fileName string) []string {
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

func ReplaceInFile(filepath, oldVersion, newVersion string) {

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
