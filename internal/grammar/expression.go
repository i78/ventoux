package grammar

import (
	"fmt"
	"math"
)

// todo this does not belong here.
type Variables = map[string]*Expression

type (
	ExprString struct {
		Value string `@String`
	}

	ExprNumber struct {
		Value float64 `@Int | @Float`
	}

	ExprIdent struct {
		Name string `@Ident`
	}

	ExprParens struct {
		Inner ExprPrecAll `"(" @@ ")"`
	}

	ExprUnary struct {
		Op   string      `@("-" | "!")`
		Expr ExprOperand `@@`
	}

	ExprAddSub struct {
		Head ExprPrec2       `@@`
		Tail []ExprAddSubExt `@@+`
	}

	ExprAddSubExt struct {
		Op   string    `@("+" | "-")`
		Expr ExprPrec2 `@@`
	}

	ExprMulDiv struct {
		Head ExprPrec3       `@@`
		Tail []ExprMulDivExt `@@+`
	}

	ExprMulDivExt struct {
		Op   string    `@("*" | "/")`
		Expr ExprPrec3 `@@`
	}

	ExprPow struct {
		Head ExprPrec3    `@@`
		Tail []ExprPowExt `@@+`
	}

	ExprPowExt struct {
		Op   string    `@("^" )`
		Expr ExprPrec3 `@@`
	}

	ExprBitshift struct {
		Head ExprPrec3         `@@`
		Tail []ExprBitshiftExt `@@+`
	}

	ExprBitshiftExt struct {
		Op   string    `@("<<" | ">>" )`
		Expr ExprPrec3 `@@`
	}

	ExprRem struct {
		Head ExprOperand  `@@`
		Tail []ExprRemExt `@@+`
	}

	ExprRemExt struct {
		Op   string      `@"%"`
		Expr ExprOperand `@@`
	}

	ExprFunction struct {
		FunctionName   string      `parser:"@Ident"`
		ParameterNames string      `parser:"@Ident '='"`
		Expression     *Expression `parser:"@@';'"` // todo this must be lazy
	}

	ExprFnCall struct {
		//FunctionName *ExprIdent      `@`
		FunctionName string          `@Ident`
		Tail         []ExprFnCallExt `@@`
	}

	ExprFnCallExt struct {
		Expr ExprPrec3 `"("@@*")"`
	}

	ExprPrecAll interface{ exprPrecAll() }
	ExprPrec2   interface{ exprPrec2() }
	ExprPrec3   interface{ exprPrec3() }
	ExprOperand interface{ exprOperand() }
	//ExprFunctionCall interface{ exprFunctionCall() }

	ExprEvaluatable interface{ Evaluate(*Variables) *Expression }
	ExprTerminal    interface{ Terminal() *Expression }
)

// These expression types can be matches as individual operands
func (ExprIdent) exprOperand()    {}
func (ExprNumber) exprOperand()   {}
func (ExprString) exprOperand()   {}
func (ExprParens) exprOperand()   {}
func (ExprUnary) exprOperand()    {}
func (ExprFnCall) exprOperand()   {}
func (ExprFunction) exprOperand() {}

// These expression types can be matched at precedence level 3
func (ExprIdent) exprPrec3()  {}
func (ExprNumber) exprPrec3() {}
func (ExprString) exprPrec3() {}
func (ExprParens) exprPrec3() {}
func (ExprUnary) exprPrec3()  {}
func (ExprRem) exprPrec3()    {}
func (ExprFnCall) exprPrec3() {}

// These expression types can be matched at precedence level 2
func (ExprIdent) exprPrec2()    {}
func (ExprNumber) exprPrec2()   {}
func (ExprString) exprPrec2()   {}
func (ExprParens) exprPrec2()   {}
func (ExprUnary) exprPrec2()    {}
func (ExprRem) exprPrec2()      {}
func (ExprMulDiv) exprPrec2()   {}
func (ExprPow) exprPrec2()      {}
func (ExprBitshift) exprPrec2() {}
func (ExprFnCall) exprPrec2()   {}

// These expression types can be matched at the minimum precedence level
func (ExprIdent) exprPrecAll()    {}
func (ExprNumber) exprPrecAll()   {}
func (ExprString) exprPrecAll()   {}
func (ExprParens) exprPrecAll()   {}
func (ExprUnary) exprPrecAll()    {}
func (ExprRem) exprPrecAll()      {}
func (ExprMulDiv) exprPrecAll()   {}
func (ExprPow) exprPrecAll()      {}
func (ExprBitshift) exprPrecAll() {}
func (ExprAddSub) exprPrecAll()   {}
func (ExprFnCall) exprPrecAll()   {}
func (ExprFunction) exprPrecAll() {}

type Expression struct {
	X ExprPrecAll `@@`
}

func (e *Expression) ToString() string {
	if it, ok := e.X.(ExprString); ok {
		return it.Value
	} else if it, ok := e.X.(ExprNumber); ok {
		return fmt.Sprintf("%f", it.Value)
	}
	return ""
}

func (eid ExprIdent) Evaluate(variables *Variables) *Expression {
	if it, exists := (*variables)[eid.Name]; exists {
		return it.Evaluate(variables)
	}
	// Else, return lazy lets talk about it.
	return &Expression{X: eid}
}

func (en ExprNumber) Terminal() *Expression {
	return &Expression{X: en}
}

func (en ExprNumber) Evaluate(*Variables) *Expression {
	return &Expression{X: en}
}

func (es ExprString) Evaluate(*Variables) *Expression {
	return &Expression{X: es}
}

// todo type
func (e Expression) Evaluate(variables *Variables) *Expression {
	if it, isTerminal := e.X.(ExprTerminal); isTerminal {
		return it.Terminal()
	} else if it, isEvaluable := e.X.(ExprEvaluatable); isEvaluable {
		return it.Evaluate(variables)
	}
	return nil
}

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

func (ep ExprParens) Evaluate(variables *Variables) *Expression {
	return ep.Inner.(ExprEvaluatable).Evaluate(variables)
}

func (efn ExprFunction) Evaluate(variables *Variables) *Expression {
	(*variables)[efn.FunctionName] = &Expression{X: efn}
	return nil
}

func (efc ExprFnCall) Evaluate(variables *Variables) *Expression {
	if _, exists := (*variables)[efc.FunctionName]; exists {
		fmt.Println("ALIFE")
	}
	return nil
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
