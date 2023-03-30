package test

import (
	machine2 "dreese.de/ventoux/internal/machine"
	"dreese.de/ventoux/internal/parser"
	"fmt"
	"github.com/alecthomas/assert/v2"
	"os"
	"testing"
)

func TestConditionals(t *testing.T) {
	t.Run("Conditionals Example should return expected result", func(t *testing.T) {
		if source, err := os.ReadFile("../examples/expressions/conditionals.vx"); err == nil {
			ast, err := parser.GetParser().ParseString("", string(source))
			assert.NoError(t, err)
			tapped := ""
			tap := func(s string) { tapped = fmt.Sprintf("%s%s", tapped, s) }
			machine := machine2.NewMachine(tap)

			machine.EvalProgram(ast)

			assert.Equal(t, "average", tapped)
		} else {
			t.Fatal("Unable to load source file")
		}
	})

	cases := []struct {
		description string
		code        string
		expected    string
	}{
		{"Should return nothing on empty bracket",
			`{}`,
			""},
		{"Should return default option when nothing else provided",
			`{
				-> 23
			}`,
			"23.000000"},
		{"Should return first matching terminal option",
			`{
	               true 	-> 42
	   						-> 23
	   		}`,
			"42.000000"},
		{"Should return first matching variable option",
			`
			k = true
	   		{
	               k 	-> 42
	   					-> 23
	   		}`,
			"42.000000"},
		{"Should return first matching mapping style option",
			`
			prefix = 33
			{
				prefix = 49 → "germany"
				prefix = 33	→ "france"
			}`,
			"france"},
		{"Should return first matching evaluable option",
			`
			k = 42
			{
	        	k < 10		→ 42
	            k < 20		→ 43
	   			k < 100 	→ 1
	   						→ 23
	   		}`,
			"1.000000"},
		{"Should also respect ranges",
			`
			k = 42
			{
				10 < k < 20		→ 42
	            50 < k < 100 	→ 43	
	   							→ 23
	   		}`,
			"23.000000"},
	}

	for _, testcase := range cases {
		t.Run(testcase.description, func(t *testing.T) {
			ast := readStringOrPanic(t, testcase.code)
			captured, machine := machineWithStdoutCapture()
			machine.EvalProgram(ast)
			assert.Equal(t, testcase.expected, *captured)
		})
	}
}
