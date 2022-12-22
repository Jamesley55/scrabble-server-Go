package models

type ScoreMultiplier struct {
	TimeTwo         int
	TimeThree       int
	Point7Letters   int
	MaxAmountLetter int
}

var SCORE_MULTIPLIER = ScoreMultiplier{
	TimeTwo:         2,
	TimeThree:       3,
	Point7Letters:   50,
	MaxAmountLetter: 7,
}
