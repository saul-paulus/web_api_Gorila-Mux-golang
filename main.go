package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/v1/api/mahasiswa", getAllMahasiswa).Methods("GET")
	router.HandleFunc("/v1/api/mahasiswa", createMahasiswa).Methods("POST")
	router.HandleFunc("/v1/api/mahasiswa/id={id}", getIdMahasiswa).Methods("GET")
	router.HandleFunc("/v1/api/mahasiswa/id={id}", updateMahasiswa).Methods("PUT")
	router.HandleFunc("/v1/api/mahasiswa/id={id}", deleteMahasiswa).Methods("DELETE")

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
var mahasiswa Mahasiswa

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
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(dataMahasiswa)
}

func createMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewDecoder(r.Body).Decode(&mahasiswa)
	dataMahasiswa = append(dataMahasiswa, mahasiswa)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Data berhasil disimpan"}`))
}

func getIdMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	params := mux.Vars(r)
	id := params["id"]
	resInput, _ := strconv.Atoi(id)
	for _, u := range dataMahasiswa {
		if u.ID == resInput {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
}

func updateMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	var updateMhs Mahasiswa
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	json.NewDecoder(r.Body).Decode(&updateMhs)
	for i, u := range dataMahasiswa {

		if u.ID == id {
			// dataMahasiswa[i].ID = reqBody.ID
			u.Nim = updateMhs.Nim
			u.Nama = updateMhs.Nama
			u.Jurusan = updateMhs.Jurusan

			dataMahasiswa = append(dataMahasiswa[:i], u)
		}
	}
	w.Write([]byte(`{"message":"Data berhasil diupdate"}`))
}

func deleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, u := range dataMahasiswa {
		if u.ID == id {
			dataMahasiswa = append(dataMahasiswa[:i], dataMahasiswa[i+1:]...)
		}
	}
	w.Write([]byte(`{"message":"Data berhasil di hapus"}`))
}
