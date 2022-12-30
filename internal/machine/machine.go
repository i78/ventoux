package machine

import (
	"dreese.de/ventoux/internal/grammar"
)

type Machine struct {
	grammar.Variables
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

func (machine *Machine) EvalExpr(e *grammar.Expression) *grammar.Expression {
	return e.Evaluate(&machine.Variables)
}

func (machine *Machine) EvalAssign(a *grammar.Assign) *grammar.Expression {
	machine.Variables[a.Left] = a.Expression
	return nil
}
