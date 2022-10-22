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
	TmpDir       string
	insertText   string
	css          CSS
	spinner, err = yacspin.New(SpinnerConfig)
	cd           = "cd " + ProjectName + " && "
	cmd          = ""
	cssInsert    = "html { scrollbar-gutter: stable both-edges; }\n"
	gists        = &Gists{
		PackageJSON:   "https://gist.github.com/b6e6e41894e0d6b3ef7aba33214415ce.git",
		IndexHTML:     "https://gist.github.com/b9f38f17a0b2cf3f28d2715011e03fb1.git",
		IndexTS:       "https://gist.github.com/3bb7a5789c91bc0229dcbfe209f0fc67.git",
		RouterTS:      "https://gist.github.com/dfc984c549436801c8baa0f81509d58a.git",
		SketchTS:      "https://gist.github.com/377dc488457d4559027bd86b0cb8c293.git",
		TemplateVUE:   "https://gist.github.com/94f18653a1e0468de83faa163d7cdbcf.git",
		AppVUE:        "https://gist.github.com/fbce0f55f9d6f2d220a92de6b12b415e.git",
		SketchVUE:     "https://gist.github.com/04cdaab8de42293f08cdb43aa1c6a2a4.git",
		Gitignore:     "https://gist.github.com/d9b4506685d58cbe0ad715a55a922f3d.git",
		Postcssrc:     "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		PopperJS:      "https://gist.github.com/202c5d3b1bb088b615a243690124a3bd.git",
		BootstrapSCSS: "https://gist.github.com/b29599aa95343ad7ff3a704c0e9b2d81.git",
		TailwindSCSS:  "https://gist.github.com/35a637a7e185333b08c730b7d64189d3.git",
		Tailwindconf:  "https://gist.github.com/ed36d206090bd1faeea8d0c1921e19fc.git",
	}

	requiredGists = [7]string{
		gists.PackageJSON,
		gists.IndexHTML,
		gists.IndexTS,
		gists.RouterTS,
		gists.TemplateVUE,
		gists.AppVUE,
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

	// =========================================================================== //
	// CHECK ALL REQUIREMENTS FOR A COMMAND
	// =========================================================================== //
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
			Warn.Println("A directory already exists for " + ProjectName + " in current directory!")
			os.Exit(1)
		}
	}

	// =========================================================================== //
	// INITIALIZE PROJECT REPOSITORY
	// =========================================================================== //

	Execution.Println("Creating project " + ProjectName)
	if err := os.MkdirAll(ProjectName+"/src/components", 0750); err != nil {
		Warn.Println("Error occurred. Check logs stinky!")
		log.Fatal(err)
	}
	cmd = "git init " + ProjectName
	Exec(&cmd)
	Action.Println("Git repo initialized!")

	// =========================================================================== //
	// INSTALL ALL CORE DEPENDENCIES
	// =========================================================================== //

	SpinWrap(
		spinner,
		11,
		" Installing core dependencies",
		func() {
			var sb strings.Builder
			for _, d := range coreDependencies {
				if d != "parcel" {
					sb.WriteString(d + " ")
				}
			}

			cmd = cd + "yarn add --dev " + coreDependencies[0] + " " + cd + "yarn add " + sb.String()
			Exec(&cmd)
		},
	)

	spinner.Frequency(45 * time.Millisecond)
	p := fmt.Sprintf("%s/", ProjectName)

	// =========================================================================== //
	// DOWNLOAD AND ORGANIZE ALL GISTS
	// =========================================================================== //

	TmpDir, err := os.MkdirTemp("", "neosvuecli")
	if err != nil {
		os.RemoveAll(TmpDir)
		log.Fatal(err)
	}

	// Allow reading and writing into tmpDir
	cmd = "chmod 777 " + TmpDir
	Exec(&cmd)

	defer os.RemoveAll(TmpDir)

	SpinWrap(
		spinner,
		43,
		" Downloading gists",
		func() {
			// IDs are the names of the downloaded folders
			var clone *string
			var ids *[]string
			clone, ids = gists.Clone(requiredGists[:], TmpDir)
			Exec(clone)
			for _, id := range *ids {
				files, _ := os.ReadDir(TmpDir + "/" + id)
				err := os.Rename(
					fmt.Sprintf("%s/%s/%s", TmpDir, id, files[1].Name()),
					fmt.Sprintf("./%s/%s", ProjectName, files[1].Name()),
				)

				if err != nil {
					os.RemoveAll(TmpDir)
					log.Fatal(err)
				}

				// Delete gist download folders
				if err := os.RemoveAll(id); err != nil {
					Warn.Println(err)
					os.Exit(1)
				}
			}
		},
	)

	spinner.Frequency(80 * time.Millisecond)

	// =========================================================================== //
	// MOVE AND ORGANIZE DOWNLOADED FILES
	// =========================================================================== //

	SpinWrap(
		spinner,
		27,
		" Moving things around",
		func() {
			Exec(
				func() *string {
					commands := [5]string{
						"mv " + p + "App.vue " + p + "src/",
						"mv " + p + "Template.vue " + p + "src/components/",
						"mv " + p + "index.ts " + p + "src/",
						"mv " + p + "router.ts " + p + "src/",
						"mv " + p + "temp.gitignore " + p + `.gitignore`,
					}
					final := strings.Join(commands[:], " && ")
					return &final
				}(),
			)
		},
	)

	// =========================================================================== //
	// CREATE CSS FILES
	// =========================================================================== //

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
				cmd = cd + "yarn add bulma"
				Exec(&cmd)
				insertText = "@import '~bulma';\n" + cssInsert
				Insert(&insertText)
			} else if *createSimple {
				cmd = cd + "yarn add simpledotcss"
				Exec(&cmd)
				insertText = "@import url('../node_modules/simpledotcss/simple.min.css');\n\n" + cssInsert
				Insert(&insertText)
			} else if *createVanilla {
				Insert(&cssInsert)
			}
		},
	)

	// =========================================================================== //
	// INITIALIZE YARN
	// =========================================================================== //
	SpinWrap(
		spinner,
		31,
		" Initializing yarn",
		func() {
			cmd = cd + "yarn && " + cd + "touch README.md"
			Exec(&cmd)
		},
	)

	Joy.Print("Environment setup complete")
}
