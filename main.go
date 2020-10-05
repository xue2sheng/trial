package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var port = "3333"
var assets = ""

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	log.Printf("Launching server at %s port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	log.Println(exPath)

	var help = false
	flag.BoolVar(&help, "help", help, "Show this help.")
	flag.StringVar(&port, "port", port, "Port to launch the static server.")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
}
