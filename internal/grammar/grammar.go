package grammar

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
)

type TopDec struct {
	Pos             lexer.Position
	Assign          *Assign          `parser:"@@"`
	Expression      *Expr            `parser:"| @@"`
	ValueOrVariable *ValueOrVariable `parser:"| @@"`
}

type Operator string

type Assign struct {
	Pos             lexer.Position
	Left            string           `parser:"@Identifier '='"`
	Expression      *Expr            `parser:"[@@"`
	ValueOrVariable *ValueOrVariable `parser:"| @@]"`
	// Value      Value           `parser:"@@"`
}

type Expr struct {
	Left     *ValueOrVariable `parser:"@@"`
	Operator Operator         `parser:"@Operator"`
	Right    *ValueOrVariable `parser:"@@"`
}

type ValueOrVariable struct {
	VariableIdentifier *string `parser:"@Identifier"`
	Value              *Value  `parser:"| @@"`
}

type Value struct {
	//Pos         lexer.Position
	StringValue *string `parser:"@String"`
	//Float       *float64 `parser: "| @Float"`
	NumberValue *float64 `parser:"| @Number"`
}

func (v *Value) String() string {
	if v.StringValue != nil {
		return *v.StringValue
	} else {
		return fmt.Sprintf("%f", *v.NumberValue)
	}
}
