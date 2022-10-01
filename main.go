package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/theckman/yacspin"
)

// Add yacspin

const (
	MAXARGS = 1     // Max num of args in command
	TESTING = false // Are we in testing mode? proj file
)

var (
	css          CSS
	projectName  = "default"
	spinner, err = yacspin.New(SpinnerConfig)

	// Commands
	cd = "cd " + projectName + " && "

	gists = Gists{
		PackageJSON:   "https://gist.github.com/b6e6e41894e0d6b3ef7aba33214415ce.git",
		IndexHTML:     "https://gist.github.com/b9f38f17a0b2cf3f28d2715011e03fb1.git",
		IndexTS:       "https://gist.github.com/3bb7a5789c91bc0229dcbfe209f0fc67.git",
		TemplateVUE:   "https://gist.github.com/94f18653a1e0468de83faa163d7cdbcf.git",
		Postcssrc:     "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		PopperJS:      "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		BootstrapSCSS: "https://gist.github.com/b29599aa95343ad7ff3a704c0e9b2d81.git",
		TailwindSCSS:  "https://gist.github.com/35a637a7e185333b08c730b7d64189d3.git",
		Tailwindconf:  "https://gist.github.com/ed36d206090bd1faeea8d0c1921e19fc.git",
	}

	requiredGists = [4]string{
		gists.PackageJSON,
		gists.IndexHTML,
		gists.IndexTS,
		gists.TemplateVUE,
	}

	coreDependencies = [4]string{
		"parcel",
		"vue",
		"vue-router@4",
		"rimraf",
	}

	tailwind = Tailwind{
		postcssrc:    gists.Postcssrc,
		configJS:     "",
		indexSCSS:    gists.TailwindSCSS,
		dependencies: []string{"tailwindcss", "postcss"},
	}

	bootstrap = Bootstrap{
		popperJS:      gists.PopperJS,
		bootstrapSCSS: gists.BootstrapSCSS,
		dependencies:  []string{"bootstrap", "@popperjs/core"},
	}
)

func main() {

	if TESTING {
		TestingIsTrue.Println("TESTING is TRUE...")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	createTailwind := flag.Bool("t", false, "select tailwind as CSS")
	createBootstrap := flag.Bool("b", false, "select bootstrap as CSS")
	createVanilla := flag.Bool("v", false, "select vanilla CSS")

	flag.Parse()

	// Check all requirements for a command
	{
		// Too many args
		if len(flag.Args()) > MAXARGS {
			Warn.Println("🤪 Too many arguments provided silly! Please provide only at max 1")
			os.Exit(1)
		}

		// No args supplied
		noFlags := true
		flagCheck := []bool{*createTailwind, *createBootstrap, *createVanilla}

		for _, flag := range flagCheck {
			if flag {
				noFlags = false
			}
		}

		if noFlags {
			Warn.Println("You're stupid, dumb, and everything in-between. Sorry, not sorry.")
			fmt.Println("🤔 This is the usage:")
			fmt.Println("$ neos-vue-cli -[CSS flavor] [Project Name]")
			fmt.Println("    -t  Tailwindcss\n    -b  Bootstrap\n    -v  Vanilla")

			os.Exit(1)
		}

		// If there's a name, set it
		if len(flag.Args()) != 0 {
			projectName = flag.Args()[0]
		}
	}

	// Check if Directory exists and Initialize Repo
	// https://programming-idioms.org/idiom/212/check-if-folder-exists/3702/go
	{
		info, err := os.Stat("./" + projectName)
		dirExists := !os.IsNotExist(err) && info.IsDir()

		if dirExists {
			fmt.Println("A directory already exists for " + "./" + projectName + " in current directory!")
			os.Exit(1)
		} else {
			// Initialize Repo
			Joy.Println("Creating project " + projectName + "!")
			Exec("git init " + projectName)
		}
	}

	spinner.Start()

	// Install all core dependencies
	{
		fullDependencyList := ""
		for _, d := range coreDependencies {
			if d != "parcel" {
				fullDependencyList += fullDependencyList + d + " "
			}
		}

		Exec(cd + "yarn add --dev " + coreDependencies[0])
		Exec(cd + "yarn add " + fullDependencyList)
	}

	// Download + organize main files
	{
		p := fmt.Sprintf("./%s/", projectName)
		Exec(fmt.Sprintf("mkdir %ssrc", p))

		// Download gists
		clone, ids := gists.Clone(requiredGists[:])
		Exec(clone)

		// With ids, move contents out of folders
		Exec(
			func() string {
				final := ""
				for _, id := range ids {
					final += fmt.Sprintf("mv ./%s/* ", id) + p + " && "
				}
				return final[0 : len(final)-4]
			}(),
		)

		// Then delete gist download folders
		Exec(
			func() string {
				final := ""
				for _, id := range ids {
					final += "rm -rf " + fmt.Sprintf("./%s", id) + " && "
				}
				return final[0 : len(final)-4]
			}(),
		)

		// Commands to organize files
		Exec(
			func() string {
				final := ""
				commands := [5]string{
					"mkdir " + p + "src/components",
					"cp " + p + "Template.vue " + p + "src/",
					"mv " + p + "src/Template.vue " + p + "src/App.vue",
					"mv " + p + "Template.vue " + p + "src/components/",
					"mv " + p + "index.ts " + p + "src/",
				}

				for _, cmd := range commands {
					final += cmd + " && "
				}

				return final[0 : len(final)-4]
			}(),
		)
	}

	// Create CSS files
	{
		if *createTailwind {
			css.tailwind = tailwind
		} else if *createBootstrap {
			css.bootstrap = bootstrap
		} else if *createVanilla {
			Exec(fmt.Sprintf("touch ./%s/src/index.scss", projectName))
		}
		Exec(cd + "yarn")
	}

	spinner.Stop()

	Joy.Println("✅ Finished ✅")
	fmt.Println("Enjoy your project, I guess... I hate web development")
	testing()
}

func testing() {
	if TESTING {
		TestingIsTrue.Print("TESTING is TRUE...")
		Action.Println(" Deleting created directory ./" + projectName)
		Exec("rm -rf " + projectName)
		os.Exit(0)
	}
}
