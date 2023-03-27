package grammar

import "math"

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

	operationResult := false

	for _, it := range ecm.Tail {
		var right float64
		right = it.Expr.(ExprEvaluatable).Evaluate(variables).X.(ExprNumber).Value

		switch it.Op {
		case ">":
			operationResult = left > right // todo chaining??
		case "<":
			operationResult = left < right
		case ">=":
			operationResult = left >= right
		}
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
