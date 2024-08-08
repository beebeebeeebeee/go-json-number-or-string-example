package app

import (
	"errors"
	"strconv"
)

type JsonInt64 int64

func (j *JsonInt64) UnmarshalJSON(data []byte) error {
	data = removeQuotes(data)

	w, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return errors.New("invalid int64 value: " + err.Error())
	}

	*j = JsonInt64(w)
	return nil
}

func (j *JsonInt64) Int64() int64 {
	return int64(*j)
}

type JsonFloat64 float64

func (j *JsonFloat64) UnmarshalJSON(data []byte) error {
	data = removeQuotes(data)

	w, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return errors.New("invalid float64 value: " + err.Error())
	}

	*j = JsonFloat64(w)
	return nil
}

func (j *JsonFloat64) Float64() float64 {
	return float64(*j)
}

type JsonString string

func (j *JsonString) UnmarshalJSON(data []byte) error {
	data = removeQuotes(data)

	*j = JsonString(data)
	return nil
}

func (j *JsonString) String() string {
	return string(*j)
}

type JsonBoolean bool

func (j *JsonBoolean) UnmarshalJSON(data []byte) error {
	data = removeQuotes(data)

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

func removeQuotes(data []byte) []byte {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		return data[1 : len(data)-1]
	}

	return data
}
