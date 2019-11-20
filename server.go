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
	mux.HandleFunc("/main.wasm", sendWasm)
	mux.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wasm_exec.js")
	})
	log.Fatal(http.ListenAndServe(*listen, mux))
}

// indexPage serves the index page content
func indexPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(indexBuf.Bytes())
	if err != nil {
		log.Fatalf("Writing index buffer returned error: %v", err)
	}
}

// sendWasm serves main.wasm
func sendWasm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "wasm/main.wasm")
}
