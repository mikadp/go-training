//https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

func getFileData() (inputFile, error) {
	//validet that we are getting the correct number of arguments
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A filepath argument is required")

	}

	// Defining option flags
	//We need to define three arguments: the flags name, the default value and short descroption
	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse() //this will parse all the arguments from the terminal

	fileLocation := flag.Arg(0) // The only argument

	//Validating whether or not we recived "comma" or "semicolon" from the parsed arguments
	// if didin't receive any of those. return error.

	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("Only comma or semicolon separators are allowed")
	}

	//Return correspongin struct instance with all the required data
	return inputFile{fileLocation, *separator, *pretty}, nil
}

func checkIfValidFile(filename string) (bool, error) {
	//Checking if entered file is CSV by using the filepath package
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return false, fmt.Errorf("File %s is not CSV", filename)
	}

	// Checking if filepath entered belongs to an existing file.
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %s does not exist", filename)
	}

	return true, nil
}

func main() {
}
