package main

import (
	"go-json-number-or-string-example/internal/app"
	"log"
)

func main() {
	a := app.NewApp()

	db := []byte("[{\"valF64\": 1.23, \"valI64\": 188, \"str\": \"hello\", \"bol\": \"true\"}, {\"valF64\": \"4.56\", \"valI64\": \"512\", \"str\": 123, \"bol\": false}]")

	d, err := a.DecodeData(db)
	if err != nil {
		log.Fatalf("error while decoding data: %v", err)
	}

	for _, v := range d {
		log.Println(v.ValF64, v.ValI64, v.Str, v.Bol)
	}
}
