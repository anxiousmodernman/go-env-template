package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestRun(t *testing.T) {

	os.Setenv("ABCDEF", "caterpillar")

	tmp, err := ioutil.TempFile("testfixtures", "temp")
	if err != nil {
		t.Errorf("error created temp file: %v", err)
	}
	tmp.Close()
	defer os.Remove(tmp.Name())

	err = Run("testfixtures/template.properties", tmp.Name())
	if err != nil {
		t.Errorf("error calling Run: %v", err)
	}

	// Re-open temp file
	f, err := os.Open(tmp.Name())
	if err != nil {
		t.Errorf("error re-opening temp file: %v", err)
	}

	ef, err := os.Open("testfixtures/expected.properties")
	if err != nil {
		t.Errorf("error opening expected file: %v", err)
	}

	var expected []byte
	var found []byte

	found, err = ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("error reading file we templated: %v", err)
	}

	expected, err = ioutil.ReadAll(ef)
	if err != nil {
		t.Errorf("error reading expected file: %v", err)
	}

	if bytes.Compare(found, expected) != 0 {
		t.Errorf("files were not the same")
	}

}
