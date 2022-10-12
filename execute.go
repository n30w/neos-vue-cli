package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Command wrapper with error handling
func Exec(s string) {
	// Go Logging an error: https://www.honeybadger.io/blog/golang-logging/
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// The good stuff
	cmd := exec.Command("bash", "-c", s)
	_, err := cmd.Output()
	log.SetOutput(file)
	if err != nil {
		Warn.Println("Error occured, please check logs 😭")
		Warn.Fprintln(cmd.Stderr)
		Warn.Printf("Removing %s", ProjectName)
		os.RemoveAll(ProjectName)
		log.Fatal(err)
	}
}

// Inserts a file with a specified line. Useful for CSS.
// Function based on https://zetcode.com/golang/writefile/
func Insert(text string) {
	f, err := os.Create(fmt.Sprintf("%s/src/index.scss", ProjectName))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(fmt.Sprintf("%s\n", text))

	if err2 != nil {
		log.Fatal(err2)
	}
}

// Go Formatting string literal: https://stackoverflow.com/questions/17779371/golang-given-a-string-output-an-equivalent-golang-string-literal
// Also: https://groups.google.com/g/golang-nuts/c/ggd3ww3ZKcI
// And: https://www.digitalocean.com/community/tutorials/an-introduction-to-working-with-strings-in-go
