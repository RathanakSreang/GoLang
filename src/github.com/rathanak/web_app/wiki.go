package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
  )

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename :=p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  return &Page{ Title: title, Body: body}, nil
}

func main() {
  // p1 := &Page{Title: "TestingPage", Body: []byte("This is a sample page")}
  // p1.save()
  // p2, _ := loadPage("TestingPage")
  // fmt.Println(string(p2.Body))
  http.HandleFunc("/", handler)
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8282", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hi I am love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, err := loadPage(title)
  fmt.Println(err)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
