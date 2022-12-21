package cli

import (
	"dreese.de/ventoux/internal/grammar"
	machine2 "dreese.de/ventoux/internal/machine"
	parser2 "dreese.de/ventoux/internal/parser"
	"encoding/json"
	"fmt"
	"github.com/alecthomas/kong"
	"log"
	"os"
)

type RunCommand struct {
	Sourcefile string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
	Verbose    bool
}

func (sv *RunCommand) Run(ctx *kong.Context) error {
	if content, err := os.ReadFile(sv.Sourcefile); err == nil {
		parser := parser2.GetParser()
		ast, _ := parser.ParseString("", string(content))

		// todo tidy
		machine := &machine2.Machine{Variables: map[string]*grammar.Value{}}

		for _, st := range ast.TopDec {
			machine.EvalTop(st)
		}

		if sv.Verbose {
			m, _ := json.Marshal(machine)
			fmt.Println("Machine Status: \n", string(m))
		}
	} else {
		log.Fatal(err)
	}

	return nil
}
