package main

import (
	"strings"
)

// Gists to download
type Gists struct {
	PackageJSON, IndexHTML, IndexTS, RouterTS, SketchTS, TemplateVUE, AppVUE, SketchVUE, Gitignore, Postcssrc, PopperJS, BootstrapSCSS, TailwindSCSS, Tailwindconf string
}

// Clone retruns pointers to a command string and a slice of ids (folder names).
func (g *Gists) Clone(gists []string, tmp string) (*string, *[]string) {
	ids := []string{}
	gc := "git clone "

	for i, gist := range gists {
		gists[i] = gc + gist
		ids = append(ids, gist[24:len(gist)-4])
	}

	final := "cd " + tmp + " && " + strings.Join(gists, " && ")
	return &final, &ids
}

// CSS struct contains organizations for CSS data, like Tailwind and Bootstrap
type CSS struct {
	Tailwind  Tailwind
	Bootstrap Bootstrap
}

type Tailwind struct {
	Postcssrc    string
	ConfigJS     string
	IndexSCSS    string
	Dependencies []string
}

type Bootstrap struct {
	PopperJS      string
	BootstrapSCSS string
	Dependencies  []string
}
