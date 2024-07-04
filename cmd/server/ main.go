package main

import (
	"fmt"
	"net/http"

	"github.com/raj921/javascript-file-scanner/internal/handlers"
)

func main() {
	http.HandleFunc("/scan", handlers.ScanHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", serveTemplate("web/templates/index.html"))

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func serveTemplate(tmplName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, tmplName)
	}
}
