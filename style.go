package main

import (
	"time"

	"github.com/fatih/color"
	"github.com/theckman/yacspin"
)

var (
	// Console colors
	Warn          = color.New(color.FgRed, color.Bold)
	Joy           = color.New(color.FgYellow, color.Bold)
	Action        = color.New(color.FgHiGreen, color.Bold, color.BgBlack)
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

func SpinWrap(sp *yacspin.Spinner, c int, s string, f func()) {
	sp.CharSet(yacspin.CharSets[c])
	sp.Prefix(s + " ")
	sp.Start()
	f()
	sp.Stop()
}
