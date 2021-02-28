package util

import "github.com/satori/go.uuid"

func GetUUID() string {
	uuidValue, _ := uuid.NewV4()
	return uuidValue.String()
}
