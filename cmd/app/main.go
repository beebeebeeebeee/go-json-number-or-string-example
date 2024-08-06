package main

import (
	"json-number/internal/app"
	"log"
)

func main() {
	a := app.NewApp()

	db := []byte("[{\"val\": 1.23, \"str\": \"hello\"}, {\"val\": \"4.56\", \"str\": 123}]")

	d := a.DecodeData(db)

	for _, v := range d {
		log.Println(v.Val, v.Str)
	}
}
