package parser

import (
	"fmt"
	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
	"testing"
)

type SourceAstTestcase struct {
	source      string
	expectedAst func(*Program)
}

func TestLiteral(t *testing.T) {
	for _, testcase := range []SourceAstTestcase{
		{
			source: "\"Hello Ventoux!\"",
			expectedAst: func(program *Program) {
				repr.Println(program)
				//assert.Equal(t, "Hello Ventoux!", *(program.TopDec[0].Expression).(grammar.ExprString).Value)
			},
		}, {
			source: "myVariable",
			expectedAst: func(program *Program) {
				repr.Println(program)
				//assert.Equal(t, "Hello Ventoux!", program.TopDec[0].Expression.(grammar.ExprString).Value)
			},
		},
	} {
		t.Run(fmt.Sprintf("Should return expected AST for %s", testcase.source), func(t *testing.T) {
			ast, err := parser.ParseString("", testcase.source)
			//repr.Println(ast)
			testcase.expectedAst(ast)
			assert.NoError(t, err)
		})
	}
}

func TestAssign(t *testing.T) {
	t.Run("Should return assign node on assign example", func(t *testing.T) {
		const testdata = "greeting = \"Hello\"\n\"greeting\""
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		// assert.Equal(t, "Hello", *ast.TopDec[0].Assign.ValueOrVariable.Value.StringValue)
		assert.NoError(t, err)
	})
}

func TestExpressions(t *testing.T) {
	t.Run("Should return expected add expression", func(t *testing.T) {
		const testdata = "1 + 1"
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		//expected := 1.0
		//assert.Equal(t, &expected, ast.TopDec[0].Expression.Left.Value.NumberValue)
		//assert.Equal(t, "+", ast.TopDec[0].Expression.Operator)
		//assert.Equal(t, &expected, ast.TopDec[0].Expression.Right.Value.NumberValue)
		assert.NoError(t, err)
	})

	t.Run("Should return expected syntax tree for 1+2+3", func(t *testing.T) {
		const testdata = "1 + 2 + 3 * ( 9 - 4 )"
		//const testdata = `a + b - c * d / e % f`
		ast, err := parser.ParseString("", testdata)
		repr.Println(ast)
		/*expected := 1.0
		assert.Equal(t, &expected, ast.TopDec[0].Expression.Left.Value.NumberValue)
		assert.Equal(t, "+", ast.TopDec[0].Expression.Operator)
		assert.Equal(t, &expected, ast.TopDec[0].Expression.Right.Value.NumberValue)*/
		assert.NoError(t, err)
	})
}
