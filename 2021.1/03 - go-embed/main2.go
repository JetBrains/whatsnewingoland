package main

import (
	"embed"
	"net/http"
)

//go:embed templates
var templates embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		idx, err := templates.ReadFile("templates/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(idx)
	})

	http.FileServer(http.FS(templates))

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
