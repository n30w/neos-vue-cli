package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// TODO: add dependencies like fatih/color and yacspin

const (
	MAXARGS = 1    // Max num of args in command
	TESTING = true // Are we in testing mode? proj file
)

var (
	css         CSS
	projectName string

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

	// gPoint *[4]string
)

func main() {
	createTailwind := flag.Bool("t", false, "select tailwind as CSS")
	createBootstrap := flag.Bool("b", false, "select bootstrap as CSS")
	createVanilla := flag.Bool("v", false, "select vanilla CSS")
	flag.Parse()
	// Check all requirements for a command
	{
		warn := color.New(color.FgRed, color.Bold)
		// Too many args
		if len(flag.Args()) > MAXARGS {
			warn.Println("🤪 Too many arguments provided silly! Please provide only at max 1")
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
			warn.Println("You're stupid, dumb, and everything in-between. Sorry, not sorry.")
			fmt.Println("🤔 This is the usage:")
			fmt.Println("$ neos-vue-cli -[CSS flavor] [Project Name]")
			fmt.Println("    -t  Tailwindcss\n    -b  Bootstrap\n    -v  Vanilla")

			os.Exit(1)
		}

		// If there's a name, set it
		if len(flag.Args()) != 0 {
			projectName = flag.Args()[0]
		} else {
			projectName = "Default"
		}
	}

	// Does a project by the same name already exist?
	// Check that

	// Initialize Repo
	Exec("git init " + projectName)
	d := color.New(color.FgYellow, color.Bold)
	d.Println("Creating project " + projectName + "!")
	cd := "cd " + projectName + " && "

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

		// Then delete folders
		Exec(
			func() string {
				final := ""
				for _, id := range ids {
					final += "rm -rf " + fmt.Sprintf("./%s", id) + " && "
				}
				return final[0 : len(final)-4]
			}(),
		)

		Exec(fmt.Sprintf("mkdir " + p + "src/components"))
		Exec("cp " + p + "Template.vue " + p + "src/")
		Exec("mv " + p + "src/Template.vue " + p + "src/App.vue")
		Exec("mv " + p + "Template.vue " + p + "src/components/")
		Exec("mv " + p + "index.ts " + p + "src/")
	}

	// Create CSS files
	{
		if *createTailwind {
			css.tailwind = tailwind
		} else if *createBootstrap {
			css.bootstrap = bootstrap
		} else if *createVanilla {
			Exec(fmt.Sprintf("touch ./%s/src/index.scss", projectName))
			fmt.Println("✅ Finished ✅")
			fmt.Println("Enjoy your project, I guess... I hate web development")
			// testing()
		}
	}

}

func testing() {
	if TESTING {
		t := color.New(color.BgRed, color.FgHiWhite, color.Bold)
		y := color.New(color.FgYellow, color.Bold)
		t.Print("TESTING is TRUE...")
		y.Println(" Deleting created directory ./" + projectName)
		Exec("rm -rf " + projectName)
		os.Exit(0)
	}
}
