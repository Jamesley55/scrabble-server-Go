package utils

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func GenerateID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		uuidError := errors.New("connot find new random in UUID")
		fmt.Print(uuidError)
	}
	return id.String()
}
