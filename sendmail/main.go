package main

import (
	"bufio"
	"os"
	"os/exec"
)

func main() {

	//  In /etc/sendmail-to-msmtp.json if present
	ctx := ProcessContext{args: os.Args[1:], consoleReader: bufio.NewReader(os.Stdin)}
	if _, err := os.Stat("/etc/sendmail-to-msmtp.json"); err == nil {
		ctx.configurationPath = "/etc/sendmail-to-msmtp.json"
	}

	// Get the command and its arguments
	sendmailCommandAndArguments := process(&ctx)

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
