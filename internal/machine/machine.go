package machine

import (
	"bytes"
	"dreese.de/ventoux/internal/grammar"
	"dreese.de/ventoux/internal/parser"
	"encoding/gob"
	"log"
)

type ConsoleOutputFn = func(string)

type Machine struct {
	grammar.Variables
	ConsoleOutputFn
}

func NewMachine(consoleOutputFn ConsoleOutputFn) *Machine {
	return &Machine{Variables: map[string]*grammar.Expression{}, ConsoleOutputFn: consoleOutputFn}
}

func (machine *Machine) EvalProgram(program *parser.Program) *grammar.Expression {
	for _, st := range program.TopDec {
		res := machine.EvalTop(st)
		if res != nil {
			machine.ConsoleOutputFn(res.ToString())
		}
	}
	return nil
}

func (machine *Machine) EvalTop(d *grammar.TopDec) *grammar.Expression {
	if d.Assign != nil {
		return machine.EvalAssign(d.Assign)
	} else if d.Expression != nil {
		return machine.EvalExpr(d.Expression)
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

func (machine *Machine) ExportMachineState() []byte {
	var buffer bytes.Buffer
	withGobConfiguration()

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(machine)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	return buffer.Bytes()
}

func withGobConfiguration() {
	gob.Register(Machine{})
	gob.Register(grammar.ExprIdent{})
	gob.Register(grammar.ExprNumber{})
	gob.Register(grammar.ExprString{})
	gob.Register(grammar.ExprParens{})
	gob.Register(grammar.ExprUnary{})
	gob.Register(grammar.ExprRem{})
	gob.Register(grammar.ExprMulDiv{})
	gob.Register(grammar.ExprPow{})
	gob.Register(grammar.ExprBitshift{})
	gob.Register(grammar.ExprComparison{})
	gob.Register(grammar.ExprAddSub{})
}
