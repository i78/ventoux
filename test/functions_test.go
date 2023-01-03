package test

import (
	"fmt"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestFunctions(t *testing.T) {
	t.Run("Simple hello world function should reuturn 'hello'", func(t *testing.T) {
		ast := readOrPanic(t, "../examples/functions/hello_fn.vx")

		fmt.Println(ast)
		captured, machine := machineWithStdoutCapture()

		machine.EvalProgram(ast)
		assert.Equal(t, "Hello Ventoux!", captured)
	})
}
