package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, "Failed to create the file on server", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save the file on server", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s uploaded successfully", header.Filename)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
