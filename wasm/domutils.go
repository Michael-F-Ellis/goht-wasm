// +build js,wasm
// Utility functions that encapsulate some of the verbosity out of the syscall/js
// API.
package main

import (
	"syscall/js"
)

// getStringValueById returns the value property of the
// element with the given id.
func getStringValueById(id string) (value string) {
	return js.Global().Get("document").Call("getElementById", id).Get("value").String()
}

// setPropertyById stores the value v in Property property of
// the element with the given id.
func setPropertyById(id string, property string, v interface{}) {
	js.Global().Get("document").Call("getElementById", id).Set(property, v)
}

// setValueById stores the value v in Property "value"
// of the element with the given id
func setValueById(id string, v interface{}) {
	js.Global().Get("document").Call("getElementById", id).Set("value", v)
}
