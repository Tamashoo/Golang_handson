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

	helo, er := template.ParseFiles("./templates/hello.html")
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

var flg bool = true

func hello(w http.ResponseWriter, rq *http.Request, temps *template.Template) {
	item := struct {
		Flg      bool
		Title    string
		Message  string
		JMessage string
	}{
		Flg:      flg,
		Title:    "Send values",
		Message:  "This is Sample message.",
		JMessage: "これはサンプルです。",
	}

	er := temps.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
	flg = !flg
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
