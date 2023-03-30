package grammar

type (
	ExprString struct {
		Value string `parser:"@String"`
	}

	ExprNumber struct {
		Value float64 `parser:"@Int | @Float"`
	}

	ExprBoolean struct {
		Value bool `parser:"@Boolean"`
	}
)

func (en ExprNumber) Terminal() *Expression {
	return &Expression{X: en}
}

func (en ExprNumber) Evaluate(*Variables) *Expression {
	return &Expression{X: en}
}

func (en ExprBoolean) Terminal() *Expression {
	return &Expression{X: en}
}

func (es ExprString) Evaluate(*Variables) *Expression {
	return &Expression{X: es}
}
