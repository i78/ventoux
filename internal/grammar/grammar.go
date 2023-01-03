package grammar

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
)

type TopDec struct {
	Pos lexer.Position
	// Function        *Function        `parser:"@@"`
	Assign          *Assign          `parser:"@@"`
	Expression      *Expression      `parser:"| @@"`
	ValueOrVariable *ValueOrVariable // `parser:"| @@"` // todo weg
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

/*
type Function struct {
	Pos            lexer.Position
	FunctionName   *string     `parser:"@Ident"`
	ParameterNames *string     `parser:"@Ident '='"`
	Expression     *Expression `parser:"@@';'"` // todo this must be lazy
}*/

func (v *Value) String() string {
	if v.StringValue != nil {
		return *v.StringValue
	} else {
		return fmt.Sprintf("%f", *v.NumberValue)
	}
}
