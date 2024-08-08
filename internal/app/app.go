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
		ds = append(ds, v.Data())
	}

	return ds, nil
}
