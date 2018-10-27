package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	checkError(scanner.Err())

	for scanner.Scan() {

		text := scanner.Text()

		parts := strings.SplitAfter(text, "+-")

		for i := range parts {
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
