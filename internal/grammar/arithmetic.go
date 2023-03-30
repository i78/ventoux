package grammar

import "math"

type (
	ExprAddSub struct {
		Head ExprPrec2       `parser:"@@"`
		Tail []ExprAddSubExt `parser:"@@+"`
	}

	ExprAddSubExt struct {
		Op   string    `parser:"@('+' | '-')"`
		Expr ExprPrec2 `parser:"@@"`
	}

	ExprMulDiv struct {
		Head ExprPrec3       `parser:"@@"`
		Tail []ExprMulDivExt `parser:"@@+"`
	}

	ExprMulDivExt struct {
		Op   string    `parser:"@('*' | '/')"`
		Expr ExprPrec3 `parser:"@@"`
	}

	ExprPow struct {
		Head ExprPrec3    `parser:"@@"`
		Tail []ExprPowExt `parser:"@@+"`
	}

	ExprPowExt struct {
		Op   string    `parser:"@('^' )"`
		Expr ExprPrec3 `parser:"@@"`
	}

	ExprComparison struct {
		Head ExprPrec3           `parser:"@@"`
		Tail []ExprComparisonExt `parser:"@@+"`
	}

	ExprComparisonExt struct {
		Op   string    `parser:"@('<' | '>' | '>=' | '<=' | '=' )"`
		Expr ExprPrec3 `parser:"@@"`
	}

	ExprBitshift struct {
		Head ExprPrec3         `parser:"@@"`
		Tail []ExprBitshiftExt `parser:"@@+"`
	}

	ExprBitshiftExt struct {
		Op   string    `parser:"@('<<' | '>>' )"`
		Expr ExprPrec3 `parser:"@@"`
	}
)

func (eas ExprAddSub) Evaluate(variables *Variables) *Expression {
	var left float64

	left = eas.Head.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

	operationResult := left

	for _, it := range eas.Tail {
		var right float64
		right = it.Expr.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

		switch it.Op {
		case "+":
			operationResult += right
		case "-":
			operationResult -= right
		case "*":
			operationResult *= right
		case "/":
			operationResult /= right
		case "%":
			operationResult = float64(int(left) % int(right))
		}
	}

	return &Expression{X: ExprNumber{Value: operationResult}}
}

func (eas ExprMulDiv) Evaluate(variables *Variables) *Expression {
	var left float64

	left = eas.Head.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

	operationResult := left

	for _, it := range eas.Tail {
		var right float64
		right = it.Expr.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

		switch it.Op {
		case "*":
			operationResult *= right
		case "/":
			operationResult /= right
		}
	}

	return &Expression{X: ExprNumber{Value: operationResult}}
}

func (ecm ExprComparison) Evaluate(variables *Variables) *Expression {
	var left float64

	left = ecm.Head.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

	operationResult := true

	for _, it := range ecm.Tail {
		var right float64
		right = it.Expr.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

		switch it.Op {
		case ">":
			operationResult = operationResult && left > right
		case "<":
			operationResult = operationResult && left < right
		case ">=":
			operationResult = operationResult && left >= right
		case "=":
			operationResult = operationResult && left == right
		}

		left = right
	}

	return &Expression{X: ExprBoolean{Value: operationResult}}
}

func (eas ExprBitshift) Evaluate(variables *Variables) *Expression {
	operationResult := eas.Head.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

	for _, it := range eas.Tail {
		right := it.Expr.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

		switch it.Op {
		case "<<":
			operationResult = float64(int(operationResult) << int(right))
		case ">>":
			operationResult = float64(int(operationResult) >> int(right))
		}
	}

	return &Expression{X: ExprNumber{Value: operationResult}}
}

func (eas ExprPow) Evaluate(variables *Variables) *Expression {
	operationResult := eas.Head.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

	for _, it := range eas.Tail {
		right := it.Expr.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

		switch it.Op {
		case "^":
			operationResult = math.Pow(operationResult, right)
		}
	}
	return &Expression{X: ExprNumber{Value: operationResult}}
}
