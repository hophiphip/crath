package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"math.io/crath/ord"
)

const PORT = "8000"
const HOST = "0.0.0.0"

type FunctionCase struct {
	Element string `json:"element"`
	Module  string `json:"module"`
}

type FunctionSoluton struct {
	Ord string `json:"ord"`
}

func ordRequest(w http.ResponseWriter, r *http.Request) {
	var funcCase FunctionCase

	element := new(big.Int)
	module := new(big.Int)

	if err := json.NewDecoder(r.Body).Decode(&funcCase); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	element, okElement := element.SetString(funcCase.Element, 10)
	module, okModule := module.SetString(funcCase.Module, 10)

	if !okElement || !okModule {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Parsing error"))
	} else {
		solution := FunctionSoluton{Ord: ord.Ord(element, module).String()}
		json.NewEncoder(w).Encode(solution)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "index.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func requests() {
	router := mux.NewRouter().StrictSlash(true)

	fs := handlers.CombinedLoggingHandler(os.Stderr, http.FileServer(http.Dir("./static")))
	router.Handle("/static/", fs)
	router.HandleFunc("/", serveTemplate)
	router.HandleFunc("/ord", ordRequest).Methods("POST")

	log.Printf("Stareted server on: %s:%s\n", HOST, PORT)

	log.Fatal(http.ListenAndServe(HOST+":"+PORT, router))
}

func main() {
	requests()
}
