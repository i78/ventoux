package test

import (
	"fmt"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestOperators(t *testing.T) {
	t.Run("Additions", func(t *testing.T) {
		t.Run("Should return expected result", func(t *testing.T) {
			ast := readStringOrPanic(t, `1.0 + 1.0`) // todo should work without the spaces, too
			captured, machine := machineWithStdoutCapture()
			machine.EvalProgram(ast)
			assert.Equal(t, "2.000000", *captured)
		})
	})
	t.Run("Comparison", func(t *testing.T) {
		cases := []struct {
			code     string
			expected string
		}{
			{`1.0<2.0`, "true"},
			{`2.0>1.0`, "true"},
			{`1.0>=1.0`, "true"},
			{`1.0=1.0`, "true"},
			{`1.0= 1.0 = 1.0`, "true"},
			{`1.0= 1.0 = 2.0`, "false"},
			{`5 < 1 < 10`, "false"},
			{`5 < 9 < 10`, "true"},
			{`10 < 42 < 20`, "false"},
		}

		for _, testcase := range cases {
			t.Run(fmt.Sprint("should return ", testcase.expected, " for ", testcase.code), func(t *testing.T) {
				ast := readStringOrPanic(t, testcase.code)
				// repr.Print(ast)
				captured, machine := machineWithStdoutCapture()
				machine.EvalProgram(ast)
				assert.Equal(t, testcase.expected, *captured)
			})
		}
	})
}
