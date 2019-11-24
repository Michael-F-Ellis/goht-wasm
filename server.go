// server is the application server
package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
)

var (
	listen   = flag.String("listen", ":8080", "listen address")
	indexBuf bytes.Buffer // holds index page html
)

// init generates the contents of indexBuf
func init() {
	indexBuf = mkIndex()
}

// main parses the flag arguments, creates a ServeMux, registers
// HandleFuncs and starts serving http requests.
func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexPage)
	mux.HandleFunc("/static/main.wasm", fsSendStaticWasm)
	mux.HandleFunc("/static/wasm_exec.js", fsSendStaticJS)
	log.Fatal(http.ListenAndServe(*listen, mux))
}

// indexPage serves the index page content
func indexPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(indexBuf.Bytes())
	if err != nil {
		log.Fatalf("Writing index buffer returned error: %v", err)
	}
}

// fsSendStaticJS serves the requested file from the ESC FS.
func fsSendStatic(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(FSMustByte(false, r.URL.Path))
	if err != nil {
		log.Fatalf("Writing file %s returned error: %v", r.URL.Path, err)
	}
}

// fsSendStaticJS serves the requested javascript file from the ESC FS.
func fsSendStaticJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fsSendStatic(w, r)
}

// fsSendStaticWasm serves the requested WebAssembly file  from the ESC FS.
func fsSendStaticWasm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	fsSendStatic(w, r)
}
