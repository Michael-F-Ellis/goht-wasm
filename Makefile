# variables used to fetch the installed go versions
# wasm_exec.js
GOROOT := $(shell go env GOROOT)
WASMEXEC := $(GOROOT)/misc/wasm/wasm_exec.js

# pseudo-targets
.phony: all clean

# Default target
all: wasm/main.wasm wasmtut wasm_exec.js

# Executable that serves the demo
wasmtut: server.go index.go
	go build

# WebAssembly that runs client-side
wasm/main.wasm: wasm/main.go
	cd wasm ; GOOS=js GOARCH=wasm go build -o main.wasm

# Go distribution js file that implements client-side
# interface to syscall/js.
wasm_exec.js: $(WASMEXEC)
	cp $(WASMEXEC) .

# Removes all target files.
clean:
	-rm -f wasm/main.wasm wasmtut wasm_exec.js
