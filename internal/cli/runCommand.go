package cli

import (
	parser2 "dreese.de/ventoux/internal/parser"
	"fmt"
	"github.com/alecthomas/kong"
	"log"
	"os"
)

type RunCommand struct {
	Sourcefile string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
}

func (sv *RunCommand) Run(ctx *kong.Context) error {
	if content, err := os.ReadFile(sv.Sourcefile); err == nil {
		parser := parser2.GetParser()
		ast, _ := parser.ParseString("", string(content))

		for _, st := range ast.TopDec {
			fmt.Println(st.Literal)
		}
	} else {
		log.Fatal(err)
	}

	return nil
}
