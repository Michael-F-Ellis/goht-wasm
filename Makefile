# variables used to fetch the installed go versions
# wasm_exec.js
GOROOT := $(shell go env GOROOT)
WASMEXEC := $(GOROOT)/misc/wasm/wasm_exec.js
SERVER := goht-wasm
# pseudo-targets
.phony: all clean

# Default target
all: wasm/main.wasm $(SERVER) wasm_exec.js

# Executable that serves the demo
$(SERVER): server.go index.go
	go build -o $(SERVER)

# WebAssembly that runs client-side
wasm/main.wasm: wasm/main.go
	cd wasm ; GOOS=js GOARCH=wasm go build -o main.wasm

# Go distribution js file that implements client-side
# interface to syscall/js.
wasm_exec.js: $(WASMEXEC)
	cp $(WASMEXEC) .

# Removes all target files.
clean:
	-rm -f wasm/main.wasm $(SERVER) wasm_exec.js
