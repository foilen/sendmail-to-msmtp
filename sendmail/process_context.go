package main

import "bufio"

// ProcessContext is used to pass the needed information to the Process function.
type ProcessContext struct {
	args          []string
	consoleReader *bufio.Reader

	configurationPath string

	sendmailFilePath string

	emailDumpFilePrefix string
}
