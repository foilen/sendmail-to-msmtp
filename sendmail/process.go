package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func process(ctx *ProcessContext) []string {

	sendmailArguments := []string{"/usr/bin/msmtp"}

	// Initial values
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
		if configuration.EmailDumpDirectory != "" {
			ctx.emailDumpFilePrefix = configuration.EmailDumpDirectory + "/" + strconv.Itoa(int(time.Now().UnixNano()))
		}
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

	// Copy email in temporary file
	if ctx.sendmailFilePath == "" {
		tmpFile, err := ioutil.TempFile("/tmp", "")
		if err != nil {
			panic(err)
		}
		ctx.sendmailFilePath = tmpFile.Name()
	}
	fContent, err := os.OpenFile(ctx.sendmailFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer fContent.Close()

	// Open raw file if requested
	var fRawDump *os.File
	if ctx.emailDumpFilePrefix != "" {
		fRawDump, err = os.OpenFile(ctx.emailDumpFilePrefix+"-raw.eml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer fRawDump.Close()
	}

	// Read the headers
	var headerLines []string
	for {
		line, err := ctx.consoleReader.ReadString('\n')

		// EOF
		if err != nil {
			break
		}

		// Put in raw dump file if needed
		if fRawDump != nil {
			if _, err = fRawDump.WriteString(line); err != nil {
				panic(err)
			}
		}

		// Sanitize
		line = strings.Trim(line, "\n\r")

		// Add to list
		if len(line) != 0 {
			headerLines = append(headerLines, line)
		}

		// Write to file
		line = line + "\n"
		if _, err = fContent.WriteString(line); err != nil {
			panic(err)
		}

		// Stop if no more the headers
		if len(line) == 1 {
			break
		}
	}

	// Find the "From: "
	for i := 0; i < len(headerLines); i++ {

		var line = headerLines[i]
		if strings.HasPrefix(line, "From: ") {

			// Append all next lines that are starting with a space
			for j := i + 1; j < len(headerLines); j++ {
				var nextLine = headerLines[j]
				if strings.HasPrefix(nextLine, " ") || strings.HasPrefix(nextLine, "\t") {
					i++
					line += nextLine
				} else {
					break
				}
			}

			// If in the form "From: The Sender <sender-header@foilen-lab.com>"
			lBracket := strings.Index(line, "<")
			rBracket := strings.Index(line, ">")
			if lBracket > 0 && rBracket > lBracket {
				sender = line[lBracket+1 : rBracket]
			} else {
				// Else in the form "From: sender-header@foilen-lab.com"
				sender = line[6:]
			}
		}
	}

	// Process the body
	for {
		line, err := ctx.consoleReader.ReadString('\n')

		// EOF
		if err != nil {
			break
		}

		// Put in raw dump file if needed
		if fRawDump != nil {
			if _, err = fRawDump.WriteString(line); err != nil {
				panic(err)
			}
		}

		// Sanitize
		line = strings.Trim(line, "\n\r")

		// Write to file
		line = line + "\n"
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
