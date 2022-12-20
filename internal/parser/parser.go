package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	lex = lexer.MustSimple([]lexer.SimpleRule{
		{"Literal", `".*"`},
		{"comment", `//.*|/\*.*?\*/`},
		{"whitespace", `\s+`},
	})
	parser = participle.MustBuild[Program](
		participle.Lexer(lex),
		participle.Unquote("Literal"),
		participle.UseLookahead(2))
)

func GetParser() *participle.Parser[Program] {
	return parser
}

type Program struct {
	Pos    lexer.Position
	TopDec []*TopDec `@@*`
}

type TopDec struct {
	Pos lexer.Position

	Literal string `@Literal`
}
