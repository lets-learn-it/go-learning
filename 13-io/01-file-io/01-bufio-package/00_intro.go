package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./capitals.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// default behaviour is to get whole line
	// We want each word
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}
