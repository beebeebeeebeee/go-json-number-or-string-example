package main

import (
	"go-json-number-or-string-example/internal/app"
	"log"
)

func main() {
	a := app.NewApp()

	db := []byte("[{\"val\": 1.23, \"str\": \"hello\", \"bol\": true}, {\"val\": \"4.56\", \"str\": 123, \"bol\": \"false\"}]")

	d := a.DecodeData(db)

	for _, v := range d {
		log.Println(v.Val, v.Str, v.Bol)
	}
}
