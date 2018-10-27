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

		//parts := strings.Split(text, "+-")

		parts := strings.SplitAfter(text, "+-")

		//	Loop over the parts from the string.
		for i := range parts {
			fmt.Println(parts[i])
		}

		//fmt.Println(text)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
