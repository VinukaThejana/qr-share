package utils

import (
	"fmt"
	"testing"
)

func TestGetDirAndFile(t *testing.T) {
	// contains a table of directories or files with different relative or
	// absolute paths, the value boolean represents wether the key is a
	// directory or a file if the key is a directory then the value will be
	// true else the value will be false
	table := map[string]bool{
		"/Users/vinukakodituwakku":                          true,
		"/Users/vinukakodituwakku/":                         true,
		"/Users/vinukakodituwakku/Downloads/testing/go.mod": false,
		"~/Downloads/":                                      true,
		"~/Downloads":                                       true,
		"~/Downloads/testing/go.mod":                        false,
		"../cmd/":                                           true,
		"../cmd/cmd.go":                                     false,
		"../cmd":                                            true,
		"./text.go":                                         false,
		"./../cmd/cmd.go":                                   false,
		"./../cmd":                                          true,
		"./../cmd/":                                         true,
		"./../../qr":                                        true,
	}

	for filepath, isDir := range table {
		dir, file, err := GetDirAndFile(filepath)
		if err != nil {
			t.Fatal(err)
		}

		if isDir {
			if file == "" {
				fmt.Print(" ✅ ")
			} else {
				fmt.Print(" ❌ ")
				t.Fail()
			}
		} else {
			if file == "" {
				fmt.Print(" ❌ ")
				t.Fail()
			} else {
				fmt.Print(" ✅ ")
			}
		}

		fmt.Printf("\tfilepath : %v\n", filepath)
		fmt.Printf("\tdir      : %v\n", dir)
		fmt.Printf("\tfile     : %v\n", file)
		fmt.Println("")
	}
}
