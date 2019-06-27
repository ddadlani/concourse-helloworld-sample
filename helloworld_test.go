package main_test

import (
	"os"
	"testing"

	sample "github.com/ddadlani/concourse-sample"
)

func TestHandler(t *testing.T) {
	os.Setenv("TARGET", "blah")
	sample.SetTarget("test")
	target := os.Getenv("TARGET")
	if target != "test" {
		t.Error("target was not set correctly")
	}
}
