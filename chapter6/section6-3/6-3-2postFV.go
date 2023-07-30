package main

import (
	"log"
	"net/http"
	"text/template"
)

type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func setupTemp() *Temps {
	temps := new(Temps)
	temps.notemp = notemp()

	indx, er := template.ParseFiles("./templates/index.html")
	if er != nil {
		indx = temps.notemp
	}
	temps.indx = indx

	helo, er := template.ParseFiles("./templates/hello4.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo

	return temps
}

func index(w http.ResponseWriter, rq *http.Request, temps *template.Template) {
	er := temps.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

func hello(w http.ResponseWriter, rq *http.Request, temps *template.Template) {
	msg := "type name and password:"

	if rq.Method == "POST" {
		nm := rq.FormValue("name")
		pw := rq.FormValue("pass")
		msg = "name: " + nm + ", password: " + pw
	}

	item := struct {
		Title   string
		Message string
	}{
		Title:   "Send values",
		Message: msg,
	}

	er := temps.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	temps := setupTemp()

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq, temps.indx)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq, temps.helo)
	})

	http.ListenAndServe("", nil)
}
