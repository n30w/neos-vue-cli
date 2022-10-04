# Neo's Vue CLI

This is my CLI tool to setup my custom development environment. Make sure you have upx installed. Install with: `brew install --build-from-source upx`
per [This github thread](https://gist.github.com/somecuitears/c6b88944e3e59c598a57812084d373d6).

Build with `make build`

## Links

- [Shrinking a Go Binary](https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)
  - `go build -ldflags="-s -w"` strips the Binary
  - `brew install --build-from-source upx` install upx with brew
  - `upx --brute (binary name)` use upx to shrink it more
- [Creating an executable](https://stackoverflow.com/questions/28018591/how-do-i-make-a-ruby-script-into-a-bash-command)
- [Creating a makefile for Go](https://earthly.dev/blog/golang-makefile/)
- [Getting started with generics in Go](https://www.infoworld.com/article/3646036/get-started-with-generics-in-go.html)
- [String Builder](https://www.calhoun.io/concatenating-and-building-strings-in-go/)
