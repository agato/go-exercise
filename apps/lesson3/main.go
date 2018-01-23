package main

import (
	"github.com/agato/go-exercise/apps/lesson3/client"
	"github.com/agato/go-exercise/apps/lesson3/trace"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("create template1")
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	t.templ.Execute(w, data)
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	port := os.Getenv("PORT")

	flag.Parse() // フラグを解釈します

	r := client.NewRoom()
	r.Tracer = trace.New(os.Stdout)
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	http.Handle("/room", r)

	go r.Run()

	log.Println("Webサーバーを開始します。ポート: ", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
