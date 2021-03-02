package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/mahasiswa", getAllMahasiswa).Methods("GET")
	router.HandleFunc("/mahasiswa", createMahasiswa).Methods("POST")

	fmt.Println("server berjalan pada 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

type Mahasiswa struct {
	ID      int    `json:"id"`
	Nim     int    `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}

var err error

var dataMahasiswa = []Mahasiswa{
	Mahasiswa{3, 323233332, "rudi", "teknik sipil"},
	Mahasiswa{2, 232223333, "genji", "penjas"},
	Mahasiswa{4, 322333233, "zaki", "bahasa indonesia"},
	Mahasiswa{1, 232223233, "liu", "statistik"},
	Mahasiswa{22, 3232223232, "lie", "matematika"},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello saul")
}

func getAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataMahasiswa)
}

func createMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
