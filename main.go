package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println(os.Args)

	// Log arguments
	fArg, err := os.OpenFile("/tmp/sendmail-arguments.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer fArg.Close()

	if _, err = fArg.WriteString(strings.Join(os.Args[1:], " ") + "\n"); err != nil {
		panic(err)
	}

	// Log email
	t := time.Now()
	tStr := strconv.FormatInt(t.Unix(), 10)
	contentFileName := "/tmp/sendmail-content-" + tStr + ".txt"
	fmt.Println(contentFileName)
	fContent, err := os.OpenFile(contentFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer fContent.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')

		// EOF
		if err != nil {
			break
		}

		// Write to file
		if _, err = fContent.WriteString(line); err != nil {
			panic(err)
		}
	}

}
