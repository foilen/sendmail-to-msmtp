package main

import (
	"bufio"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestSenderInConfig(t *testing.T) {

	args := []string{"d1@foilen-lab.com", "d2@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "default@foilen-lab.com", "--", "d1@foilen-lab.com", "d2@foilen-lab.com"}
	configurationPath := "testdata/process_test_in_config.json"

	ctx := ProcessContext{args: args, consoleReader: reader, configurationPath: configurationPath}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}

func TestSenderInF(t *testing.T) {

	args := []string{"-f", "sender-arg@foilen-lab.com", "d1@foilen-lab.com", "d2@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "sender-arg@foilen-lab.com", "--", "d1@foilen-lab.com", "d2@foilen-lab.com"}
	configurationPath := "testdata/process_test_in_config.json"

	ctx := ProcessContext{args: args, consoleReader: reader, configurationPath: configurationPath}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}

func TestSenderInF2(t *testing.T) {

	args := []string{"-f", "sender-arg@foilen-lab.com"}
	email, err := ioutil.ReadFile("testdata/process_test_no_from.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(strings.NewReader(string(email)))
	expected := []string{"/usr/bin/msmtp", "-f", "sender-arg@foilen-lab.com"}

	ctx := ProcessContext{args: args, consoleReader: reader}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}

func TestSenderInHeader(t *testing.T) {

	args := []string{"-t"}
	email, err := ioutil.ReadFile("testdata/process_test_with_from.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(strings.NewReader(string(email)))
	expected := []string{"/usr/bin/msmtp", "-t", "-f", "sender-header@foilen-lab.com"}

	ctx := ProcessContext{args: args, consoleReader: reader}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}

func TestSenderInR(t *testing.T) {

	args := []string{"-r", "sender-arg@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "sender-arg@foilen-lab.com"}

	ctx := ProcessContext{args: args, consoleReader: reader}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}
