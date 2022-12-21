package machine

import (
	"dreese.de/ventoux/internal/grammar"
	"fmt"
)

type Machine struct {
	Variables map[string]*grammar.Value
}

func (machine *Machine) EvalTop(d *grammar.TopDec) {
	if d.Assign != nil {
		machine.EvalAssign(d.Assign)
	} else if d.Literal != nil {
		machine.EvalLit(d.Literal)
	}
}

func (machine *Machine) EvalLit(l *grammar.Literal) {
	fmt.Println(machine.EvalExpression(l.Expression).String())
}

func (machine *Machine) EvalExpression(e *grammar.Expression) *grammar.Value {
	if e.Value != nil {
		return e.Value
	} else {
		if e.VariableIdentifier != nil {
			return machine.Variables[*e.VariableIdentifier]
		}
	}
	panic("Unable to evaluate expression")
}

func (machine *Machine) EvalAssign(a *grammar.Assign) {
	//fmt.Println("ASSIGN", a.Left, " = ", a.Value)
	machine.Variables[a.Left] = machine.EvalExpression(&a.Expression)
}
