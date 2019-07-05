package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func process(ctx *ProcessContext) []string {

	sendmailArguments := []string{"/usr/bin/msmtp"}

	// Log arguments
	fArg, err := os.OpenFile("/tmp/sendmail-arguments.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer fArg.Close()

	if _, err = fArg.WriteString(strings.Join(ctx.args, " ") + "\n"); err != nil {
		panic(err)
	}

	var sender = ""
	var readMessageForRecipients = false
	recipients := []string{}

	// Get the default sender from the config file
	if ctx.configurationPath != "" {
		configuration, err := getSendmailToMsmtpConfiguration(ctx.configurationPath)
		if err != nil {
			panic(err)
		}
		sender = configuration.DefaultFrom
	}

	// Process the arguments
	var endOfOptions = false
	for i := 0; i < len(ctx.args); i++ {

		// Check if no more an option
		if ctx.args[i][0] != '-' {
			endOfOptions = true
		}

		// Find the recipients
		if endOfOptions {
			recipients = append(recipients, ctx.args[i])
		}

		// Find the "-r" or "-f"
		if ctx.args[i] == "-r" || ctx.args[i] == "-f" {
			sender = ctx.args[i+1]
			i++
		}

		// Find the "-t"
		if ctx.args[i] == "-t" {
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
		line, err := ctx.consoleReader.ReadString('\n')

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

	// Set the sendmail arguments
	if readMessageForRecipients {
		sendmailArguments = append(sendmailArguments, "-t")
	}
	if sender != "" {
		sendmailArguments = append(sendmailArguments, "-f", sender)
	}

	if len(recipients) > 0 {
		sendmailArguments = append(sendmailArguments, "--")
		sendmailArguments = append(sendmailArguments, recipients...)
	}

	return sendmailArguments

}
