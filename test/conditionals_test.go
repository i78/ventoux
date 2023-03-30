package test

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestConditionals(t *testing.T) {
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
