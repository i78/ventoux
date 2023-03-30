package grammar

import (
	"fmt"
)

type (
	ExprIdent struct {
		Name string `parser:"@Ident"`
	}

	ExprParens struct {
		Inner ExprPrecAll `parser:"'(' @@ ')'"`
	}

	ExprUnary struct {
		Op   string      `parser:"@('-' | '!')"`
		Expr ExprOperand `parser:"@@"`
	}

	ExprRem struct {
		Head ExprOperand  `parser:"@@"`
		Tail []ExprRemExt `parser:"@@+"`
	}

	ExprRemExt struct {
		Op   string      `parser:"@'%'"`
		Expr ExprOperand `parser:"@@"`
	}

	ExprPrecAll interface{ exprPrecAll() }
	ExprPrec2   interface{ exprPrec2() }
	ExprPrec3   interface{ exprPrec3() }
	ExprOperand interface{ exprOperand() }

	ExprEvaluatable interface{ Evaluate(*Variables) *Expression }
	ExprTerminal    interface{ Terminal() *Expression }
)

// These expression types can be matches as individual operands
func (ExprIdent) exprOperand()    {}
func (ExprNumber) exprOperand()   {}
func (ExprBoolean) exprOperand()  {}
func (ExprString) exprOperand()   {}
func (ExprParens) exprOperand()   {}
func (ExprUnary) exprOperand()    {}
func (ExprFnCall) exprOperand()   {}
func (ExprFunction) exprOperand() {}

// These expression types can be matched at precedence level 3
func (ExprIdent) exprPrec3()   {}
func (ExprNumber) exprPrec3()  {}
func (ExprBoolean) exprPrec3() {}
func (ExprString) exprPrec3()  {}
func (ExprParens) exprPrec3()  {}
func (ExprUnary) exprPrec3()   {}
func (ExprRem) exprPrec3()     {}
func (ExprFnCall) exprPrec3()  {}

// These expression types can be matched at precedence level 2
func (ExprIdent) exprPrec2()      {}
func (ExprNumber) exprPrec2()     {}
func (ExprBoolean) exprPrec2()    {}
func (ExprString) exprPrec2()     {}
func (ExprParens) exprPrec2()     {}
func (ExprUnary) exprPrec2()      {}
func (ExprRem) exprPrec2()        {}
func (ExprComparison) exprPrec2() {}
func (ExprMulDiv) exprPrec2()     {}
func (ExprPow) exprPrec2()        {}
func (ExprBitshift) exprPrec2()   {}
func (ExprFnCall) exprPrec2()     {}

// These expression types can be matched at the minimum precedence level
func (ExprIdent) exprPrecAll()      {}
func (ExprNumber) exprPrecAll()     {}
func (ExprBoolean) exprPrecAll()    {}
func (ExprString) exprPrecAll()     {}
func (ExprParens) exprPrecAll()     {}
func (ExprUnary) exprPrecAll()      {}
func (ExprRem) exprPrecAll()        {}
func (ExprMulDiv) exprPrecAll()     {}
func (ExprComparison) exprPrecAll() {}
func (ExprPow) exprPrecAll()        {}
func (ExprBitshift) exprPrecAll()   {}
func (ExprAddSub) exprPrecAll()     {}
func (ExprFnCall) exprPrecAll()     {}
func (ExprFunction) exprPrecAll()   {}

type Expression struct {
	X ExprPrecAll `@@`
}

func (e *Expression) ToString() string {
	if it, ok := e.X.(ExprString); ok {
		return it.Value
	} else if it, ok := e.X.(ExprNumber); ok {
		return fmt.Sprintf("%f", it.Value)
	} else if it, ok := e.X.(ExprBoolean); ok {
		return fmt.Sprintf("%t", it.Value)
	}
	return ""
}

func (eid ExprIdent) Evaluate(variables *Variables) *Expression {
	if it, exists := (*variables)[eid.Name]; exists {
		return it.Evaluate(variables)
	}
	// Else, return lazy lets talk about it.
	return &Expression{X: eid}
}

// todo type
func (e Expression) Evaluate(variables *Variables) *Expression {
	if it, isTerminal := e.X.(ExprTerminal); isTerminal {
		return it.Terminal()
	} else if it, isEvaluable := e.X.(ExprEvaluatable); isEvaluable {
		return it.Evaluate(variables)
	}
	return nil
}

func (ep ExprParens) Evaluate(variables *Variables) *Expression {
	return ep.Inner.(ExprEvaluatable).Evaluate(variables)
}
