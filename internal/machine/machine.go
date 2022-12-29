package machine

import (
	"dreese.de/ventoux/internal/grammar"
	"fmt"
)

type Machine struct {
	Variables map[string]*grammar.Value
}

func (machine *Machine) EvalTop(d *grammar.TopDec) interface{} {
	if d.Assign != nil {
		return machine.EvalAssign(d.Assign)
	} else if d.Expression != nil {
		return machine.EvalExpr(d.Expression)
	} else if d.ValueOrVariable != nil {
		return machine.EvalValueOrVariable(d.ValueOrVariable)
	}
	return nil
}

func (machine *Machine) EvalValueOrVariable(e *grammar.ValueOrVariable) *grammar.Value {
	if e.Value != nil {
		return e.Value
	} else {
		if e.VariableIdentifier != nil {
			return machine.Variables[*e.VariableIdentifier]
		}
	}
	panic("Unable to evaluate expression")
}

func (machine *Machine) EvalExpr(e *grammar.Expression) *grammar.Expression {
	// fmt.Println(">>>", e)
	if it, ok := e.X.(grammar.ExprString); ok {
		fmt.Println(it.ToString())
		return e
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

func (machine *Machine) EvalAssign(a *grammar.Assign) interface{} {
	/*var toAssign *grammar.Value
	if a.ValueOrVariable != nil {
		toAssign = machine.EvalValueOrVariable(a.ValueOrVariable)
	} else if a.Expression != nil {
		toAssign = machine.EvalExpr(a.Expression)
	}

	machine.Variables[a.Left] = toAssign*/
	return nil
}
