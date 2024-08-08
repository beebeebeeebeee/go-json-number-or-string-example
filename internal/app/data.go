package app

type DataRaw struct {
	ValF64 JsonFloat64 `json:"valF64"`
	ValI64 JsonInt64   `json:"valI64"`
	Str    JsonString  `json:"str"`
	Bol    JsonBoolean `json:"bol"`
}

type Data struct {
	ValF64 float64 `json:"valF64"`
	ValI64 int64   `json:"valI64"`
	Str    string  `json:"str"`
	Bol    bool    `json:"bol"`
}
