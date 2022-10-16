package main

// TODO:
// - Add P5.JS vue support. That thing sucks to setup.

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/theckman/yacspin"
)

const (
	Maxargs = 3 // Max num of args in command
)

var (
	ProjectName  string
	spinner, err = yacspin.New(SpinnerConfig)
	cd           = "cd " + ProjectName + " && "
	Tmpdir       = "" // Access /tmp directory DOES NOT CURRENTLY WORK
	gists        = Gists{
		PackageJSON:   "https://gist.github.com/b6e6e41894e0d6b3ef7aba33214415ce.git",
		IndexHTML:     "https://gist.github.com/b9f38f17a0b2cf3f28d2715011e03fb1.git",
		IndexTS:       "https://gist.github.com/3bb7a5789c91bc0229dcbfe209f0fc67.git",
		TemplateVUE:   "https://gist.github.com/94f18653a1e0468de83faa163d7cdbcf.git",
		Gitignore:     "https://gist.github.com/d9b4506685d58cbe0ad715a55a922f3d.git",
		Postcssrc:     "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		PopperJS:      "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		BootstrapSCSS: "https://gist.github.com/b29599aa95343ad7ff3a704c0e9b2d81.git",
		TailwindSCSS:  "https://gist.github.com/35a637a7e185333b08c730b7d64189d3.git",
		Tailwindconf:  "https://gist.github.com/ed36d206090bd1faeea8d0c1921e19fc.git",
	}

	requiredGists = [5]string{
		gists.PackageJSON,
		gists.IndexHTML,
		gists.IndexTS,
		gists.TemplateVUE,
		gists.Gitignore,
	}

	coreDependencies = [4]string{
		"parcel",
		"vue",
		"vue-router@4",
		"rimraf",
	}

	tailwind = Tailwind{
		Postcssrc:    gists.Postcssrc,
		ConfigJS:     "",
		IndexSCSS:    gists.TailwindSCSS,
		Dependencies: []string{"tailwindcss", "postcss"},
	}

	bootstrap = Bootstrap{
		PopperJS:      gists.PopperJS,
		BootstrapSCSS: gists.BootstrapSCSS,
		Dependencies:  []string{"bootstrap", "@popperjs/core"},
	}

	css       CSS
	cssInsert = "html { scrollbar-gutter: stable both-edges; }\n"
)

func main() {
	// For spinner
	if err != nil {
		fmt.Println(err)
		return
	}

	createTailwind := flag.Bool("t", false, "select Tailwind as CSS")
	createBootstrap := flag.Bool("b", false, "select Bootstrap as CSS")
	createBulma := flag.Bool("u", false, "select Bulma CSS")
	createSimple := flag.Bool("s", false, "select Simple CSS")
	// createP5 := flag.Bool("p", false, "select P5 Project")
	createVanilla := flag.Bool("v", false, "select vanilla CSS")

	// Check all requirements for a command
	{
		if len(os.Args) > Maxargs || len(os.Args) == 1 {
			Warn.Println("Invalid number of arguments provided!")
			os.Exit(1)
		}

		// If there's a name, set it
		if len(os.Args) > 2 {
			ProjectName = os.Args[2]
		} else {
			ProjectName = "Default"
		}

		// Check if Directory exists
		// https://programming-idioms.org/idiom/212/check-if-folder-exists/3702/go
		info, err := os.Stat("./" + ProjectName)
		dirExists := !os.IsNotExist(err) && info.IsDir()

		if dirExists {
			Warn.Println("A directory already exists for " + "./" + ProjectName + " in current directory!")
			os.Exit(1)
		}
	}

	// Initialize Project Repository
	Execution.Println("Creating project " + ProjectName)
	if err := os.MkdirAll(ProjectName+"/src/components", 0750); err != nil {
		Warn.Println("Error occurred. Check logs stinky!")
		log.Fatal(err)
	}
	Exec("git init " + ProjectName)
	Action.Println("Git repo initialized!")

	// Install all core dependencies
	SpinWrap(
		spinner,
		11,
		" Installing core dependencies",
		func() {
			fullDependencyList := ""
			for _, d := range coreDependencies {
				if d != "parcel" {
					fullDependencyList += fullDependencyList + d + " "
				}
			}

			Exec(cd + "yarn add --dev " + coreDependencies[0])
			Exec(cd + "yarn add " + fullDependencyList)
		},
	)

	spinner.Frequency(45 * time.Millisecond)
	p := fmt.Sprintf("%s/", ProjectName)

	// Download and organize gists
	SpinWrap(
		spinner,
		43,
		" Downloading gists", // Maybe do with goroutines?
		func() {

			// IDs are the names of the downloaded folders

			clone, ids := gists.Clone(requiredGists[:])
			Exec(clone)

			// With IDs, move contents out of folders
			// I had done this before with Exec(), but this is more
			// idiomatic, and also simpler

			for _, id := range ids {
				files, _ := os.ReadDir(id)
				os.Rename(
					fmt.Sprintf("./%s/%s", id, files[1].Name()),
					fmt.Sprintf("./%s/%s", ProjectName, files[1].Name()),
				)

				// Then delete gist download folders
				if err := os.RemoveAll(id); err != nil {
					Warn.Println(err)
					os.Exit(1)
				}
			}
		},
	)

	spinner.Frequency(80 * time.Millisecond)
	Exec("cp " + p + "Template.vue " + p + "src/")

	// Move and organize downloaded files
	SpinWrap(
		spinner,
		27,
		" Moving things around",
		func() {
			Exec(
				func() string {

					var sb strings.Builder
					commands := [4]string{
						"mv " + p + "src/Template.vue " + p + "src/App.vue",
						"mv " + p + "Template.vue " + p + "src/components/",
						"mv " + p + "index.ts " + p + "src/",
						"mv " + p + "temp.gitignore " + p + `.gitignore`,
					}

					for _, cmd := range commands {
						sb.WriteString(cmd + " && ")
					}
					return sb.String()[0 : len(sb.String())-4]
				}(),
			)
		},
	)

	// Create CSS files
	SpinWrap(
		spinner,
		44,
		" Adding CSS",
		func() {
			flag.Parse()
			if *createTailwind {
				css.Tailwind = tailwind
			} else if *createBootstrap {
				css.Bootstrap = bootstrap
			} else if *createBulma {
				Exec("cd " + ProjectName + " && yarn add bulma")
				Insert("@import '~bulma';\n" + cssInsert)
			} else if *createSimple {
				Exec("cd " + ProjectName + " && yarn add simpledotcss")
				Insert("@import url('../node_modules/simpledotcss/simple.min.css');\n\n" + cssInsert)
			} else if *createVanilla {
				Insert(cssInsert)
			}
		},
	)

	// Initalize yarn
	SpinWrap(
		spinner,
		31,
		" Initalizing yarn",
		func() {
			Exec("cd " + ProjectName + " && yarn")
			Exec("cd " + ProjectName + " && touch README.md")
		},
	)

	Joy.Println("Environment setup complete")
}
