package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tmpl *template.Template

func main() {
	// Parse template at startup
	var err error
	tmplPath := filepath.Join("internal", "templates", "index.html")
	tmpl, err = template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	// Define a handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	addr := ":8080"
	log.Printf("Server is running on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
