package main

// yarn add vue rimraf
// yarn add --dev parcel

type Dependencies struct {
	rimraf    string
	tailwind  string
	bootstrap string
	vue       string
	parcel    string
}

type Tailwind struct {
	postcssrc string
	configJS  string
	indexSCSS string
}

type Bootstrap struct {
	popperJS         string
	bootstrapIndexTS string
}

// Gists to download
type Gists struct {
	packageJSON string
	indexHTML   string
	indexTS     string
	templateVUE string
	tailwind    Tailwind
	bootstrap   Bootstrap
}

func (g *Gists) DownloadGists() {

}

func (g *Gists) Init() {

}
