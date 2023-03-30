package grammar

type (
	ExprConditional struct {
		Options []*ExprConditionalOption `parser:"'{'@@*'}'"`
	}

	ExprConditionalOption struct {
		Precondition *Expression `parser:"@@?"`
		Result       *Expression `parser:"Arrow @@"`
	}
)

// func (ExprConditional) exprOperand() {}
func (ExprConditional) exprPrecAll() {}

func (ec ExprConditional) Evaluate(variables *Variables) *Expression {
	for _, option := range ec.Options {
		if option.Precondition == nil {
			return option.Result
		}
		ev := option.Precondition.Evaluate(variables)

		if evBool, isBoolean := ev.X.(ExprBoolean); isBoolean {
			if evBool.Value == true {
				return option.Result
			}
		}
	}
	return &Expression{}
}
