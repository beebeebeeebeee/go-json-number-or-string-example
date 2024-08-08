package app

import (
	"encoding/json"
	"errors"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) DecodeData(data []byte) ([]Data, error) {
	var dr []DataRaw
	if err := json.Unmarshal(data, &dr); err != nil {
		return nil, errors.New("error while unmarshalling data: " + err.Error())
	}

	var ds []Data
	for _, v := range dr {
		ds = append(ds, Data{
			ValF64: v.ValF64.Float64(),
			ValI64: v.ValI64.Int64(),
			Str:    v.Str.String(),
			Bol:    v.Bol.Boolean(),
		})
	}

	return ds, nil
}
