package generator

import (
	"MathTrainer/internal"
	"strings"
	"testing"
)

func TestCalculateExpr_2operands_and_1operation(t *testing.T) {
	tests := []struct {
		name           string
		expr           []string
		maxRes         int
		expectedAnswer int
	}{
		{name: "adding 1", expr: strings.Split("1 + 1", " "), maxRes: 10, expectedAnswer: 2},
		{name: "adding 2", expr: strings.Split("2 + 5", " "), maxRes: 10, expectedAnswer: 7},
		{name: "adding 3", expr: strings.Split("1000 + 190", " "), maxRes: 10000, expectedAnswer: 1190},
		{name: "adding 4", expr: strings.Split("2193 + 4324", " "), maxRes: 10000, expectedAnswer: 6517},

		{name: "substraction 1", expr: strings.Split("1 - 1", " "), maxRes: 10000, expectedAnswer: 0},
		{name: "substraction 2", expr: strings.Split("5 - 2", " "), maxRes: 10000, expectedAnswer: 3},
		{name: "substraction 3", expr: strings.Split("1000 - 190", " "), maxRes: 10000, expectedAnswer: 810},
		{name: "substraction 4", expr: strings.Split("4324 - 2193", " "), maxRes: 10000, expectedAnswer: 2131},

		{name: "multiplication 1", expr: strings.Split("1 "+internal.MultiplicationSymbol+" 1", " "), maxRes: 10000, expectedAnswer: 1},
		{name: "multiplication 2", expr: strings.Split("2 "+internal.MultiplicationSymbol+" 5", " "), maxRes: 10000, expectedAnswer: 10},
		{name: "multiplication 3", expr: strings.Split("1000 "+internal.MultiplicationSymbol+" 190", " "), maxRes: 100000000, expectedAnswer: 190000},
		{name: "multiplication 4", expr: strings.Split("2193 "+internal.MultiplicationSymbol+" 4324", " "), maxRes: 1000000000, expectedAnswer: 9482532},

		{name: "division 1", expr: strings.Split("1 "+internal.DivisionSimbol+" 1", " "), maxRes: 10000, expectedAnswer: 1},
		{name: "division 2", expr: strings.Split("4 "+internal.DivisionSimbol+" 2", " "), maxRes: 10000, expectedAnswer: 2},
		{name: "division 3", expr: strings.Split("1000 "+internal.DivisionSimbol+" 100", " "), maxRes: 10000, expectedAnswer: 10},
		{name: "division 4", expr: strings.Split("381 "+internal.DivisionSimbol+" 3", " "), maxRes: 10000, expectedAnswer: 127},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewMather(test.expr, test.maxRes)
			answer, err := m.Calculate()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if test.expectedAnswer != answer {
				expr := strings.Join(test.expr, " ")
				t.Errorf("equation: %s, expected answer: %d, real answer: %d", expr, test.expectedAnswer, answer)
			}
		})
	}
}

