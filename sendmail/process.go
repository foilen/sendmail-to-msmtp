package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func process(args []string, consoleReader *bufio.Reader) []string {

	sendmailArguments := []string{"/usr/bin/msmtp"}

	// Log arguments
	fArg, err := os.OpenFile("/tmp/sendmail-arguments.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer fArg.Close()

	if _, err = fArg.WriteString(strings.Join(args, " ") + "\n"); err != nil {
		panic(err)
	}

	// TODO Get the default from the config file
	var sender = "noname@nohost"
	var readMessageForRecipients = false

	// Process the arguments
	for i := 0; i < len(args); i++ {

		// Find the "-r" or "-f"
		if args[i] == "-r" || args[i] == "-f" {
			sender = args[i+1]
			i++
		}

		// Find the "-t"
		if args[i] == "-t" {
			readMessageForRecipients = true
		}

	}

	// Log email
	t := time.Now()
	tStr := strconv.FormatInt(t.Unix(), 10)
	rStr := strconv.FormatInt(rand.Int63(), 10)
	contentFileName := "/tmp/sendmail-content-" + tStr + rStr + ".txt"
	fContent, err := os.OpenFile(contentFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer fContent.Close()

	for {
		line, err := consoleReader.ReadString('\n')

		// EOF
		if err != nil {
			break
		}

		// Sanitize
		line = strings.TrimSpace(line)

		// Find the "From: "
		if strings.HasPrefix(line, "From: ") {
			sender = line[6:]
		}

		// Write to file
		if _, err = fContent.WriteString(line); err != nil {
			panic(err)
		}
	}

	// Set the sender
	if readMessageForRecipients {
		sendmailArguments = append(sendmailArguments, "-t")
	}
	sendmailArguments = append(sendmailArguments, "-f", sender)

	// TODO Get the addresses at the end

	return sendmailArguments

}
