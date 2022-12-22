package models

type ObjectiveType string

const (
	Private ObjectiveType = "Private"
	Public  ObjectiveType = "Public"
)

type Objective struct {
	Name        string
	FullName    string
	Points      int
	IsCompleted bool
	Type        ObjectiveType
}
