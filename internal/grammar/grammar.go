package grammar

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type TopDec struct {
	Pos        lexer.Position
	Assign     *Assign     `parser:"@@"`
	Expression *Expression `parser:"| @@"`
}

type Assign struct {
	Pos        lexer.Position
	Left       string      `parser:"@Ident '='"`
	Expression *Expression `parser:"@@"`
}
