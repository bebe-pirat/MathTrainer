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
	Id             int
	OperandOrder   int
	EquationTypeId int
	MinValue       int
	MaxValue       int
}

type EquationTypeWithOperands struct {
	Id          int
	Operations  string
	NumOperands int
	NoRemainder bool
	MaxResult   int
	Operands    []Operand
}

type ShortEquationType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type EquationTypeRequest struct {
	Operations  string `json:"operations"`
	NumOperands int `json:"num_operands"`
	NoRemainder bool `json:"no_remainder"`
	MaxResult   int `json:"max_result"`
	Operands    []Operand `json:"operands"`
}
