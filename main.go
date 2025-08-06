package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		log.Println("Template execution error:", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about")
	})
	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "gallery")
	})
	http.HandleFunc("/idx", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "idx")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("HTTP server listening on 0.0.0.0:9090")
	if err := http.ListenAndServe("0.0.0.0:9090", nil); err != nil {
		log.Fatal(err)
	}

}
