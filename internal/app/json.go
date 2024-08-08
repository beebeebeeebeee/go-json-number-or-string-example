package app

import (
	"errors"
)

type JsonString string

func (j *JsonString) UnmarshalJSON(data []byte) error {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	*j = JsonString(data)
	return nil
}

func (j *JsonString) String() string {
	return string(*j)
}

type JsonBoolean bool

func (j *JsonBoolean) UnmarshalJSON(data []byte) error {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	switch string(data) {
	case "true":
		*j = true
	case "false":
		*j = false
	default:
		return errors.New("invalid boolean value")
	}

	return nil
}

func (j *JsonBoolean) Boolean() bool {
	return bool(*j)
}
