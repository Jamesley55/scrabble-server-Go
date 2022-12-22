package models

type CommandName string

type CommandError string

type Command struct {
	FullCommand string
	Name        string
}

type PlaceCommand struct {
	Command
	Row               string
	Column            int
	Direction         string
	Word              string
	WordsInDictionary bool
}

type ExchangeCommand struct {
	Command
	Letters string
}

type ClueCommand struct {
	Command
	PlayableWords []PlayableWord
}

type HelpCommand struct {
	Command
	Message string
}
