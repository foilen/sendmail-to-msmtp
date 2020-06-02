package main

import (
	"bufio"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func AssertFilesContent(t *testing.T, expectedFile string, actualFile string) {

	email, err := ioutil.ReadFile(expectedFile)
	if err != nil {
		panic(err)
	}
	expectedEmail := string(email)
	actual, err := ioutil.ReadFile(actualFile)
	if err != nil {
		panic(err)
	}
	actualEmail := string(actual)
	if actualEmail != expectedEmail {
		t.Errorf("Expected file content: \n//////////\n%s\n//////////\n ; Got file content: \n//////////\n%s\n//////////\n", expectedEmail, actualEmail)
	}

}

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

	AssertFilesContent(t, "testdata/process_test_no_from.txt", ctx.sendmailFilePath)

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

	AssertFilesContent(t, "testdata/process_test_with_from.txt", ctx.sendmailFilePath)

}

func TestSenderInHeaderWithName(t *testing.T) {

	args := []string{"-t"}
	email, err := ioutil.ReadFile("testdata/process_test_with_from_with_name.txt")
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

	AssertFilesContent(t, "testdata/process_test_with_from_with_name.txt", ctx.sendmailFilePath)

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

func TestMultipart(t *testing.T) {

	args := []string{"-t", "-i", "-f", "sender@foilen-lab.com"}
	email, err := ioutil.ReadFile("testdata/process_test_multipart.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(strings.NewReader(string(email)))
	expected := []string{"/usr/bin/msmtp", "-t", "-f", "sender@foilen-lab.com"}

	ctx := ProcessContext{args: args, consoleReader: reader}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

	AssertFilesContent(t, "testdata/process_test_multipart.txt", ctx.sendmailFilePath)

}

func TestMultipartTab(t *testing.T) {

	args := []string{"-t", "-i", "-f", "sender@foilen-lab.com"}
	email, err := ioutil.ReadFile("testdata/process_test_multipart-tab.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(strings.NewReader(string(email)))
	expected := []string{"/usr/bin/msmtp", "-t", "-f", "sender@foilen-lab.com"}

	ctx := ProcessContext{args: args, consoleReader: reader}
	actual := process(&ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

	AssertFilesContent(t, "testdata/process_test_multipart-tab.txt", ctx.sendmailFilePath)

}
