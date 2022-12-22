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
		assert.Equal(t, "Hello Ventoux!", *ast.TopDec[0].ValueOrVariable.Value.StringValue)
		assert.Equal(t, "Hello!", *ast.TopDec[1].ValueOrVariable.Value.StringValue)
		assert.NoError(t, err)
	})

}

func TestAssign(t *testing.T) {
	t.Run("Should return assign node on assign example", func(t *testing.T) {
		const testdata = "greeting = \"Hello\"\n\"greeting\""
		ast, err := parser.ParseString("", testdata)
		assert.Equal(t, "Hello", *ast.TopDec[0].Assign.ValueOrVariable.Value.StringValue)
		assert.NoError(t, err)
	})
}

func TestExpressions(t *testing.T) {
	t.Run("Should return expected add expression", func(t *testing.T) {
		const testdata = "1 + 1"
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		expected := 1.0
		assert.Equal(t, &expected, ast.TopDec[0].Expression.Left.Value.NumberValue)
		assert.Equal(t, "+", ast.TopDec[0].Expression.Operator)
		assert.Equal(t, &expected, ast.TopDec[0].Expression.Right.Value.NumberValue)
		assert.NoError(t, err)
	})
}
