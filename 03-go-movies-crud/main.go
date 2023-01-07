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

//no database will be used we will be sing structs and

//struct is basically like an object in javascript while slices are similar to arrays

type Movie struct {
	ID       string    `JSON:"id"`
	Isbn     string    `JSON:"isbn"`
	Title    string    `JSON:"title"`
	Director *Director `JSON:"director"` //pointer to the director struct, so all declarations in the director can be referenced here
}

type Director struct {
	Firstname string `JSON:"firstname"`
	Lastname  string `JSON:"lastname"`
}

var movies []Movie

// get all the movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") //set response content type to content / json
	json.NewEncoder(w).Encode(movies)                  //Json encode the response to send back
}

// delete a movie and return the remaining movies
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies { //looping through the array
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //slicing that one we don't want still don't understand this yet
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies { //using the blank operator cos goland will flag unused variables
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie //create a variable and assign the global movie struct

	_ = json.NewDecoder(r.Body).Decode(&movie) //convert the current movie body sent via api to json and
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies) //return all the movies with the updated movie
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set json content type
	params := mux.Vars(r)

	//delete an existing movie
	for index, item := range movies {
		if item.ID == params["id"] {
			log.Print("got here")
			movies = append(movies[:index], movies[index+1:]...)

			fmt.Printf("%+v\n", movies)

			var movie Movie
			//add a new movie into the setup
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "Adewumi", Lastname: "Adebayou"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{Firstname: "Hio", Lastname: "Sola"}})
	movies = append(movies, Movie{ID: "3", Isbn: "103929", Title: "Movie Trree", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	http.Handle("/", r)

	fmt.Printf("Server starting on port 8000\n") //prinf the place the serve is starting

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Couldn't connect to server")
	}

}
