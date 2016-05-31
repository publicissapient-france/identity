package main

import (
	"html/template"
	"net/http"
	"os"
	"log"
	"io"
	"encoding/json"
)

type Identity struct {
	Name     string        `json:"name"`
	Filename string        `json:"filename"`
	Hostname string        `json:"hostname"`
	Hits     int           `json:"hits"`
}

var templates = template.Must(template.ParseFiles("identity.html"))

func getImage(url string, filename string) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	file, err := os.Create("static/img/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func loadIdentity() (*Identity, error) {
	name := os.Getenv("NAME")
	if len(name) == 0 {
		name = "N/A"
	}

	filename := os.Getenv("FILENAME")
	if len(filename) == 0 {
		filename = "identity.png"
	}

	url := os.Getenv("URL")
	if len(url) != 0 {
		getImage(url, filename);
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Identity{Name: name, Filename: filename, Hostname: hostname, Hits: 0}, nil
}

var identity, err = loadIdentity()

func makeHandler(fn func(http.ResponseWriter, *http.Request, *Identity)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fn(w, r, identity)
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request, identity *Identity) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	js, err := json.Marshal(identity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(js)
}

func identityHandler(w http.ResponseWriter, r *http.Request, identity *Identity) {
	identity.Hits = identity.Hits + 1

	err := templates.ExecuteTemplate(w, "identity.html", identity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/identity", makeHandler(identityHandler))
	http.HandleFunc("/identity/json", makeHandler(jsonHandler))

	http.HandleFunc("/static/", staticHandler)

	http.ListenAndServe(":8080", nil)
}