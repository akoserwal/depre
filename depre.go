package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//COMPILE scope
const COMPILE = "compile"

//TEST scope
const TEST = "test"

//INFO build type
const INFO = "INFO"

func main() {

	const regexPattern = "[^a-zA-Z0-9.:-]"

	scanner := bufio.NewScanner(os.Stdin)
	checkError(scanner.Err())
	var depMap = make(map[string]int)

	for scanner.Scan() {

		text := scanner.Text()

		reg, err := regexp.Compile(regexPattern)
		checkError(err)

		processedString := reg.ReplaceAllString(text, "")
		deps := strings.Split(processedString, INFO)

		countDependencyOccurance(deps, depMap)
		checkError(err)

	}

	iterateMapAndPrint(depMap)

}

func iterateMapAndPrint(m map[string]int) {
	for k, v := range m {
		fmt.Printf("[%s] occurance:[%d]\n", k, v)

	}
}

func countDependencyOccurance(deps []string, depMap map[string]int) map[string]int {

	for i := range deps {
		if strings.Contains(deps[i], COMPILE) || strings.Contains(deps[i], TEST) {
			value, ok := depMap[deps[i]]
			if ok == true {
				depMap[deps[i]] = value + 1
			} else {
				depMap[deps[i]] = 1

			}
		}
	}
	return depMap
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
