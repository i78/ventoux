package machine

import (
	"dreese.de/ventoux/internal/grammar"
	"fmt"
)

type Machine struct {
	Variables map[string]*grammar.Expression
}

func (machine *Machine) EvalTop(d *grammar.TopDec) *grammar.Expression {
	if d.Assign != nil {
		return machine.EvalAssign(d.Assign)
	} else if d.Expression != nil {
		return machine.EvalExpr(d.Expression)
	} else if d.ValueOrVariable != nil {
		//return machine.EvalValueOrVariable(d.ValueOrVariable)
		return nil
	}
	return nil
}

func (machine *Machine) EvalValueOrVariable(e *grammar.ValueOrVariable) *grammar.Value {
	fmt.Println("!WARN!!")
	/*if e.Value != nil {
		return e.Value
	} else {
		if e.VariableIdentifier != nil {
			return machine.Variables[*e.VariableIdentifier]
		}
	}*/
	return nil
	// panic("Unable to evaluate expression")
}

func (machine *Machine) EvalExpr(e *grammar.Expression) *grammar.Expression {
	// todo this does not feel right here.
	if _, ok := e.X.(grammar.ExprString); ok {
		// fmt.Println(it.ToString())
		return e
	} else if _, ok := e.X.(grammar.ExprNumber); ok {
		return e
	} else if ident, ok := e.X.(grammar.ExprIdent); ok {
		// todo Lazy evaluation? Not bad. :)
		return machine.Variables[ident.Name]
	}

	// left := machine.EvalValueOrVariable(e.Left).NumberValue
	// right := machine.EvalValueOrVariable(e.Left).NumberValue // todo
	//right := machine.EvalValueOrVariable(e.Right).NumberValue
	/*
		var operationResult float64
		//switch e.Operator {
		switch "+" {
		case "+":
			operationResult = *left + *right
		case "-":
			operationResult = *left - *right
		case "*":
			operationResult = *left * *right
		case "/":
			operationResult = *left / *right
		case "^":
			operationResult = math.Pow(*left, *right)
		case "%":
			operationResult = float64(int(*left) % int(*right))
		case "<<":
			operationResult = float64(int(*left) << int(*right))
		case ">>":
			operationResult = float64(int(*left) >> int(*right))
		}
	*/
	return nil
}

func (machine *Machine) EvalAssign(a *grammar.Assign) *grammar.Expression {
	/*var toAssign *grammar.Value
	if a.ValueOrVariable != nil {
		toAssign = machine.EvalValueOrVariable(a.ValueOrVariable)
	} else if a.Expression != nil {
		toAssign = machine.EvalExpr(a.Expression)
	}
	*/
	machine.Variables[a.Left] = a.Expression
	return nil
}
