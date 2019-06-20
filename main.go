package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(os.Args)

	f, err := os.OpenFile("/tmp/sendmail-arguments.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(strings.Join(os.Args[1:], " ") + "\n"); err != nil {
		panic(err)
	}

}
