package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Exec is a command wrapper with error handling
func Exec(s *string) {
	// Go Logging an error: https://www.honeybadger.io/blog/golang-logging/
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// The good stuff
	cmd := exec.Command("bash", "-c", *s)
	o, err := cmd.Output()
	log.SetOutput(file)
	if err != nil {
		Warn.Println("Error occurred, please check logs 😭")
		Warn.Fprintln(cmd.Stderr)
		Warn.Printf("Removing %s\n", ProjectName)
		os.RemoveAll(ProjectName)
		log.Fatal(o, err)
	}
}

// Insert inserts a file with a specified line. Useful for CSS.
// Function based on https://zetcode.com/golang/writefile/
func Insert(text *string) {
	f, err := os.Create(fmt.Sprintf("%s/src/index.scss", ProjectName))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", *text)); err != nil {
		log.Fatal(err)
	}

}

// Go Formatting string literal: https://stackoverflow.com/questions/17779371/golang-given-a-string-output-an-equivalent-golang-string-literal
// Also: https://groups.google.com/g/golang-nuts/c/ggd3ww3ZKcI
// And: https://www.digitalocean.com/community/tutorials/an-introduction-to-working-with-strings-in-go
