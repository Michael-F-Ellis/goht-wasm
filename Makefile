# variables used to fetch the installed go versions
# wasm_exec.js
GOROOT := $(shell go env GOROOT)
WASMEXEC := $(GOROOT)/misc/wasm/wasm_exec.js

## The executable
SERVER := goht-wasm

# pseudo-targets
.phony: all clean

# Default target
all: static/main.wasm $(SERVER) static/wasm_exec.js

# WebAssembly that runs client-side
static/main.wasm: wasm/main.go
	cd wasm ; GOOS=js GOARCH=wasm go build -o ../static/main.wasm

# Go distribution js file that implements client-side
# interface to syscall/js.
static/wasm_exec.js: $(WASMEXEC)
	cp $(WASMEXEC) static/

# Generate the go files that create the esc filesystem
static.go: static/wasm_exec.js static/main.wasm
	esc -ignore static/*.go -o static.go static

# Executable that serves the demo
$(SERVER): server.go index.go static.go
	go build -o $(SERVER)


# Removes all target files.
clean:
	-rm -f static/main.wasm $(SERVER) static/wasm_exec.js static.go
