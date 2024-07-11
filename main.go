package main

import (
	"fmt"
	"groupie-tracker/create"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist/", ArtistHandler)

	openbrowser("http://localhost:8080")
	fmt.Println("Server listening on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
func openbrowser(url string) error {
	var cmd string
	var args []string
    // Determine the appropriate command and arguments based on the user's operating system

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	    // Append the URL to the arguments

	args = append(args, url)

	    // Append the URL to the arguments

	return exec.Command(cmd, args...).Start()

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	    // Check if the request method is GET

	if r.Method == "GET" {
        // Check if the request URL path is not the root ("/")

		if r.URL.Path != "/" {
			t, _ := template.ParseFiles("templates/error.html")
			t.Execute(w, http.StatusNotFound)
			return
		}

		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			// If there's an error parsing the template, return a 500 Internal Server Error
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}

		// Call the `create.GetArtists()` function to retrieve a list of artists
		artist, err := create.GetArtists()
		if err != nil {

// If there's an error getting the artists, return a 500 Internal Server Error
			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		 // Execute the "index.html" template and pass the list of artists as the data
		t.Execute(w, artist)
	} else {
	        // If the request method is not GET, return a 400 Bad Request error
		http.Error(w, "400: Bad Request", http.StatusBadRequest)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET

	if r.Method == "GET" {
        // Parse the "profile.html" template

		t, err := template.ParseFiles("templates/profile.html")
		if err != nil {
			            // If there's an error parsing the template, return a 500 Internal Server Error

			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}
		        // Call the `create.GetArtists()` function to retrieve a list of artists

		artist, err := create.GetArtists()
		if err != nil {
			            // If there's an error getting the artists, return a 500 Internal Server Error

			http.Error(w, "500: internal server error", http.StatusInternalServerError)
			return
		}

        // Extract the artist name from the URL path

		urlString := string(r.URL.Path)[8:]

        // Iterate through the list of artists to find the one that matches the URL path

		for i, v := range artist {
			if v.Name == urlString {
	// If a matching artist is found, execute the "profile.html" template and pass the artist data
				t.Execute(w, artist[i])
				return
			}
		}
// If no matching artist is found, render the "error.html" template with a 404 Not Found status
		t, _ = template.ParseFiles("templates/error.html")
		t.Execute(w, http.StatusNotFound)
		return
	} else {
		http.Error(w, "400: bad request.", http.StatusBadRequest)
	}

}