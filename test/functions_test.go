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
}
