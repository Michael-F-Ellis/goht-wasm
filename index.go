// The application index file
package main

import (
	"bytes"

	. "github.com/Michael-F-Ellis/goht"
)

// mkIndex constructs and renders the index content that is
// served to initiate the application.
func mkIndex() (b bytes.Buffer) {
	head := Head(``,
		Meta(`charset="utf-8"`),
		Title(``, "Go wasm"),
	)
	body := Body(`id="thebody"`,
		// wasm_exec.js is shipped with Go to implement
		// the JS side of syscall/js
		Script(`src="/wasm_exec.js"`, ""),
		// The boot script that loads and launches the application's
		// WebAssembly content
		Script(``, `
	if (!WebAssembly.instantiateStreaming) { // polyfill
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}

		const go = new Go();

		let mod, inst;

		WebAssembly.instantiateStreaming(fetch("/main.wasm"), go.importObject).then(async (result) => {
			mod = result.module;
			inst = result.instance;
			await go.run(inst);
		});`),
	)
	// concatenate and render the html document
	html := Html(``, head, body)
	err := Render(html, &b, 0)
	if err != nil {
		panic(err)
	}
	return
}
