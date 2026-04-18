package model

type EquationType struct {
	Id          int
	Class       int
	Name        string
	Description string
	Operations  string
	NumOperands int
	NoRemainder bool
	MaxResult   int
}

type Operand struct {
	Id           int
	OperandOrder int
	MinValue     int
	MaxValue     int
}

type EquationTypeWithOperands struct {
	Id          int
	Operations  string
	NumOperands int
	NoRemainder bool
	MaxResult   int
	Operands    []Operand
}
