package main

// Gists to download
type Gists struct {
	packageJSON   string
	indexHTML     string
	indexTS       string
	templateVUE   string
	postcssrc     string
	popperJS      string
	bootstrapSCSS string
	tailwindSCSS  string
	tailwindconf  string
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
