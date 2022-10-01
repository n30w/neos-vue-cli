package main

import "fmt"

// Gists to download
type Gists struct {
	PackageJSON, IndexHTML, IndexTS, TemplateVUE, Postcssrc, PopperJS, BootstrapSCSS, TailwindSCSS, Tailwindconf string
}

// If I had a repo struct, I could make an interface for this clone function then have both Repo and Gist implement this method
func (g *Gists) Clone(gists []string) (string, []string) {
	final := ""
	ids := []string{}
	gc := "git clone "
	for _, gist := range gists {
		final += gc + fmt.Sprintf("%s && ", gist)
		ids = append(ids, gist[24:len(gist)-4])
	}
	return final[0 : len(final)-4], ids
}

// Organizations for CSS data, like Tailwind and Bootstrap
type CSS struct {
	tailwind  Tailwind
	bootstrap Bootstrap
}

type Tailwind struct {
	postcssrc    string
	configJS     string
	indexSCSS    string
	dependencies []string
}

type Bootstrap struct {
	popperJS      string
	bootstrapSCSS string
	dependencies  []string
}
