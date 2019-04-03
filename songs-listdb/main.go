package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Song to create a song list
type Song struct {
	ID       int    `json:id`
	Title    string `json:title`
	Author   string `json:author`
	MusicURL string `json:musicurl`
}

var songs []Song

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/songs", getSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", getSong).Methods("GET")
	router.HandleFunc("/songs", addSong).Methods("POST")
	router.HandleFunc("/songs/{id}", updateSong).Methods("PUT")
	router.HandleFunc("/songs/{id}", removeSong).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func getSongs(w http.ResponseWriter, r *http.Request) {

}

func getSong(w http.ResponseWriter, r *http.Request) {
}

func addSong(w http.ResponseWriter, r *http.Request) {

}

func updateSong(w http.ResponseWriter, r *http.Request) {

}

func removeSong(w http.ResponseWriter, r *http.Request) {
}
