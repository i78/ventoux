package grammar

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
)

type TopDec struct {
	Pos             lexer.Position
	Assign          *Assign          `parser:"@@"`
	Expression      *Expression      `parser:"| @@"`
	ValueOrVariable *ValueOrVariable // `parser:"| @@"`
}

type Operator string

type Assign struct {
	Pos        lexer.Position
	Left       string      `parser:"@Ident '='"`
	Expression *Expression `parser:"@@"`
	//Expression      *Expression      `parser:"[@@"`
	//ValueOrVariable *ValueOrVariable `parser:"| @@]"`
	// Value      Value           `parser:"@@"`
}

type Expr struct {
	Left *ValueOrVariable `parser:"@@"`
	//Operator Operator         `parser:"@Operator"`
	//Right    *ValueOrVariable `parser:"@@"`
	Right []*OpTerm `parser:"@@*"`
}

type OpTerm struct {
	Operator Operator         `parser:"@Operator"`
	Right    *ValueOrVariable `parser:"@@"`
}

// todo is this a "Terminal"?
type ValueOrVariable struct {
	VariableIdentifier *string `parser:"@Ident"`
	Value              *Value  `parser:"| @@"`
}

type Value struct {
	//Pos         lexer.Position
	StringValue *string `parser:"@String"`
	//Float       *float64 `parser: "| @Float"`
	// Bool TODO
	NumberValue *float64 `parser:"| @Number"`
}

func (v *Value) String() string {
	if v.StringValue != nil {
		return *v.StringValue
	} else {
		return fmt.Sprintf("%f", *v.NumberValue)
	}
}

// ---
type (
	ExprString struct {
		Value string `@String`
	}

	ExprNumber struct {
		Value float64 `@Int | @Float`
	}

	ExprIdent struct {
		Name string `@Ident`
	}

	ExprParens struct {
		Inner ExprPrecAll `"(" @@ ")"`
	}

	ExprUnary struct {
		Op   string      `@("-" | "!")`
		Expr ExprOperand `@@`
	}

	ExprAddSub struct {
		Head ExprPrec2       `@@`
		Tail []ExprAddSubExt `@@+`
	}

	ExprAddSubExt struct {
		Op   string    `@("+" | "-")`
		Expr ExprPrec2 `@@`
	}

	ExprMulDiv struct {
		Head ExprPrec3       `@@`
		Tail []ExprMulDivExt `@@+`
	}

	ExprMulDivExt struct {
		Op   string    `@("*" | "/")`
		Expr ExprPrec3 `@@`
	}

	ExprRem struct {
		Head ExprOperand  `@@`
		Tail []ExprRemExt `@@+`
	}

	ExprRemExt struct {
		Op   string      `@"%"`
		Expr ExprOperand `@@`
	}

	ExprPrecAll interface{ exprPrecAll() }
	ExprPrec2   interface{ exprPrec2() }
	ExprPrec3   interface{ exprPrec3() }
	ExprOperand interface{ exprOperand() }
)

// essai
func (s ExprString) ToString() string {
	return s.Value
}

// These expression types can be matches as individual operands
func (ExprIdent) exprOperand()  {}
func (ExprNumber) exprOperand() {}
func (ExprString) exprOperand() {}
func (ExprParens) exprOperand() {}
func (ExprUnary) exprOperand()  {}

// These expression types can be matched at precedence level 3
func (ExprIdent) exprPrec3()  {}
func (ExprNumber) exprPrec3() {}
func (ExprString) exprPrec3() {}
func (ExprParens) exprPrec3() {}
func (ExprUnary) exprPrec3()  {}
func (ExprRem) exprPrec3()    {}

// These expression types can be matched at precedence level 2
func (ExprIdent) exprPrec2()  {}
func (ExprNumber) exprPrec2() {}
func (ExprString) exprPrec2() {}
func (ExprParens) exprPrec2() {}
func (ExprUnary) exprPrec2()  {}
func (ExprRem) exprPrec2()    {}
func (ExprMulDiv) exprPrec2() {}

// These expression types can be matched at the minimum precedence level
func (ExprIdent) exprPrecAll()  {}
func (ExprNumber) exprPrecAll() {}
func (ExprString) exprPrecAll() {}
func (ExprParens) exprPrecAll() {}
func (ExprUnary) exprPrecAll()  {}
func (ExprRem) exprPrecAll()    {}
func (ExprMulDiv) exprPrecAll() {}
func (ExprAddSub) exprPrecAll() {}

type Expression struct {
	X ExprPrecAll `@@`
}