func TestCalculateExpr_3operands_and_2SameOperation(t *testing.T) {
	tests := []struct {
		name           string
		expr           []string
		maxRes         int
		expectedAnswer int
	}{
		{name: "adding 1", expr: strings.Split("1 + 1 + 5", " "), maxRes: 10, expectedAnswer: 7},
		{name: "adding 2", expr: strings.Split("2 + 5 + 2 ", " "), maxRes: 10, expectedAnswer: 9},
		{name: "adding 3", expr: strings.Split("1000 + 190 + 923", " "), maxRes: 10000, expectedAnswer: 2113},
		{name: "adding 4", expr: strings.Split("2193 + 4324 + 823", " "), maxRes: 10000, expectedAnswer: 7340},

		{name: "substraction 1", expr: strings.Split("10 - 1 - 6", " "), maxRes: 10000, expectedAnswer: 3},
		{name: "substraction 2", expr: strings.Split("5 - 2 - 1", " "), maxRes: 10000, expectedAnswer: 2},
		{name: "substraction 3", expr: strings.Split("1000 - 190 - 810", " "), maxRes: 10000, expectedAnswer: 0},
		{name: "substraction 4", expr: strings.Split("4324 - 2193 - 1221", " "), maxRes: 10000, expectedAnswer: 910},

		{name: "multiplication 1", expr: strings.Split("1 "+internal.MultiplicationSymbol+" 1 "+internal.MultiplicationSymbol+" 3", " "), maxRes: 10000, expectedAnswer: 3},
		{name: "multiplication 2", expr: strings.Split("2 "+internal.MultiplicationSymbol+" 5 "+internal.MultiplicationSymbol+" 21", " "), maxRes: 10000, expectedAnswer: 210},
		{name: "multiplication 3", expr: strings.Split("1000 "+internal.MultiplicationSymbol+" 190 "+internal.MultiplicationSymbol+" 12", " "), maxRes: 100000000, expectedAnswer: 2280000},
		{name: "multiplication 4", expr: strings.Split("2193 "+internal.MultiplicationSymbol+" 4324 "+internal.MultiplicationSymbol+" 65", " "), maxRes: 1000000000, expectedAnswer: 616364580},

		{name: "division 1", expr: strings.Split("6 "+internal.DivisionSimbol+" 2 "+internal.DivisionSimbol+" 3", " "), maxRes: 10000, expectedAnswer: 1},
		{name: "division 2", expr: strings.Split("4 "+internal.DivisionSimbol+" 2 "+internal.DivisionSimbol+" 2", " "), maxRes: 10000, expectedAnswer: 1},
		{name: "division 3", expr: strings.Split("1000 "+internal.DivisionSimbol+" 100 "+internal.DivisionSimbol+" 10", " "), maxRes: 10000, expectedAnswer: 1},
		{name: "division 4", expr: strings.Split("381 "+internal.DivisionSimbol+" 3 "+internal.DivisionSimbol+" 127", " "), maxRes: 10000, expectedAnswer: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewMather(test.expr, test.maxRes)
			answer, err := m.Calculate()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if test.expectedAnswer != answer {
				expr := strings.Join(test.expr, " ")
				t.Errorf("equation: %s, expected answer: %d, real answer: %d", expr, test.expectedAnswer, answer)
			}
		})
	}
}
func TestCalculateExpr_MixedOperations_3Operands(t *testing.T) {
	tests := []struct {
		name           string
		expr           []string
		maxRes         int
		expectedAnswer int
	}{
		{name: "add and mul 1", expr: strings.Split("2 + 3 × 4", " "), maxRes: 100, expectedAnswer: 14},
		{name: "add and mul 2", expr: strings.Split("10 + 2 × 7", " "), maxRes: 100, expectedAnswer: 24},
		{name: "sub and div 1", expr: strings.Split("10 - 6 ÷ 2", " "), maxRes: 100, expectedAnswer: 7},
		{name: "sub and div 2", expr: strings.Split("20 - 15 ÷ 3", " "), maxRes: 100, expectedAnswer: 15},
		{name: "mul and sub", expr: strings.Split("5 × 4 - 6", " "), maxRes: 100, expectedAnswer: 14},
		{name: "div and add", expr: strings.Split("8 ÷ 2 + 3", " "), maxRes: 100, expectedAnswer: 7},
		{name: "all ops", expr: strings.Split("12 + 6 ÷ 3 - 2 × 2", " "), maxRes: 100, expectedAnswer: 10}, // 12+2-4=10
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewMather(test.expr, test.maxRes)
			answer, err := m.Calculate()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if answer != test.expectedAnswer {
				exprStr := strings.Join(test.expr, " ")
				t.Errorf("expr: %s, expected %d, got %d", exprStr, test.expectedAnswer, answer)
			}
		})
	}
}

func TestCalculateExpr_4Operands(t *testing.T) {
	tests := []struct {
		name           string
		expr           []string
		maxRes         int
		expectedAnswer int
	}{
		{name: "add,mul,sub", expr: strings.Split("2 + 3 × 4 - 1", " "), maxRes: 100, expectedAnswer: 13},
		{name: "sub,mul,add", expr: strings.Split("10 - 2 × 3 + 4", " "), maxRes: 100, expectedAnswer: 8},
		{name: "div,add,mul", expr: strings.Split("20 ÷ 4 + 3 × 2", " "), maxRes: 100, expectedAnswer: 11},
		{name: "add,div,sub", expr: strings.Split("9 + 6 ÷ 3 - 2", " "), maxRes: 100, expectedAnswer: 9},
		{name: "div,mul,sub", expr: strings.Split("18 ÷ 3 × 2 - 4", " "), maxRes: 100, expectedAnswer: 8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewMather(test.expr, test.maxRes)
			answer, err := m.Calculate()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if answer != test.expectedAnswer {
				exprStr := strings.Join(test.expr, " ")
				t.Errorf("expr: %s, expected %d, got %d", exprStr, test.expectedAnswer, answer)
			}
		})
	}
}

