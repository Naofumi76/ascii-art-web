package main

import (
	ascii "ascii-art-web/asciiart" // Adjust the import path as per your project structure
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// initiate the server
	server := &http.Server{
		Addr:              ":8080",           // port number
		Handler:           nil,               // default handler
		ReadHeaderTimeout: 10 * time.Second,  // read timeout
		WriteTimeout:      10 * time.Second,  // write timeout
		IdleTimeout:       120 * time.Second, // idle timeout
		MaxHeaderBytes:    1 << 20,           // max header bytes
	}

	// Handlers
	http.HandleFunc("/", HandleHome)              // handle main page
	http.HandleFunc("/download", DownloadHandler) // handle download endpoint

	// Start the server
	log.Fatal(server.ListenAndServe())
}

type PageData struct {
	Resultat string
	Title    string
}

var (
	tpl, _    = template.ParseFiles("body.html")
	err404, _ = template.ParseFiles("err404.html")
	data      = PageData{
		Resultat: "",
		Title:    "Ascii-Art",
	}
)

// HandleHome handles the main page requests
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// Management of 404 error
	if r.URL.Path != "/" {
		err := err404.Execute(w, nil)
		log.Printf("Error %v", http.StatusNotFound)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Handle form submission
	if r.Method == http.MethodPost {
		opt := r.FormValue("options")
		result := r.FormValue("send_value")
		if result != "" {
			tmp, err := ascii.AsciiArt(result, opt) // assuming asciiart.AsciiArt is your function to generate ASCII art
			if err != nil {
				data.Resultat = "Error 500 : Wrong input !"
				log.Printf("Error %v", http.StatusInternalServerError)
			} else {
				data.Resultat = tmp
			}
		}
	}

	// Render the template
	err := tpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Status OK : %v", http.StatusOK)
	}
}

// DownloadHandler handles the download requests
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Extract text from query parameter
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "No text to download", http.StatusBadRequest)
		return
	}

	// Set headers for download
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Length", fmt.Sprint(len(text)))

	// Write text content
	if _, err := w.Write([]byte(text)); err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}
}
