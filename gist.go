package main

// yarn add vue rimraf
// yarn add --dev parcel

// Gists to download
type Gists struct {
	packageJSON string
	indexHTML   string
	indexTS     string
	templateVUE string
	postcssrc   string
	popperJS    string
}

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
