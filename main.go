package main

import (
	"net/http"
	"path"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the HTML template
		tmpl, err := template.ParseFiles("templates/image.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Construct the image URL
		imagePath := "images/image.jpg"
		imageURL := path.Join("/", imagePath)

		// Render the HTML template with the image URL
		data := struct {
			ImageURL string
		}{ImageURL: imageURL}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))

	http.ListenAndServe(":8080", nil)
}
