# go-json-number-or-string-example

this is the example to show how to use json.Number and a custom jsonString struct.

for the real world json, sometime the data we expected is a number but the json return as a string, and the data we expected is a string but the json return as a number. 

so we need to use json.Number and custom jsonString struct to handle this case.

## quick run
```bash
$ go run cmd/app/main.go
```

