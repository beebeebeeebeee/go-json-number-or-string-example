package app

import "encoding/json"

type DataRaw struct {
	Val json.Number `json:"val"`
	Str JsonString  `json:"str"`
}

type Data struct {
	Val float64 `json:"val"`
	Str string  `json:"str"`
}
