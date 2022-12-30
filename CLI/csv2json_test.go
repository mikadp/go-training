package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_getFileData(t *testing.T) {
	//Defining test slice. Each unit test should have the following properties:
	tests := []struct {
		name    string    //The name of the test
		want    inputFile //InputFile instance that function returns
		wantErr bool      // error
		osArgs  []string  // The commang arguments used for this test
	}{
		//Declare each unit test input and output data
		{"Default parameters", inputFile{"test.csv", "comma", false}, false, []string{"cmd", "test.csv"}},
		{"No parameters", inputFile{}, true, []string{"cmd"}},
		{"semicolon enabled", inputFile{"test.csv", "semicolon", false}, false, []string{"cmd", "--separator=semicolon", "test.csv"}},
		{"Pretty enabled", inputFile{"test.csv", "comma", true}, false, []string{"cmd", "--pretty", "test.csv"}},
		{"Pretty and semicolon enabled", inputFile{"test.csv", "semicolon", true}, false, []string{"cmd", "--pretty", "--separator=semicolon", "test.csv"}},
		{"Separator not identified", inputFile{}, true, []string{"cmd", "--separator=pipe", "test.csv"}},
	}
	// iterating over the previous test slice
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Saving the original os.Args reference
			actualOsArgs := os.Args
			//defer function will run after the test is done
			defer func() {
				os.Args = actualOsArgs                                           //restoring the original os.Args reference
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reseting the Flag command line. So that we can parse flags again
			}()

			os.Args = tt.osArgs             // setting the specific command args for this test
			got, err := getFileData()       // running the function we want to test
			if (err != nil) != tt.wantErr { // Asserting whether or not we get the correct error value
				t.Errorf("getFileData() error = %v, wantErr %v", err, tt.wantErr)
				return

			}
			if !reflect.DeepEqual(got, tt.want) { // Asserting whether or not we get the corret wanted value
				t.Errorf("getFileData() = %v, want %v", got, tt.want)
			}

		})
	}
}
