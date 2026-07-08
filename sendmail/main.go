package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {

	//  In /etc/sendmail-to-msmtp.json if present (or SENDMAIL_TO_MSMTP_CONFIG if set)
	ctx := ProcessContext{args: os.Args[1:], consoleReader: bufio.NewReader(os.Stdin)}
	configurationPath := os.Getenv("SENDMAIL_TO_MSMTP_CONFIG")
	if configurationPath == "" {
		configurationPath = "/etc/sendmail-to-msmtp.json"
	}
	if _, err := os.Stat(configurationPath); err == nil {
		ctx.configurationPath = configurationPath
	}

	// Get the command and its arguments
	sendmailCommandAndArguments := process(&ctx)

	//  Copy the sendmail file in the dump if needed
	if ctx.emailDumpFilePrefix != "" {

		input, err := ioutil.ReadFile(ctx.sendmailFilePath)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(ctx.emailDumpFilePrefix+"-sendmail.eml", input, 0600)
		if err != nil {
			panic(err)
		}

	}

	// Run the msmtp command
	cmd := exec.Command(sendmailCommandAndArguments[0], sendmailCommandAndArguments[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fileIn, err := os.Open(ctx.sendmailFilePath)
	if err != nil {
		panic(err)
	}
	defer fileIn.Close()
	defer os.Remove(ctx.sendmailFilePath)
	cmd.Stdin = fileIn
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

}
