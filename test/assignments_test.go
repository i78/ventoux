package _test

import (
	machine2 "dreese.de/ventoux/internal/machine"
	"dreese.de/ventoux/internal/parser"
	"fmt"
	"github.com/alecthomas/assert/v2"
	"os"
	"testing"
)

func TestAssignments(t *testing.T) {
	if source, err := os.ReadFile("../examples/assign/assign.vx"); err == nil {
		ast, err := parser.GetParser().ParseString("", string(source))
		assert.NoError(t, err)
		tapped := ""
		tap := func(s string) { tapped = fmt.Sprintf("%s%s", tapped, s) }
		machine := machine2.NewMachine(tap)

		machine.EvalProgram(ast)

		assert.Equal(t, "Hello Variables!################HelloThis really works!This really works!3.141592", tapped)
	} else {
		t.Fatal("Unable to load source file")
	}
}
