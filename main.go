package main

import (
	"bufio"
	"os"
)

func main() {

	process(os.Args[1:], bufio.NewReader(os.Stdin))

}
