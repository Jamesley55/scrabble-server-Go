package models

type MessageType string

type ChatMessage struct {
	PlayerName string
	Data       string
	From       MessageType
}
