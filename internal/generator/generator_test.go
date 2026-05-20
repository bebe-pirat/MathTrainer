package generator

import (
	"MathTrainer/internal"
	"MathTrainer/internal/model"
	"strings"
	"testing"
)

func TestGenerateEquation_2operands_1operation(t *testing.T) {
	tests := []struct {
		name string
		data model.EquationTypeWithOperands
	}{
		{
			name: "adding",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "+",
				NumOperands: 2,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     10,
						MaxValue:     50,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     10,
						MaxValue:     50,
					},
				},
			},
		},
		{
			name: "substraction",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "-",
				NumOperands: 2,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     10,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     10,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "multiplication",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "*",
				NumOperands: 2,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "division",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "/",
				NumOperands: 2,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := GenerateEquation(test.data)
			if err != nil {
				t.Errorf("failed to generate equation with test data, test data: %v", test.data)
			}

			if res.CorrectAnswer > test.data.MaxResult && res.CorrectAnswer < 0 {
				t.Errorf("expected to have result less than %d, but it has: %d", test.data.MaxResult, res.CorrectAnswer)

			}

			switch test.data.Operations {
			case "+", "-":
				if !strings.Contains(res.Text, test.data.Operations) {
					t.Errorf("expected to have following operation in equation text: %s, but it doesn't have any, equation text: %s", test.data.Operations, res.Text)
				}
			case "*":
				if !strings.Contains(res.Text, internal.MultiplicationSymbol) {
					t.Errorf("expected to have following operation in equation text: %s, but it doesn't have any, equation text: %s", test.data.Operations, res.Text)
				}
			case "/":
				if !strings.Contains(res.Text, internal.DivisionSimbol) {
					t.Errorf("expected to have following operation in equation text: %s, but it doesn't have any, equation text: %s", test.data.Operations, res.Text)
				}
			}

			if !strings.Contains(res.Text, "= ?") {
				t.Errorf("missing '= ?' in %q", res.Text)
			}
		})
	}
}

func TestGenerateEquation_3operands_2SameOperations(t *testing.T) {
	tests := []struct {
		name string
		data model.EquationTypeWithOperands
	}{
		{
			name: "adding",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "+",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   150,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "substraction",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "-",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "multiplication",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "*",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   1000,
				Operands: []model.Operand{
					{
						Id:           1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "division",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "/",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{
						Id:           1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := GenerateEquation(test.data)
			if err != nil {
				t.Errorf("failed to generate equation with test data, test data: %v", test.data)
			}

			if res.CorrectAnswer > test.data.MaxResult && res.CorrectAnswer < 0 {
				t.Errorf("expected to have result less than %d, but it has: %d", test.data.MaxResult, res.CorrectAnswer)

			}

			switch test.data.Operations {
			case "+", "-":
				if !strings.Contains(res.Text, test.data.Operations) {
					t.Errorf("expected to have following operation in equation text: %s, but it doesn't have any, equation text: %s", test.data.Operations, res.Text)
				}
			case "*":
				if !strings.Contains(res.Text, internal.MultiplicationSymbol) {
					t.Errorf("expected to have following operation in equation text: %s, but it doesn't have any, equation text: %s", test.data.Operations, res.Text)
				}
			case "/":
				if !strings.Contains(res.Text, internal.DivisionSimbol) {
					t.Errorf("expected to have following operation in equation text: %s, but it doesn't have any, equation text: %s", test.data.Operations, res.Text)
				}
			}

			if !strings.Contains(res.Text, "= ?") {
				t.Errorf("missing '= ?' in %q", res.Text)
			}
		})
	}
}

func TestGenerateEquation_3operands_2DiffOperations(t *testing.T) {
	tests := []struct {
		name string
		data model.EquationTypeWithOperands
	}{
		{
			name: "adding&substraction",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "+-",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   150,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "substraction&division",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "/-",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{Id: 1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "multiplication&substraction",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "*-",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   1000,
				Operands: []model.Operand{
					{
						Id:           1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
		{
			name: "division&adding",
			data: model.EquationTypeWithOperands{
				Id:          1,
				Operations:  "/+",
				NumOperands: 3,
				NoRemainder: true,
				MaxResult:   100,
				Operands: []model.Operand{
					{
						Id:           1,
						OperandOrder: 1,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           2,
						OperandOrder: 2,
						MinValue:     1,
						MaxValue:     100,
					},
					{
						Id:           3,
						OperandOrder: 3,
						MinValue:     1,
						MaxValue:     100,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := GenerateEquation(test.data)
			if err != nil {
				t.Errorf("failed to generate equation with test data, test data: %v", test.data)
			}

			if res.CorrectAnswer > test.data.MaxResult && res.CorrectAnswer < 0 {
				t.Errorf("expected to have result less than %d, but it has: %d", test.data.MaxResult, res.CorrectAnswer)

			}

			operationsStr := strings.ReplaceAll(test.data.Operations, "*", internal.MultiplicationSymbol)
			operations := []rune(strings.ReplaceAll(operationsStr, "/", internal.DivisionSimbol))

			countOperations := 0
			for _, op := range operations {
				if strings.Contains(res.Text, string(op)) {
					countOperations++
				}
			}

			if countOperations < 0 {
				t.Errorf("expected to have following operations in equation text: %s, but it doesn't have any, equation text: %s", string(operations), res.Text)
			}

			if !strings.Contains(res.Text, "= ?") {
				t.Errorf("missing '= ?' in %q", res.Text)
			}
		})
	}
}

func TestGenerateEquation_RangeChecks(t *testing.T) {
	eqType := model.EquationTypeWithOperands{
		NumOperands: 2,
		Operations:  "+-",
		Operands: []model.Operand{
			{OperandOrder: 1, MinValue: 1, MaxValue: 10},
			{OperandOrder: 2, MinValue: 1, MaxValue: 10},
		},
		MaxResult: 20,
		Id:        1,
	}

	for i := 0; i < 100; i++ {
		eq, err := GenerateEquation(eqType)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if eq.CorrectAnswer > eqType.MaxResult && eq.CorrectAnswer > -1 {
			t.Errorf("answer %d exceeds max %d", eq.CorrectAnswer, eqType.MaxResult)
		}
		if !strings.Contains(eq.Text, "= ?") {
			t.Errorf("missing '= ?' in %q", eq.Text)
		}
	}
}
