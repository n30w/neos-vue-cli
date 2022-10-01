package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

// Executes commands
// Go Formatting string literal: https://stackoverflow.com/questions/17779371/golang-given-a-string-output-an-equivalent-golang-string-literal
// Also: https://groups.google.com/g/golang-nuts/c/ggd3ww3ZKcI
// And: https://www.digitalocean.com/community/tutorials/an-introduction-to-working-with-strings-in-go

func Exec(s string) {

	c := color.New(color.FgMagenta)

	// Go Logging an error
	// https://www.honeybadger.io/blog/golang-logging/
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	c.Println("Executing " + s)

	cmd := exec.Command("bash", "-c", s)
	// If the file doesn't exist, create it or append to the file
	stdout, err := cmd.Output()
	log.SetOutput(file)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error occured, please check logs 😭")
	}
	fmt.Println(string(stdout))
}

// Builds and retruns a string of commands to execute
// func Multiple(cmd, args string, x func() string, files []string) string { // x for extracted
// 	final := ""
// 	z := x
// 	for _, file := range files {
// 		final += cmd + " " + z
// 	}
// 	return final[0 : len(final)-4]
// }

// func M(cmd string, f func(), files []string) string {
// 	final := ""
// 	for _, file := range files {
// 		final += cmd + " " + f(file)
// 	}
// 	return final
// }