func TestCalculateExpr_7Operands(t *testing.T) {
	tests := []struct {
		name           string
		expr           []string
		maxRes         int
		expectedAnswer int
	}{
		{
			name:           "complex 1",
			expr:           strings.Split("5 + 3 × 2 - 4 ÷ 2 + 7 - 1 × 3", " "),
			maxRes:         100,
			expectedAnswer: 13,
		},
		{
			name:           "complex 2",
			expr:           strings.Split("10 ÷ 2 + 3 × 4 - 6 + 2 × 3 - 1", " "),
			maxRes:         100,
			expectedAnswer: 16,
		},
		{
			name:           "complex 3",
			expr:           strings.Split("100 - 20 × 2 + 30 ÷ 3 - 5 × 2 + 4", " "),
			maxRes:         200,
			expectedAnswer: 64,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewMather(test.expr, test.maxRes)
			answer, err := m.Calculate()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if answer != test.expectedAnswer {
				exprStr := strings.Join(test.expr, " ")
				t.Errorf("expr: %s, expected %d, got %d", exprStr, test.expectedAnswer, answer)
			}
		})
	}
}

func TestCalculateExpr_ExpectedError_UnderZero(t *testing.T) {
	tests := []struct {
		name   string
		expr   []string
		maxRes int
		errMsg string
	}{
		{
			name:   "subtraction results negative",
			expr:   strings.Split("5 - 10", " "),
			maxRes: 100,
			errMsg: "Under zero",
		},
		{
			name:   "addition results negative (sum <= 0)",
			expr:   strings.Split("-3 + 2", " "),
			maxRes: 100,
			errMsg: "Under zero",
		},
		{
			name:   "multiplication results negative",
			expr:   []string{"-5", internal.MultiplicationSymbol, "2"},
			maxRes: 100,
			errMsg: "Under zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMather(tt.expr, tt.maxRes)
			_, err := m.Calculate()
			if err == nil {
				t.Errorf("expected error containing %q, got nil", tt.errMsg)
				return
			}
			if !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
			}
		})
	}
}

func TestCalculateExpr_ExpectedError_ExceedMax(t *testing.T) {
	tests := []struct {
		name   string
		expr   []string
		maxRes int
		errMsg string
	}{
		{
			name:   "addition exceeds maxResult",
			expr:   strings.Split("100 + 50", " "),
			maxRes: 100,
			errMsg: "Over max border",
		},
		{
			name:   "multiplication exceeds maxResult",
			expr:   strings.Split("10 × 20", " "),
			maxRes: 150,
			errMsg: "Over max border",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMather(tt.expr, tt.maxRes)
			_, err := m.Calculate()
			if err == nil {
				t.Errorf("expected error containing %q, got nil", tt.errMsg)
				return
			}
			if !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
			}
		})
	}
}

func TestCalculateExpr_ExpectedError_DivisionZero(t *testing.T) {
	tests := []struct {
		name   string
		expr   []string
		maxRes int
		errMsg string
	}{
		{
			name:   "division by zero",
			expr:   strings.Split("10 ÷ 0", " "),
			maxRes: 100,
			errMsg: "Division by zero",
		},
		{
			name:   "non-integer division",
			expr:   strings.Split("7 ÷ 2", " "),
			maxRes: 100,
			errMsg: "Division by zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMather(tt.expr, tt.maxRes)
			_, err := m.Calculate()
			if err == nil {
				t.Errorf("expected error containing %q, got nil", tt.errMsg)
				return
			}
			if !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
			}
		})
	}
}

func TestCalculateExpr_ExpectedError_UnknownSymbol(t *testing.T) {
	tests := []struct {
		name   string
		expr   []string
		maxRes int
		errMsg string
	}{
		{
			name:   "unknown operator",
			expr:   strings.Split("5 % 3", " "),
			maxRes: 100,
			errMsg: "Invalid expression",
		},
		{
			name:   "invalid token",
			expr:   []string{"5", "?", "3"},
			maxRes: 100,
			errMsg: "Invalid expression",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMather(tt.expr, tt.maxRes)
			_, err := m.Calculate()
			if err == nil {
				t.Errorf("expected error containing %q, got nil", tt.errMsg)
				return
			}
			if !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
			}
		})
	}
}

func TestCalculateExpr_ExpectedError_InvalidExpression(t *testing.T) {
	tests := []struct {
		name   string
		expr   []string
		maxRes int
		errMsg string
	}{
		{
			name:   "empty expression",
			expr:   []string{},
			maxRes: 100,
			errMsg: "Invalid expression",
		},
		{
			name:   "only operands without operator",
			expr:   strings.Split("1 2 3", " "),
			maxRes: 100,
			errMsg: "Invalid expression",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMather(tt.expr, tt.maxRes)
			_, err := m.Calculate()
			if err == nil {
				t.Errorf("expected error containing %q, got nil", tt.errMsg)
				return
			}
			if !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("expected error containing %q, got %q", tt.errMsg, err.Error())
			}
		})
	}
}
