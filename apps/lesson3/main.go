package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"os"
	"agato/go-exercise/apps/lesson3/trace"
	"flag"
	"agato/go-exercise/apps/lesson3/client"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
