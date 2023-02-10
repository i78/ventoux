package test

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestFunctions(t *testing.T) {
	t.Run("Simple hello world function should return 'hello'", func(t *testing.T) {
		ast := readOrPanic(t, "../examples/functions/hello_fn.vx")
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "Hello FnVentoux!", *captured)
	})

	t.Run("Should return expected result for adder function", func(t *testing.T) {
		ast := readOrPanic(t, "../examples/functions/adder_fn.vx")
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "10.000000", *captured)
	})

	t.Run("Should return expected result for nested adder function", func(t *testing.T) {
		ast := readOrPanic(t, "../examples/functions/adder_fn_nested.vx")
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "20.000000", *captured)
	})

	t.Run("Should return expected result for curried adder function", func(t *testing.T) {
		ast := readOrPanic(t, "../examples/functions/adder_fn_curried.vx")
		captured, machine := machineWithStdoutCapture()
		machine.EvalProgram(ast)
		assert.Equal(t, "43.000000", *captured)
	})
}
