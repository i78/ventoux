package parser

import (
	"dreese.de/ventoux/internal/grammar"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	lex = lexer.MustSimple([]lexer.SimpleRule{
		{`String`, `"(?:\\.|[^"])*"`},
		// {"Int", `\d+`},
		{"Number", `[+-]?([0-9]*[.])?[0-9]+`},
		{"Identifier", `[a-zA-Z0-9]+`},
		{"comment", `//.*|/\*.*?\*/`},
		{"Equals", `=`},
		{"whitespace", ` `},
		{"eol", `[\n\r]+`},
	})
	parser = participle.MustBuild[Program](
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.UseLookahead(2))
)

func GetParser() *participle.Parser[Program] {
	return parser
}

type Program struct {
	Pos    lexer.Position
	TopDec []*grammar.TopDec `@@*`
}
