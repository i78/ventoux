package test

import (
	machine2 "dreese.de/ventoux/internal/machine"
	"dreese.de/ventoux/internal/parser"
	"fmt"
	"github.com/alecthomas/assert/v2"
	"os"
	"testing"
)

func machineWithStdoutCapture() (*string, *machine2.Machine) {
	tapped := ""
	tap := func(s string) {
		tapped = fmt.Sprintf("%s%s", tapped, s)
	}
	machine := machine2.NewMachine(tap)
	return &tapped, machine
}

func readOrPanic(t *testing.T, filename string) *parser.Program {
	if source, err := os.ReadFile(filename); err == nil {
		ast, err := parser.ParseCode(string(source))
		assert.NoError(t, err)
		return ast
	} else {
		t.Fatal("unable to read source file")
		return nil
	}
}

func readStringOrPanic(t *testing.T, source string) *parser.Program {
	ast, err := parser.ParseCode(source)
	assert.NoError(t, err)
	return ast
}
