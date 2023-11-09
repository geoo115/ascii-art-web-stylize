package main

import (
	art "asciiartweb/ascii"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// PageData struct holds the data to be passed to the template.
type PageData struct {
	AsciiArt string
}

func main() {

	fs := http.FileServer(http.Dir("static")) // Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleDefault) // Define the handlers for different routes
	http.HandleFunc("/ascii-art", handleAsciiArt)
	fmt.Println("Server started. Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleDefault(w http.ResponseWriter, r *http.Request) { // Check if the request URL path is exactly "/" and Send a 404 Not Found error page
	if r.URL.Path != "/" { // r.URL.Path
		renderErrorPage(w, http.StatusNotFound)
		return
	}
	renderTemplate(w, "templates/index.html", nil)
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) == 0 {
		renderErrorPage(w, http.StatusBadRequest)
		return
	}

	result, err := art.Generate(text, banner)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := PageData{
		AsciiArt: result,
	}

	json.NewEncoder(w).Encode(data)
}

// renderTemplate renders the specified template with the provided data.
func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl, err := template.ParseFiles(templateFile) // Parse the template file
	if err != nil {
		renderErrorPage(w, http.StatusNotFound) // Send a 404 Not Found error page
		return
	}
	err = tmpl.Execute(w, data) // Execute the template with the provided data
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError) // Send a 500 Internal Server Error page
	}
}

type ErrorData struct {
	StatusCode string
}

func renderErrorPage(w http.ResponseWriter, status int) {
	data := ErrorData{
		StatusCode: fmt.Sprintf("%d", status),
	}

	// Parse the error page template with the status code
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%d.html", status))
	if err != nil {
		log.Printf("Error parsing error template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status) // Set the status code for the response

	// Execute the error page template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing error template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
