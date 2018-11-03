package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	checkError(scanner.Err())
	var depMap = make(map[string]int)

	for scanner.Scan() {

		text := scanner.Text()

		reg, err := regexp.Compile("[^a-zA-Z0-9.:-]")
		checkError(err)

		processedString := reg.ReplaceAllString(text, "")
		deps := strings.Split(processedString, "INFO")

		for i := range deps {
			if strings.HasPrefix(deps[i], "-") == true && strings.HasSuffix(deps[i], "compile") {
				if depMap[deps[i]] == 0 {
					depMap[deps[i]] = 1
				} else {
					depMap[deps[1]] = depMap[deps[1]] + 1
				}
			}
		}
	}

	for k, v := range depMap {
		fmt.Printf("dependency:[%s] occurance:[%d]\n", k, v)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
