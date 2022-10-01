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
	Action        = color.New(color.FgYellow, color.Bold)
	TestingIsTrue = color.New(color.BgRed, color.FgHiWhite, color.Bold)
	Execution     = color.New(color.FgMagenta, color.Bold)

	SpinnerConfig = yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[27],
		Suffix:          "doing its thang",
		SuffixAutoColon: true,
		StopCharacter:   "✓",
		StopColors:      []string{"fgGreen"},
	}
)
