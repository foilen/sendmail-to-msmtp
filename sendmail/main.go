package main

import (
	"bufio"
	"os"
)

func main() {

	//  In /etc/sendmail-to-msmtp.json if present
	ctx := ProcessContext{args: os.Args, consoleReader: bufio.NewReader(os.Stdin)}
	if _, err := os.Stat("/etc/sendmail-to-msmtp.json"); err == nil {
		ctx.configurationPath = "/etc/sendmail-to-msmtp.json"
	}

	process(&ctx)

	// TODO Run the msmtp command (provide the file as input, stream the out/err outputs and return the same status code)

}
