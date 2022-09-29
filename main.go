package main

import (
	"flag"
	"fmt"
	"os"
)

// TODO: add dependencies like fatih/color and yacspin

// Tutorial: https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f

/*

This will be a program dedicated to creating the environment
required to setup a NeosWebDev.

It is a command line tool

The course of action follows:

1. run neosvuecli create [options]
	- options are:
		- (REQUIRED) name : name of project
		- directory : if not stated it will just create one in the current
		  directory
		- private : sets whether you want the project to be private
		- license : type of license
		- t : add tailwind
		- b : add bootstrap


Things that the command should do:

git init (name)

cd (name)

yarn add --dev parcel && yarn add vue (https://classic.yarnpkg.com/lang/en/docs/cli/init/)

mkdir src && mkdir src/components

download gists

mv gists

add tailwind or bootstrap

touch tsconfig.json

git add .

git commit -m "initial commit"

git push origin main

finish!

*/

var (
	css         CSS
	projectName string

	gists = Gists{
		packageJSON: "https://gist.github.com/b6e6e41894e0d6b3ef7aba33214415ce.git",
		indexHTML:   "https://gist.github.com/b9f38f17a0b2cf3f28d2715011e03fb1.git",
		indexTS:     "https://gist.github.com/3bb7a5789c91bc0229dcbfe209f0fc67.git",
		templateVUE: "https://gist.github.com/94f18653a1e0468de83faa163d7cdbcf.git",
		postcssrc:   "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		popperJS:    "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
	}

	commands = Commands{
		git:   []string{"init", "add", "commit", "-m", "inital commit", "push"},
		mkdir: []string{"src", "src/components"},
		// touch: []string{"src/tsconfig.json"},
		cd: []string{"src"},
		// mv:    []string{"-t", "./src index.ts index.scss", "Template.vue src/components/"},
		yarn: []string{"yarn add --dev ", "yarn add "},
	}

	coreDependencies = [4]string{
		"parcel",
		"vue",
		"vue-router@4",
		"rimraf",
	}

	// destinations = map[string]string{
	// 	"src":               "src",
	// 	"src/components":    "src/components",
	// 	"src/tsconfig.json": "src/tsconfig.json",
	// 	"./src":             "./src",
	// }

	tailwind = Tailwind{
		postcssrc:    gists.postcssrc,
		configJS:     "",
		indexSCSS:    "@tailwind base;@tailwind components;@tailwind utilities;",
		dependencies: []string{"tailwindcss", "postcss"},
	}

	bootstrap = Bootstrap{
		popperJS:      gists.popperJS,
		bootstrapSCSS: "@import '../node_modules/bootstrap/scss/functions';@import '../node_modules/bootstrap/scss/variables';@import '../node_modules/bootstrap/scss/mixins';@import '../node_modules/bootstrap/scss/root';@import '../node_modules/bootstrap/scss/reboot';@import '../node_modules/bootstrap/scss/type';@import '../node_modules/bootstrap/scss/images';@import '../node_modules/bootstrap/scss/containers';@import '../node_modules/bootstrap/scss/grid';",
		dependencies:  []string{"bootstrap", "@popperjs/core"},
	}
)

// func notEnoughArgs() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Expected 'create' subcommand!")
// 		os.Exit(1)
// 	}
// }

func init() {

}

func main() {
	const MAXARGS = 1
	fmt.Println("Initializing...")

	// createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createTailwind := flag.Bool("t", false, "select tailwind as CSS")
	createBootstrap := flag.Bool("b", false, "select bootstrap as CSS")
	createVanilla := flag.Bool("v", false, "select vanilla CSS")

	flag.Parse()
	fmt.Println(*createVanilla)
	if len(flag.Args()) > MAXARGS {
		fmt.Println("Too many arguments provided! Please provide only at max 1")
		os.Exit(1)
	}

	if len(flag.Args()) != 0 {
		projectName = flag.Args()[0]
	} else {
		projectName = "default"
	}
	fmt.Println(projectName)
	Exec("mkdir watda")
	pwd, _ := os.Getwd()
	// Initialize Repo
	Exec("git init " + projectName)
	Exec("cd " + projectName)
	fmt.Println(pwd)
	// Install all core dependencies
	{
		fullDependencyList := ""
		for _, d := range coreDependencies {
			if d != "parcel" {
				fullDependencyList += fullDependencyList + d + " "
			}
		}

		Exec("yard add --dev " + coreDependencies[0])
		Exec("yarn add " + fullDependencyList)
	}

	// Download gists
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
	// createCmd.Parse(os.Args[2:])
	if *createTailwind {
		css.tailwind = tailwind
		//write this to tailwind.config.js
		//content: [
		//"./src/**/*.{html, js, ts, jsx, tsx, vue}",
		//"./src/*.vue",
		//"./src/components/*.vue",
		//],

	} else if *createBootstrap {
		css.bootstrap = bootstrap
		// write this to index.ts
		// import * as bootstrap from 'bootstrap';

	} else if *createVanilla {
		Exec("touch src/index.scss")
		fmt.Println("Finished!")
	}

	// switch os.Args[1] {
	// case "create":
	// case "help":
	// 	fmt.Println("Usage: neosvuecli (create | help) (-t | -b) (project name)")
	// 	os.Exit(1)
	// default:
	// 	notEnoughArgs()
	// }
}
