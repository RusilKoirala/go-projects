package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

var Movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Movies)

	if err != nil {
		log.Panic(err)
	}

}

func delMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Movies {
		if item.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "No movie ID was sent",
			})
			return
		}
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Movie not found",
	})
}

func getMoviebyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.IntN(10000000))
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func editMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			Movies = append(Movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}

}

func main() {
	r := mux.NewRouter()

	Movies = append(Movies, Movie{ID: "1", Isbn: "423242", Title: "Lord of the Rings", Director: &Director{Name: "Sir Peter Jackson"}})

	// Routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviebyID).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", editMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", delMovie).Methods("DELETE")

	// Start the server
	fmt.Println("Server has started on http://localhost:3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}

}
