package grammar

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
)

type TopDec struct {
	Pos        lexer.Position
	Assign     *Assign     `parser:"@@"`
	Expression *Expression `parser:"| @@"`
}

type Operator string

type Assign struct {
	Pos        lexer.Position
	Left       string      `parser:"@Ident '='"`
	Expression *Expression `parser:"@@"`
}

type Expr struct {
	Left  *ValueOrVariable `parser:"@@"`
	Right []*OpTerm        `parser:"@@*"`
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
	Pos         lexer.Position
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
