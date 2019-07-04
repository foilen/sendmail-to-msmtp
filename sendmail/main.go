package main

import (
	"bufio"
	"os"
)

func main() {

	ctx := ProcessContext{args: os.Args, consoleReader: bufio.NewReader(os.Stdin)}
	process(&ctx)

	// TODO Run the msmtp command (provide the file as input, stream the out/err outputs and return the same status code)

}
