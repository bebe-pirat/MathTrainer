package generator

import (
	"MathTrainer/internal"
	"MathTrainer/internal/model"
	"math/rand/v2"
	"strconv"
)

func GenerateEquation(equationType model.EquationTypeWithOperands) (model.Equation, error) {
	vars := make([]string, equationType.NumOperands)
	ops := make([]string, equationType.NumOperands-1)
	expr := make([]string, cap(vars)+cap(ops))
	var eqStr string = ""

	var correctAnswer int = 0
	var err error
	for {
		runes := []rune(equationType.Operations)
		for i := 0; i < equationType.NumOperands; i++ {
			operandRange := equationType.Operands[i]
			vars[i] = strconv.Itoa(rand.IntN(operandRange.MaxValue-operandRange.MinValue) + operandRange.MinValue)
			expr[i*2] = vars[i]
			eqStr += vars[i]

			if i < cap(ops) {
				ops[i] = string(runes[rand.IntN(len(runes))])

				switch ops[i] {
				case "/":
					ops[i] = internal.DivisionSimbol
				case "*":
					ops[i] = internal.MultiplicationSymbol
				}

				expr[i*2+1] = ops[i]
				eqStr += ops[i]
			}
		}

		eqStr += "= ?"

		m := NewMather(expr, equationType.MaxResult)
		correctAnswer, err = m.Calculate()
		if err == nil {
			break
		} else {
			eqStr = ""
		}
	}

	return model.Equation{
		Text:           eqStr,
		CorrectAnswer:  correctAnswer,
		EquationTypeId: equationType.Id,
	}, nil
}
