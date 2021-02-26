package main

import (
  "fmt"
  "net/http"
  "html/template"
  "github.com/sputnik5459/username-finder/internal/engine"
)

type TemplateData struct {
  Title string
  Result string
}

func index(w http.ResponseWriter, r *http.Request){
  p := TemplateData{Title: "Index"}

  tmpls, _ := template.ParseFiles(
    "web/templates/index.html",
    "web/templates/find_username_form.html",
  )

  fmt.Println(tmpls.Execute(w,p))
}

func processor(w http.ResponseWriter, r *http.Request) {
  if r.Method != "Post" {
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }

  fname := r.FormValue("username_field")
  engine.FindUsername(fname)
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/process", processor)
  http.ListenAndServe(":8080", nil)
}
