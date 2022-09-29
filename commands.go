package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Commands struct {
	git   []string
	mkdir []string
	// touch []string
	cd []string
	// mv    []string
	yarn []string
}

// exec command wrapper
// Go Formatting string literal: https://stackoverflow.com/questions/17779371/golang-given-a-string-output-an-equivalent-golang-string-literal
// Also: https://groups.google.com/g/golang-nuts/c/ggd3ww3ZKcI
// And: https://www.digitalocean.com/community/tutorials/an-introduction-to-working-with-strings-in-go
func Exec(s string) {
	// Go Logging an error
	// https://www.honeybadger.io/blog/golang-logging/
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fmt.Println("entered execution...")
	cmd := exec.Command("bash", "-c", s)
	// If the file doesn't exist, create it or append to the file
	stdout, err := cmd.Output()
	log.SetOutput(file)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error occured, please check logs 😭😭😭")
	}
	fmt.Println(string(stdout))
}
