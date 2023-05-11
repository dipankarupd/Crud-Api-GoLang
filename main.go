package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

// get all functions:
func getMovies(w http.ResponseWriter, r *http.Request) {

	// taking the content from the slice and creating a json for the instance
	// and sending the response in json format

	//this line sets the response header to indicate that the response will be in JSON format. The
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete movie function:
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// param has all the variables whose request has been sent
	param := mux.Vars(r)

	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range movies {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	var movie Movie

	for index, item := range movies {

		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(10000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

var movies []Movie

func main() {
	// creating a new Router using the gorilla mux:
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "456987", Title: "Everything Everywhere All at Once", Director: &Director{FirstName: "Dipankar", LastName: "Upadhyaya"}})
	movies = append(movies, Movie{ID: "2", Isbn: "456123", Title: "Saili", Director: &Director{FirstName: "Panas", LastName: "Babu"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("Delete")

	fmt.Printf("starting the server at port 8000\n")

	// starting the server:
	log.Fatal(http.ListenAndServe(":8000", r))
}
