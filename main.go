package main

import (
    "html/template"
    "net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, _ := template.ParseFiles("templates/" + tmpl + ".html")
    t.Execute(w, nil)
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

    http.ListenAndServe(":8080", nil)
}
