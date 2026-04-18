package generator

import (
	"MathTrainer/internal"
	"strconv"
)

type Mather struct {
	infix             []string
	postfix           []string
	operationPriotiry map[string]int
	maxResult         int
}

func NewMather(infix_ []string, maxResult int) *Mather {
	return &Mather{
		infix:   infix_,
		postfix: make([]string, 0),
		operationPriotiry: map[string]int{
			internal.SummationSymbol:      1,
			internal.SubstractionSybmol:   1,
			internal.MultiplicationSymbol: 2,
			internal.DivisionSimbol:       2,
		},
		maxResult: maxResult,
	}
}

func (m *Mather) infixExprToPostfix() {
	output := make([]string, 0)
	stack := make([]string, 0)

	for _, token := range m.infix {
		_, err := strconv.Atoi(token)

		if err == nil {
			output = append(output, token)
			continue
		}

		switch token {
		case internal.OpenClaw:
			stack = append(stack, token)

		case internal.CloseClaw:
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if top == internal.OpenClaw {
					break
				}
				output = append(output, top)
			}

		case internal.SummationSymbol, internal.SubstractionSybmol, internal.DivisionSimbol, internal.MultiplicationSymbol:
			for len(stack) > 0 {
				top := stack[len(stack)-1]

				if top == internal.OpenClaw || !m.isOperator(token) {
					break
				}

				if m.operationPriotiry[token] > m.operationPriotiry[top] {
					break
				}

				output = append(output, top)
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, token)
		default:

		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top != internal.CloseClaw {
			output = append(output, top)
		}
	}

	m.postfix = output
}

func (m *Mather) isOperator(token string) bool {
	_, exists := m.operationPriotiry[token]
	return exists
}

func (m *Mather) calculatePostfix() (int, error) {
	if len(m.postfix) <= 0 {
		return 0, &CalculationError{"Invalid expression"}
	}

	stack := make([]int, 0)

	for _, token := range m.postfix {
		integer, err := strconv.Atoi(token)

		if err == nil {
			stack = append(stack, integer)
			continue
		}

		if m.isOperator(token) {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result, err := m.calculateOperation(b, a, token)

			if err != nil {
				return 0, err
			}

			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, &CalculationError{"Invalid expression"}
	}

	return stack[0], nil
}

func (m *Mather) calculateOperation(a, b int, op string) (int, error) {
	switch op {
	case internal.SummationSymbol:
		if a+b <= 0 {
			return 0, &CalculationError{"Under zero"}
		}
		if a+b > m.maxResult {
			return 0, &CalculationError{"Over max border"}
		}
		return a + b, nil

	case internal.SubstractionSybmol:
		if a-b <= 0 {
			return 0, &CalculationError{"Under zero"}
		}
		return a - b, nil

	case internal.MultiplicationSymbol:
		if a*b <= 0 {
			return 0, &CalculationError{"Under zero"}
		}
		if a*b > m.maxResult {
			return 0, &CalculationError{"Over max border"}
		}
		return a * b, nil

	case internal.DivisionSimbol:
		if b == 0 || a%b != 0 {
			return 0, &CalculationError{"Division by zero"}
		}
		return a / b, nil
	}
	return 0, &CalculationError{"Unknown operation"}
}

type CalculationError struct {
	Message string
}

func (e *CalculationError) Error() string {
	return "Calculation error: " + e.Message
}

func (m *Mather) Calculate() (int, error) {
	m.infixExprToPostfix()

	return m.calculatePostfix()
}
