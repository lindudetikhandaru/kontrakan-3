package model

type Kontrakan struct {
	ID         int     `json:"id"`
	Nama       string  `json:"nama"`
	Kota       string  `json:"kota"`
	HargaBulan float64 `json:"harga_bulan"`
	Status     string  `json:"status"` // "available" atau "booked"
}
