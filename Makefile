all: tinygo go 
.PHONY: all

tinygo:
	tinygo build -o ./wasm/tiny.wasm -target wasm ./main.go

go:
	GOOS=js GOARCH=wasm go build -o ./wasm/go.wasm