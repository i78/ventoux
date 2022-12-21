package grammar

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
)

type TopDec struct {
	Pos     lexer.Position
	Assign  *Assign  `parser:"@@"`
	Literal *Literal `parser:"| @@"`
}

type Literal struct {
	//Value Value `@@`
	Expression *Expression `parser:"@@"`
}

type Assign struct {
	Pos  lexer.Position
	Left string `parser:"@Identifier '='"`
	//Value Value  `parser:"@@"`
	Expression Expression `parser:"@@"`
}

type Expression struct {
	VariableIdentifier *string `parser:"@Identifier"`
	Value              *Value  `parser:"| @@"`
}

type Value struct {
	Pos         lexer.Position
	StringValue *string `parser:"@String"`
	//Float       *float64 `parser: "| @Float"`
	NumberValue *float64 `parser:"| @Number""`
}

func (v *Value) String() string {
	if v.StringValue != nil {
		return *v.StringValue
	} else {
		return fmt.Sprintf("%f", *v.NumberValue)
	}
}
