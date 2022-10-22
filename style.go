package main

import (
	"time"

	"github.com/fatih/color"
	"github.com/theckman/yacspin"
)

// TODO: Possibly add Uilive pkg to it?
// can display live output line by line via carriage return

var (
	Warn          = color.New(color.FgRed, color.Bold)
	Joy           = color.New(color.FgYellow, color.Bold)
	Action        = color.New(color.FgHiGreen, color.Bold)
	TestingIsTrue = color.New(color.BgRed, color.FgHiWhite, color.Bold)
	Execution     = color.New(color.FgMagenta, color.Bold)

	SpinnerConfig = yacspin.Config{
		Frequency:     100 * time.Millisecond,
		Prefix:        " ",
		Message:       "",
		SpinnerAtEnd:  true,
		StopCharacter: "✓",
		StopColors:    []string{"fgGreen"},
	}
)

func SpinWrap(spinner *yacspin.Spinner, spinnerType int, message string, f func()) {
	spinner.CharSet(yacspin.CharSets[spinnerType])
	spinner.Prefix(message + " ")
	spinner.Start()
	f()
	spinner.Stop()
}
