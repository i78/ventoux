package machine

import (
	"dreese.de/ventoux/internal/grammar"
	parser2 "dreese.de/ventoux/internal/parser"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestEvalExpression(t *testing.T) {

	t.Run("should successfully add two constants", func(t *testing.T) {
		parser := parser2.GetParser()
		ast, _ := parser.ParseString("", "8 + 16")
		m := Machine{}
		result := m.EvalExpr(ast.TopDec[0].Expression)
		assert.Equal(t, 24, result.X.(grammar.ExprNumber).Value)
	})

	t.Run("should successfully add a constant and a variable value", func(t *testing.T) {
		parser := parser2.GetParser()
		ast, _ := parser.ParseString("", "a=8\na + 8")
		m := Machine{
			grammar.Variables{},
		}
		m.EvalTop(ast.TopDec[0])
		result := m.EvalExpr(ast.TopDec[1].Expression)
		assert.Equal(t, 16, result.X.(grammar.ExprNumber).Value)
	})

}

func TestSaveVmState(t *testing.T) {
	t.Run("should save even empty machine state when requested", func(t *testing.T) {
		m := Machine{}
		state := m.ExportMachineState()
		assert.NotZero(t, state)
	})

	t.Run("should save machine state with variables when requested", func(t *testing.T) {
		m := Machine{
			map[string]*grammar.Expression{"dummy": {X: grammar.ExprString{Value: "hello!"}}},
		}
		state := m.ExportMachineState()
		assert.NotZero(t, state)
	})
}
