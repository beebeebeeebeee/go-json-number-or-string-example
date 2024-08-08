package app

type Data struct {
	ValF64 float64 `json:"valF64"`
	ValI64 int64   `json:"valI64"`
	Str    string  `json:"str"`
	Bol    bool    `json:"bol"`
}

type DataRaw struct {
	ValF64 JsonFloat64 `json:"valF64"`
	ValI64 JsonInt64   `json:"valI64"`
	Str    JsonString  `json:"str"`
	Bol    JsonBoolean `json:"bol"`
}

func (r *DataRaw) Data() Data {
	return Data{
		ValF64: r.ValF64.Float64(),
		ValI64: r.ValI64.Int64(),
		Str:    r.Str.String(),
		Bol:    r.Bol.Boolean(),
	}
}
