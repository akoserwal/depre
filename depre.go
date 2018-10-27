package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Dependencies tree
type Dependencies struct {
	dep string
}

//MavenPackage structure
type MavenPackage struct {
	rootPackage  string
	dependencies []Dependencies
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	checkError(scanner.Err())

	for scanner.Scan() {

		text := scanner.Text()

		parts := strings.SplitAfter(text, "[INFO] +-")

		for i := range parts {
			parts[i] = strings.TrimLeft(parts[i], "[INFO] +-")
			fmt.Println(parts[i])

		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
