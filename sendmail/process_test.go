package main

import (
	"bufio"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestSenderInF(t *testing.T) {

	args := []string{"-f", "sender-arg@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "sender-arg@foilen-lab.com"}

	actual := process(args, reader)

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

	actual := process(args, reader)

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

	actual := process(args, reader)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}

func TestSenderInR(t *testing.T) {

	args := []string{"-r", "sender-arg@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "sender-arg@foilen-lab.com"}

	actual := process(args, reader)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}
