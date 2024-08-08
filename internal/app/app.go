package app

import (
	"encoding/json"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) DecodeData(data []byte) []Data {
	var dr []DataRaw
	if err := json.Unmarshal(data, &dr); err != nil {
		panic(err)
	}

	var ds []Data
	for _, v := range dr {
		val, err := v.Val.Float64()
		if err != nil {
			panic(err)
		}

		ds = append(ds, Data{
			Val: val,
			Str: v.Str.String(),
			Bol: v.Bol.Boolean(),
		})
	}

	return ds
}
