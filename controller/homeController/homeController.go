package homecontroller

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

func Navbar(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/layout/navbar.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Files(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	// Open the file
	file, err := os.Open("assets/files/"+name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the content type header to application/pdf
	w.Header().Set("Content-Type", "application/pdf")

	// Copy the PDF file to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Photo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	extension := filepath.Ext(name)
	// Open the file
	file, err := os.Open("assets/img/"+name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/"+strings.TrimPrefix(extension, "."))

	// Copy the PDF file to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}



