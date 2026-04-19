package internal

const LevelsInSectionCoef = 3

// how much stars is given for percentage of right answers
const (
	OneStarsPercent   = 0.5
	TwoStarsPercent   = 0.7
	ThreeStarsPercent = 0.9
	MinStars          = 0
	MaxStars          = 3
)

// addition xp for answers
const (
	RightAnswerRegularExpressionXP = 10
	RightAnswerMixedExpression     = 15
	WrongAnswer                    = -5

	LevelWithoutMistakes = 50
)

// equation types weights
const (
	WeakWeight   = 2.0
	MediumWeight = 1.0
	StrongWeight = 0.5
	NewWeight    = 1.5
)

const (
	WeakCategoryAccurary   = 0.7
	MegiumCategoryAccuracy = 0.9
)

const CountEquationInSet = 10

const (
	DivisionSimbol       = "÷"
	MultiplicationSymbol = "×"
	SummationSymbol      = "+"
	SubstractionSybmol   = "-"
	OpenClaw             = "("
	CloseClaw            = ")"
)
