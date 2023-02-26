package main

import (
	"bufio"
	"os"
)

func main() {
	// create reader to read from STDIN
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()

	// create writer to write on STDOUT
	writer := bufio.NewWriter(os.Stdout)

	// convert string to byte arry
	writer.Write([]byte(text))

	// finally flush buffer
	writer.Flush()

}
