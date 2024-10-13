package main

import (
	"log"
	"os"
	"testing"
)

func TestByteCount(t *testing.T) {

	expected_byte := 342190
	expected_word := 58164
	expected_line := 7145

	data, err := os.Stat("./test.txt")
	if err != nil {
		log.Println("err getting stat of test file")
		t.Fail()
	}
	file_content, err := os.ReadFile("./test.txt")

	if err != nil {
		log.Println("err getting content test file")
		t.Fail()
	}

	if expected_byte != byteCount(data) {
		log.Println("byte count off")
		t.Fail()
	}
	if expected_word != wordCount(file_content) {
		log.Println("byte count off")
		t.Fail()
	}

	if expected_line != lineCount(file_content) {
		log.Println("byte count off")
		t.Fail()
	}
}
