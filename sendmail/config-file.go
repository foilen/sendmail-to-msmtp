package main

import (
	"encoding/json"
	"io/ioutil"
)

// SendmailToMsmtpConfiguration is the json configuration file
type SendmailToMsmtpConfiguration struct {
	DefaultFrom string
}

func getSendmailToMsmtpConfiguration(path string) (*SendmailToMsmtpConfiguration, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var SendmailToMsmtpConfiguration SendmailToMsmtpConfiguration
	err = json.Unmarshal(jsonBytes, &SendmailToMsmtpConfiguration)

	return &SendmailToMsmtpConfiguration, err
}
