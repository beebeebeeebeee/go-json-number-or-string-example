package app

import "encoding/json"

type DataRaw struct {
	Val json.Number `json:"val"`
	Str JsonString  `json:"str"`
	Bol JsonBoolean `json:"bol"`
}

type Data struct {
	Val float64 `json:"val"`
	Str string  `json:"str"`
	Bol bool    `json:"bol"`
}
