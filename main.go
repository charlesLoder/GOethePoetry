package main

import (
	"html/template"
	"log"
	"net/http"
)


func main(){
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
      tmpl := template.Must(template.ParseFiles("templates/index.html"))
      tmpl.Execute(w, nil)
    })

    http.HandleFunc("/poem", func(w http.ResponseWriter, r *http.Request) {
      log.Println("Finding a poem...")

      tmpl := template.Must(template.ParseFiles("templates/fragments/poem.html"))

      data, err := GetRandomPoem()

      if err != nil {
        log.Println("Error getting poem: ", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      tmpl.Execute(w, data)
    })

    log.Println("Listening on http://localhost:8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}