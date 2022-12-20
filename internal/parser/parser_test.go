package parser

import (
	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Should parse helloworld", func(t *testing.T) {
		const testdata = "\"Hello Ventoux!\""
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		assert.NoError(t, err)
	})
}
