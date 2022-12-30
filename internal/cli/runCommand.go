package cli

import (
	"dreese.de/ventoux/internal/grammar"
	machine2 "dreese.de/ventoux/internal/machine"
	parser2 "dreese.de/ventoux/internal/parser"
	"encoding/json"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/alecthomas/repr"
	"log"
	"os"
)

type RunCommand struct {
	Sourcefile string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
	PrintAst   bool
	Verbose    bool
}

func (sv *RunCommand) Run(ctx *kong.Context) error {
	if content, err := os.ReadFile(sv.Sourcefile); err == nil {
		parser := parser2.GetParser()
		ast, err := parser.ParseString("", string(content))

		if err != nil {
			panic(err)
		}

		if sv.PrintAst {
			repr.Print(ast)
		}

		// todo tidy
		machine := &machine2.Machine{Variables: map[string]*grammar.Expression{}}

		for _, st := range ast.TopDec {
			res := machine.EvalTop(st)
			if res != nil {
				fmt.Println(res.ToString())
			}
		}

		if sv.Verbose {
			m, _ := json.MarshalIndent(machine, "", "  ")
			fmt.Println("Machine Status: \n", string(m))
		}
	} else {
		log.Fatal(err)
	}

	return nil
}
