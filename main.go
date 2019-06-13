package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {

	defer Recover()

	checked := checkParams(os.Args, 1)

	if !checked {
		panic("Error")
		return
	}
	GetParams := os.Args[1]

	extension := strings.Split(GetParams, ".")

	if len(extension) <= 1 {
		panic("Error File Not Support")
		return
	}

	checkExtension := checkFormat(extension[1])

	if !checkExtension {
		panic("Format Not Support")
		return
	}

	csvFile, err := os.Open(GetParams)

	if err != nil {
		panic("Cannot Open File")
	}

	read := csv.NewReader(csvFile)

	for {
		record, err := read.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		for v := range record {
			runner(record[v])
		}

	}

	// runner()

}

func checkFormat(argv string) bool {
	if argv != "csv" {
		return false
	}

	return true
}

func checkParams(key []string, val int) bool {
	if len(key) > val {
		return true
	}
	return false

}

func runner(imageId string) {
	command := exec.Command("docker", "rmi", "-f", imageId)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Image ID", imageId)
}

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Warning", r)
	}
}
