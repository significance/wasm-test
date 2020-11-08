package main
import "syscall/js"

// function definition
func add(this js.Value, i []js.Value) interface{} {
  return js.ValueOf(i[0].Int()+i[1].Int())
}

// function definition
func subtract(this js.Value, i []js.Value) interface{} {
  return js.ValueOf(i[0].Int()-i[1].Int())
}

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))
    js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}