package parser

import (
	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
	"testing"
)

func TestLiteral(t *testing.T) {
	t.Run("Should return literal node on helloworld example", func(t *testing.T) {
		const testdata = "\"Hello Ventoux!\"\n\"Hello!\""
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		assert.Equal(t, "Hello Ventoux!", *ast.TopDec[0].Literal.Expression.Value.StringValue)
		assert.NoError(t, err)
	})

}

func TestAssign(t *testing.T) {
	t.Run("Should return assign node on assign example", func(t *testing.T) {
		const testdata = "greeting = \"Hello\"\n\"greeting\""
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		assert.Equal(t, "Hello", *ast.TopDec[0].Assign.Expression.Value.StringValue)
		assert.NoError(t, err)
	})
}
