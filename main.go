package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// TODO: add dependencies like fatih/color and yacspin

const MAXARGS = 1

var (
	css         CSS
	projectName string

	gists = Gists{
		packageJSON:   "https://gist.github.com/b6e6e41894e0d6b3ef7aba33214415ce.git",
		indexHTML:     "https://gist.github.com/b9f38f17a0b2cf3f28d2715011e03fb1.git",
		indexTS:       "https://gist.github.com/3bb7a5789c91bc0229dcbfe209f0fc67.git",
		templateVUE:   "https://gist.github.com/94f18653a1e0468de83faa163d7cdbcf.git",
		postcssrc:     "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		popperJS:      "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		bootstrapSCSS: "https://gist.github.com/b29599aa95343ad7ff3a704c0e9b2d81.git",
		tailwindSCSS:  "https://gist.github.com/35a637a7e185333b08c730b7d64189d3.git",
		tailwindconf:  "https://gist.github.com/ed36d206090bd1faeea8d0c1921e19fc.git",
	}

	coreDependencies = [4]string{
		"parcel",
		"vue",
		"vue-router@4",
		"rimraf",
	}

	tailwind = Tailwind{
		postcssrc:    gists.postcssrc,
		configJS:     "",
		indexSCSS:    gists.tailwindSCSS,
		dependencies: []string{"tailwindcss", "postcss"},
	}

	bootstrap = Bootstrap{
		popperJS:      gists.popperJS,
		bootstrapSCSS: gists.bootstrapSCSS,
		dependencies:  []string{"bootstrap", "@popperjs/core"},
	}
)

func main() {
	createTailwind := flag.Bool("t", false, "select tailwind as CSS")
	createBootstrap := flag.Bool("b", false, "select bootstrap as CSS")
	createVanilla := flag.Bool("v", false, "select vanilla CSS")

	flag.Parse()

	if len(flag.Args()) > MAXARGS {
		fmt.Println("Too many arguments provided! Please provide only at max 1")
		os.Exit(1)
	}

	// If there's a name, set it
	if len(flag.Args()) != 0 {
		projectName = flag.Args()[0]
	} else {
		projectName = "Default"
	}

	fmt.Println("Creating project", projectName+"!")
	cd := "cd " + projectName + " && "

	// Initialize Repo
	Exec("git init " + projectName)

	// Install all core dependencies
	{
		fullDependencyList := ""
		for _, d := range coreDependencies {
			if d != "parcel" {
				fullDependencyList += fullDependencyList + d + " "
			}
		}

		Exec(cd + "yard add --dev " + coreDependencies[0])
		Exec(cd + "yarn add " + fullDependencyList)
	}

	// Download main files
	{
		g := "git clone "
		Exec("mkdir src")
		Exec(g + gists.packageJSON)
		Exec(g + gists.indexHTML)
		Exec(g + gists.indexTS)
		Exec(g + gists.templateVUE)

		// Move to appropriate folders
		Exec("mkdir src/components")
		Exec("mv Template.vue ./src/components")
		Exec("mv index.ts ./src")
	}

	// Create CSS files
	{
		if *createTailwind {
			css.tailwind = tailwind
		} else if *createBootstrap {
			css.bootstrap = bootstrap
		} else if *createVanilla {
			Exec("touch src/index.scss")
			fmt.Println("Finished!")
		}
	}
}

// Executes commands
// Go Formatting string literal: https://stackoverflow.com/questions/17779371/golang-given-a-string-output-an-equivalent-golang-string-literal
// Also: https://groups.google.com/g/golang-nuts/c/ggd3ww3ZKcI
// And: https://www.digitalocean.com/community/tutorials/an-introduction-to-working-with-strings-in-go
func Exec(s string) {
	// Go Logging an error
	// https://www.honeybadger.io/blog/golang-logging/
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fmt.Println("Executing " + s + "...")
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
