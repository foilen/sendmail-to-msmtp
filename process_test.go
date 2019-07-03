package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestSenderInF(t *testing.T) {

	args := []string{"-f", "sender@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "sender@foilen-lab.com"}

	actual := process(args, reader)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}
func TestSenderInR(t *testing.T) {

	args := []string{"-r", "sender@foilen-lab.com"}
	reader := bufio.NewReader(strings.NewReader(""))
	expected := []string{"/usr/bin/msmtp", "-f", "sender@foilen-lab.com"}

	actual := process(args, reader)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s ; Got: %s", strings.Join(expected, " "), strings.Join(actual, " "))
	}

}
