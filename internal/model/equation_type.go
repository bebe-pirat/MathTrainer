package model

type EquationType struct {
	Id          int    `json:"id"`
	Class       int    `json:"class"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Operations  string `json:"operations"`
	NumOperands int    `json:"num_operands"`
	NoRemainder bool   `json:"no_remainder"`
	MaxResult   int    `json:"max_result"`
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
	Operations  string    `json:"operations"`
	NumOperands int       `json:"num_operands"`
	NoRemainder bool      `json:"no_remainder"`
	MaxResult   int       `json:"max_result"`
	Operands    []Operand `json:"operands"`
}

type CreateOperandRequest struct {
	OperandOrder int `json:"operand_order"`
	MinValue     int `json:"min_value"`
	MaxValue     int `json:"max_value"`
}

type CreateEquationTypeRequest struct {
	Class       int                    `json:"class"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Operations  string                 `json:"operations"`
	NumOperands int                    `json:"num_operands"`
	NoRemainder bool                   `json:"no_remainder"`
	MaxResult   int                    `json:"max_result"`
	Operands    []CreateOperandRequest `json:"operands"`
}

type UpdateOperandRequest struct {
	Id           int `json:"id"`
	OperandOrder int `json:"operand_order"`
	MinValue     int `json:"min_value"`
	MaxValue     int `json:"max_value"`
}

type UpdateEquationTypeRequest struct {
	Class       int                    `json:"class"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Operations  string                 `json:"operations"`
	NumOperands int                    `json:"num_operands"`
	NoRemainder bool                   `json:"no_remainder"`
	MaxResult   int                    `json:"max_result"`
	Operands    []UpdateOperandRequest `json:"operands"`
}

type OperandResponse struct {
	Id           int `json:"id"`
	OperandOrder int `json:"operand_order"`
	MinValue     int `json:"min_value"`
	MaxValue     int `json:"max_value"`
}

type SectionAndEquationType struct {
	SectionId        int    `json:"section_id"`
	SectionName      string `json:"section_name"`
	EquationTypeId   int    `json:"equation_type_id"`
	EquationTypeName string `json:"equation_type_name"`
	Class            int    `json:"class"`
}
