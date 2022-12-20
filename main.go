package main

import (
	"dreese.de/ventoux/internal/cli"
	"github.com/alecthomas/kong"
)

var CLI struct {
	Run cli.RunCommand `cmd:"" help:"Executes a Ventoux Script"`
}

// type verboseFlag bool

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
