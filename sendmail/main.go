package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {

	//  In /etc/sendmail-to-msmtp.json if present, unless overridden by SENDMAIL_TO_MSMTP_CONFIG_PATH
	ctx := ProcessContext{args: os.Args[1:], consoleReader: bufio.NewReader(os.Stdin)}
	configurationPath := "/etc/sendmail-to-msmtp.json"
	if envConfigurationPath := os.Getenv("SENDMAIL_TO_MSMTP_CONFIG_PATH"); envConfigurationPath != "" {
		configurationPath = envConfigurationPath
	}
	if _, err := os.Stat(configurationPath); err == nil {
		ctx.configurationPath = configurationPath
	}

	// The msmtp binary path can be overridden with SENDMAIL_TO_MSMTP_MSMTP_PATH
	ctx.msmtpPath = os.Getenv("SENDMAIL_TO_MSMTP_MSMTP_PATH")

	// Always show the values that will actually be used
	resolvedMsmtpPath := ctx.msmtpPath
	if resolvedMsmtpPath == "" {
		resolvedMsmtpPath = "/usr/bin/msmtp"
	}
	fmt.Fprintf(os.Stderr, "sendmail-to-msmtp: configuration file: %s (found: %t)\n", configurationPath, ctx.configurationPath != "")
	fmt.Fprintf(os.Stderr, "sendmail-to-msmtp: msmtp binary: %s\n", resolvedMsmtpPath)

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
