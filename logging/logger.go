package main

import (
	"log"
	"os"
	"path"
)

var (
	Warn   *log.Logger
	Error  *log.Logger
	Notice *log.Logger
)

var pwd, err = os.Getwd()

func main() {
	warnFile, err := os.OpenFile(path.Join(pwd, "warnings.log"), os.O_RDWR|os.O_APPEND, 0660)
	defer warnFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	errorFile, err := os.OpenFile(path.Join(pwd, "errors.log"), os.O_RDWR|os.O_APPEND, 0660)
	defer errorFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	Warn = log.New(warnFile, "WARNING: ", log.LstdFlags)

	Warn.Println("Messages written to a file called 'warnings.log' are likely to be ignored")

	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime)
	Error.SetOutput(errorFile)
	Error.Println("Error messages, on the other hand, tend to catch attention!")
}
