package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Zakaria")

	if result != "Hello Zakaria" {
		// untuk memunculkan error
		panic("Result is not Hello Zakaria")
	}
}

func TestHelloWorlZakaria(t *testing.T) {
	result := HelloWorld("Zakaria")

	if result != "Hello Error" {
		// untuk memunculkan error
		panic("Result is not Hello Zakaria")
	}
}
