package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Executes commands
func Exec(s string) {
	// Go Logging an error
	// https://www.honeybadger.io/blog/golang-logging/
	// If the file doesn't exist, create it or append to the file
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	cmd := exec.Command("bash", "-c", s)
	_, err := cmd.Output()
	log.SetOutput(file)
	if err != nil {
		fmt.Println("Error occured, please check logs 😭")
		fmt.Fprintln(cmd.Stderr)
		Warn.Printf("Removing %s", ProjectName)
		exec.Command("bash", "-c", fmt.Sprintf("rm -rf %s", ProjectName))
		log.Fatal(err)
	}
}

// Go Formatting string literal: https://stackoverflow.com/questions/17779371/golang-given-a-string-output-an-equivalent-golang-string-literal
// Also: https://groups.google.com/g/golang-nuts/c/ggd3ww3ZKcI
// And: https://www.digitalocean.com/community/tutorials/an-introduction-to-working-with-strings-in-go
