package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	tplDir = "templates/*"
	tpl *template.Template

	localIP = ""
	dockerIP = ""
	currentIP = ""

	client *http.Client
)

func init() {
	tpl = template.Must(template.ParseGlob(tplDir))
	if tpl == nil {
		log.Fatalf("template.Must(template.ParseGlob(\"%s\")", tplDir)
	}

	client = &http.Client{
		Timeout: time.Second * 15,
	}

	dockerPORT := os.Getenv("dockerPORT")
	if dockerPORT == "" {
		dockerPORT = "9001"
	}

	dockerHOST := os.Getenv("dockerHOST")
	if dockerHOST == "" {
		dockerHOST = "rest-api"
	}

	dockerIP = fmt.Sprintf("http://%s:%s", dockerHOST, dockerPORT)

	_, err := http.Get(dockerIP)
	if err != nil {
		log.Printf("http.Get(dockerIP): %s\n", err.Error())
		localHOST := os.Getenv("localHOST")
		if localHOST == "" {
			localHOST = "0.0.0.0"
		}

		currentIP = fmt.Sprintf("http://%s:%s", localHOST, dockerPORT)
	} else {
		currentIP = dockerIP
	}
	log.Printf("Backend is running at %s ...\n", currentIP)
}

func main() {
	uiHOST := os.Getenv("uiHOST")
	if uiHOST == "" {
		uiHOST = "0.0.0.0"
	}

	uiPORT := os.Getenv("uiPORT")
	if uiPORT == "" {
		uiPORT = "9042"
	}

	addr := fmt.Sprintf("%s:%s", uiHOST, uiPORT)

	r := chi.NewRouter()

	r.Post("/", createBook)
	r.Get("/", getBooks)
	r.Get("/{id}", getBookById)
	r.Get("/ping", getHeartbeat)
	r.Get("/create", createBook)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		root := http.Dir("static")
		fsrvr := http.StripPrefix("/static/", http.FileServer(root))
		fsrvr.ServeHTTP(w, r)
	})

	log.Printf("Serving at %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
