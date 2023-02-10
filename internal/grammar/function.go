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
		localVariables := efc.buildLocalFunctionEnvironment(variables, fn)
		return efc.executeWithResultEvaluation(fn, localVariables)
	} else {
		panic("function does not exist")
	}
}

func (efc ExprFnCall) executeWithResultEvaluation(fn *Expression, localVariables Variables) *Expression {
	res := fn.Evaluate(&localVariables)
	return res.Evaluate(&localVariables)
}

func (efc ExprFnCall) buildLocalFunctionEnvironment(variables *Variables, fn *Expression) Variables {
	localVariables := Variables{}
	for idx, k := range fn.X.(ExprFunction).ParameterNames {
		localVariables[k] = efc.Tail[idx].Expr.Evaluate(variables)
	}
	return localVariables
}
