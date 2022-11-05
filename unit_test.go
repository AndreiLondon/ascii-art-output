package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	outputFileName := "test00.txt"
	banner := "shadow"
	message := `First\nTest`
	testFileName := "test00Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test01.txt"
	banner = "standard"
	message = `hello`
	testFileName = "test01Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test02.txt"
	banner = "standard"
	message = `123 -> #$%`
	testFileName = "test02Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test03.txt"
	banner = "shadow"
	message = `432 -> #$%&@`
	testFileName = "test03Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test04.txt"
	banner = "shadow"
	message = `There`
	testFileName = "test04Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test05.txt"
	banner = "thinkertoy"
	message = "123 -> \"#$%@"
	testFileName = "test05Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test06.txt"
	banner = "thinkertoy"
	message = "2 you"
	testFileName = "test06Result.txt"
	execute(t, banner, message, outputFileName, testFileName)

	outputFileName = "test07.txt"
	banner = "standard"
	message = "Testing long output!"
	testFileName = "test07Result.txt"
	execute(t, banner, message, outputFileName, testFileName)
}

func execute(t *testing.T, banner string, message string, outputFileName string, testFileName string) {
	convertAndWrite(outputFileName, message, banner)

	data, err := os.ReadFile(testFileName)
	if err != nil {
		t.Fatalf("Failed to read test file")
	}
	str1 := string(data)

	data, err = os.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("Failed to read output file")
	}
	str2 := string(data)

	if str1 != str2 {
		t.Fatalf(`convertText(path) fail.
		Want:
%v 
		Found:
%v`, str1, str2)
	}
}
