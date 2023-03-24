package grammar

func (efn ExprFunction) Evaluate(variables *Variables) *Expression {
	(*variables)[efn.FunctionName] = &Expression{X: efn}
	return efn.Expression
}

func (efe ExprFnCallExt) Evaluate(_ *Variables) Expression {
	return efe.Expr
}

func (efc ExprFnCall) Evaluate(variables *Variables) *Expression {
	if fn, exists := (*variables)[efc.FunctionName]; exists {
		if partialInvocation, fnIsPartialInvoked := fn.X.(ExprFnCall); fnIsPartialInvoked {
			args := append(partialInvocation.Tail, efc.Tail...)

			unwrappedFnCall := ExprFnCall{
				FunctionName: partialInvocation.FunctionName,
				Tail:         args,
			}
			return unwrappedFnCall.Evaluate(variables)
		}

		localVariables := efc.buildLocalFunctionEnvironment(variables, fn)
		return efc.executeWithResultEvaluation(fn, localVariables)
	} else {
		panic("function does not exist")
	}
}

func (efc ExprFnCall) executeWithResultEvaluation(fn *Expression, localVariables Variables) *Expression {
	res := fn.Evaluate(&localVariables)
	return res.Evaluate(&localVariables) // curry here?
}

func (efc ExprFnCall) buildLocalFunctionEnvironment(variables *Variables, fn *Expression) Variables {
	if len(efc.Tail) != len(fn.X.(ExprFunction).ParameterNames) {
		panic("argument length of call and function do not match")
	}
	localVariables := Variables{}
	for idx, k := range fn.X.(ExprFunction).ParameterNames {
		localVariables[k] = efc.Tail[idx].Expr.Evaluate(variables)
	}
	return localVariables
}
