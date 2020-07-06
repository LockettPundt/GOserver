package routes

import (
	"fmt"
	"net/http"

	models "github.com/LockettPundt/GOserver/models"
)

// HelloHandler is the controller for the /hello route
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

// FormHandler is the controller for the /form.html route
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	title := r.FormValue("title")
	releaseDate := r.FormValue("releaseDate")
	rating := r.FormValue("rating")
	status := r.FormValue("status")
	models.AddMovie(title, releaseDate, rating, status)

	fmt.Fprintf(w, "Title = %s\n", title)
	fmt.Fprintf(w, "date = %s\n", releaseDate)
	fmt.Fprintf(w, "rating = %T\n", rating)
	fmt.Fprintf(w, "status = %s\n", status)

}
