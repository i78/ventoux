package parser

import (
	"dreese.de/ventoux/internal/grammar"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	lex = lexer.MustSimple([]lexer.SimpleRule{
		// Rem: Participle respects the ordering of this list.
		{"comment", `//.*|/\*.*?\*/`},
		{`String`, `"(?:\\.|[^"])*"`},
		//{"Number", `XKXK[+-]?([0-9]*[.])?[0-9]+`},
		{"Float", `[+-]?([0-9]*[.])?[0-9]+`},
		{"Int", `[+-]?[0-9]+`},
		{"Operator", `[-,()*/+%{};&!=:^]|>>|<<`},
		{"Ident", `[a-zA-Z]+[0-9]?`},
		{"Equals", `=`},
		{"whitespace", ` `},
		{"eol", `[\n\r]+`},
	})
	parser = participle.MustBuild[Program](
		participle.Lexer(lex),
		participle.Unquote("String"),

		participle.Union[grammar.ExprOperand](grammar.ExprUnary{}, grammar.ExprIdent{}, grammar.ExprNumber{}, grammar.ExprString{}, grammar.ExprParens{}),
		// Register the grammar.ExprPrec3 union so we can parse expressions at precedence level 3
		participle.Union[grammar.ExprPrec3](grammar.ExprRem{}, grammar.ExprUnary{}, grammar.ExprIdent{}, grammar.ExprNumber{}, grammar.ExprString{}, grammar.ExprParens{}),
		// Register the grammar.ExprPrec2 union so we can parse expressions at precedence level 2
		participle.Union[grammar.ExprPrec2](grammar.ExprMulDiv{}, grammar.ExprPow{}, grammar.ExprBitshift{}, grammar.ExprRem{}, grammar.ExprUnary{}, grammar.ExprIdent{}, grammar.ExprNumber{}, grammar.ExprString{}, grammar.ExprParens{}),
		// Register the grammar.ExprPrecAll union so we can parse expressions at the minimum precedence level
		participle.Union[grammar.ExprPrecAll](grammar.ExprAddSub{}, grammar.ExprMulDiv{}, grammar.ExprPow{}, grammar.ExprBitshift{}, grammar.ExprRem{}, grammar.ExprUnary{}, grammar.ExprIdent{}, grammar.ExprNumber{}, grammar.ExprString{}, grammar.ExprParens{}),
		participle.Elide("comment"),
		participle.UseLookahead(99999))
)

func GetParser() *participle.Parser[Program] {
	return parser
}

type Program struct {
	Pos    lexer.Position
	TopDec []*grammar.TopDec `parser:"@@*"`
}
