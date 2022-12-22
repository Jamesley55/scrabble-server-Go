package models

import grid "scrabble/app/models/grid"

type SCORE_PERCENTILE int

type Position struct {
	Row int
	Col int
}

type PlayableWord struct {
	Word      string
	Position  Position
	Direction string
	Score     *int
}

type ClueFunctionParams struct {
	PlayableWords []PlayableWord
	Anchors       []Position
	Grid          []grid.Node
	PlayerLetters []string
	WordDirection string
	PrefixLimit   int
}
