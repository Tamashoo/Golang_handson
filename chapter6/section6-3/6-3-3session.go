package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
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

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

func hello(w http.ResponseWriter, rq *http.Request, temps *template.Template) {
	msg := "login name & password:"

	ses, _ := cs.Get(rq, "hello-session")

	if rq.Method == "POST" {
		ses.Values["login"] = nil
		ses.Values["name"] = nil
		nm := rq.PostFormValue("name")
		pw := rq.PostFormValue("pass")
		if nm == pw {
			ses.Values["login"] = true
			ses.Values["name"] = nm
		}
		ses.Save(rq, w)
	}

	flg, _ := ses.Values["login"].(bool)
	lname, _ := ses.Values["name"].(string)
	if flg {
		msg = "logined: " + lname
	}

	item := struct {
		Title   string
		Message string
	}{
		Title:   "Session",
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
