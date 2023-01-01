package cli

import (
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
	Sourcefile                string `arg:"" optional:"" name:"path" help:"Path to Ventoux script to execute." type:"path"`
	ExportVirtualMachineState string `optional:"" help:"Save the VM state to after program execution." type:"path"`
	PrintAst                  bool   `help:"Print the abstract syntax tree of the given Ventoux script to stdout before running"`
	Verbose                   bool   `help:"Provide additional information"`
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

		machine := machine2.NewMachine(func(s string) {
			fmt.Println(s)
		})

		machine.EvalProgram(ast)

		if sv.Verbose {
			m, _ := json.MarshalIndent(machine, "", "  ")
			fmt.Println("Machine Status: \n", string(m))
		}

		if sv.ExportVirtualMachineState != "" {
			state := machine.ExportMachineState()
			if err := os.WriteFile(sv.ExportVirtualMachineState, state, 0766); err != nil {
				panic("Unable to persist virtual machine state!")
			}
		}
	} else {
		log.Fatal(err)
	}

	return nil
}
