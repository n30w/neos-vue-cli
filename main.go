package main

import "fmt"

// TODO: add dependencies like fatih/color and yacspin

// Tutorial: https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f

// Please bundle add file templates like .json and .vue and .html to binary.

/*

This will be a program dedicated to creating the environment
required to setup a NeosWebDev.

It is a command line tool

The course of action follows:

1. run neosvuecli create [options]
	- options are:
		- (REQUIRED) name : name of project
		- directory : if not stated it will just create one in the current
		  directory
		- private : sets whether you want the project to be private
		- license : type of license
		- t : add tailwind
		- b : add bootstrap


Things that the command should do:

git init (name)

cd (name)

yarn init --yes

or to initalize a skip and private project
yarn init -yp

yarn add --dev parcel && yarn add vue (https://classic.yarnpkg.com/lang/en/docs/cli/init/)

add tailwind or bootstrap

touch tsconfig.json

touch index.html

write to index.html

mkdir src && cd src

touch App.vue index.scss index.ts

write to App.vue and write to index.ts

mkdir components && touch Temp.vue

write to Temp.vue (Temp.vue have same content as App.vue)

git add .

git commit -m "initial commit"

git push origin main

finish!

*/

func main() {
	fmt.Println("vim-go")
}
