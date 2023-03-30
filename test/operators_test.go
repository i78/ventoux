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
			{`1.0 < 2.0`, "true"},        // todo should work without the spaces, too
			{`2.0 > 1.0`, "true"},        // todo should work without the spaces, too
			{`1.0 >= 1.0`, "true"},       // todo should work without the spaces, too
			{`1.0 = 1.0`, "true"},        // todo should work without the spaces, too
			{`1.0 = 1.0 = 1.0`, "true"},  // todo should work without the spaces, too
			{`1.0 = 1.0 = 2.0`, "false"}, // todo should work without the spaces, too
			{`5 < 1 < 10`, "false"},      // todo should work without the spaces, too
			{`5 < 9 < 10`, "true"},       // todo should work without the spaces, too
			{`10 < 42 < 20`, "false"},    // todo should work without the spaces, too
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
