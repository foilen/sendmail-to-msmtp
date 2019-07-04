package main

import (
	"bufio"
	"os"
)

func main() {

	process(os.Args[1:], bufio.NewReader(os.Stdin))

	// TODO Run the msmtp command (provide the file as input, stream the out/err outputs and return the same status code)

}
