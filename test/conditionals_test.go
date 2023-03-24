package test

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestConditionals(t *testing.T) {
	t.Run("Should return nothing on empty bracket", func(t *testing.T) {
		ast := readStringOrPanic(t, `{}`)
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "", *captured)
	})

	t.Run("Should return default option when nothing else provided", func(t *testing.T) {
		ast := readStringOrPanic(t, `{
			-> 23
		}`)
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "23.000000", *captured)
	})

	t.Run("Should return first matching terminal option", func(t *testing.T) {
		ast := readStringOrPanic(t, `{
            true 	-> 42
					-> 23
		}`)
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "42.000000", *captured)
	})

	t.Run("Should return first matching variable option", func(t *testing.T) {
		ast := readStringOrPanic(t, `
		k = true
		{
            k 	-> 42
				-> 23
		}`)
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "42.000000", *captured)
	})
}
