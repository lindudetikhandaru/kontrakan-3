package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yayayaya/internal/model"
	"yayayaya/internal/repository"
	"yayayaya/internal/usecase"
)

func main() {
	repo := repository.NewKontrakanRepository()
	uc := usecase.NewKontrakanUsecase(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/kontrakan", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // <-- Tambahkan baris ini supaya web bisa akses API

		// GET: Lihat Data
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(uc.GetAllKontrakan())
			return
		}

		// POST: Tambah Data (CRUD)
		if r.Method == http.MethodPost {
			var k model.Kontrakan
			json.NewDecoder(r.Body).Decode(&k)
			uc.AddKontrakan(k)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"msg": "Berhasil tambah!"})
			return
		}
	})

	// Endpoint Booking (Transaksi Kompleks)
	mux.HandleFunc("/api/booking", func(w http.ResponseWriter, r *http.Request) {
		id := 1 // Contoh ID 1
		err := uc.BookingKontrakan(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Booking berhasil untuk ID %d!", id)
	})

	fmt.Println("🚀 Server berjalan di http://localhost:8080/api/kontrakan")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Gagal:", err)
	}
}
