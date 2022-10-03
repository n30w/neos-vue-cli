package main

import "fmt"

// Gists to download
type Gists struct {
	PackageJSON, IndexHTML, IndexTS, TemplateVUE, Gitignore, Postcssrc, PopperJS, BootstrapSCSS, TailwindSCSS, Tailwindconf string
}

func (g *Gists) Clone(gists []string) (string, []string) {
	final := ""
	ids := []string{}
	gc := "git clone " // TODO: Make this thing a temp directory via os pkg
	for _, gist := range gists {
		final += gc + fmt.Sprintf("%s && ", gist)
		ids = append(ids, gist[24:len(gist)-4])
	}
	return final[0 : len(final)-4], ids
}

// Organizations for CSS data, like Tailwind and Bootstrap
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
