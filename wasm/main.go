// +build js,wasm

package main

import (
	"bytes"
	. "github.com/Michael-F-Ellis/goht"
	"strconv"
	"syscall/js"
)

// ui defines text fields and buttons displayed in the application
var ui = Null(
	Input(`type="text" id="value1"`),
	Input(`type="text" id="value2"`),
	Button(`onClick="add('value1', 'value2', 'result');" id="addButton"`, "Add"),
	Button(`onClick="subtract('value1', 'value2', 'result');" id="subtractButton"`, "Subtract"),
	Input(`type="text" id="result"`),
)

// injectUI renders ui elements into a div at the
// end of the body element.
func injectUI() {
	doc := js.Global().Get("document")
	div := doc.Call("getElementById", "content")
	b := bytes.Buffer{}
	_ = Render(ui, &b, 0)
	div.Set("innerHTML", b.String())
}

// getInput Values returns the current values of the two
// input fields as floats.
func getInputValues(i []js.Value) (v1, v2 float64) {
	value1 := getStringValueById(i[0].String())
	value2 := getStringValueById(i[1].String())
	v1, _ = strconv.ParseFloat(value1, 64)
	v2, _ = strconv.ParseFloat(value2, 64)
	return
}

// add sums the values in the input fields and
// puts the result into the output field
func add(this js.Value, i []js.Value) interface{} {
	in1, in2 := getInputValues(i)
	setValueById(i[2].String(), in1+in2)
	return nil
}

// subtract places the difference between the first input field
// and the second input field into the output field.
func subtract(this js.Value, i []js.Value) interface{} {
	int1, int2 := getInputValues(i)
	setValueById(i[2].String(), int1-int2)
	return nil
}

// registerCallbacks maps Go functions to JS event listeners
// named in application button event scripts.
func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

// main adds the ui elements to the docs, registers button
// callbacks and waits forever.
func main() {
	c := make(chan struct{}, 0)
	injectUI()
	registerCallbacks()
	println("Go WebAssembly Initialized")

	<-c
}
