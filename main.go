package main

import (
	"dreese.de/ventoux/internal/cli"
	"github.com/alecthomas/kong"
)

var CLI struct {
	Run cli.RunCommand `cmd:"" help:"Executes a Ventoux Script"`
}

// type verboseFlag bool
/*
                           888
                           888
                           888
888  888  .d88b.  88888b.  888888 .d88b.  888  888 888  888
888  888 d8P  Y8b 888 "88b 888   d88""88b 888  888 `Y8bd8P'
Y88  88P 88888888 888  888 888   888  888 888  888   X88K
 Y8bd8P  Y8b.     888  888 Y88b. Y88..88P Y88b 888 .d8""8b.
  Y88P    "Y8888  888  888  "Y888 "Y88P"   "Y88888 888  888
*/
func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
