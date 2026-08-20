package repository

import "yayayaya/internal/model"

type KontrakanRepository struct {
	Data []model.Kontrakan
}

func NewKontrakanRepository() *KontrakanRepository {
	return &KontrakanRepository{
		Data: []model.Kontrakan{
			{ID: 1, Nama: "Kontrakan Ancol Sejahtera", Kota: "Jakarta", HargaBulan: 2500000, Status: "available"},
			{ID: 2, Nama: "Kost Mahasiswa Depok Indah", Kota: "Depok", HargaBulan: 1500000, Status: "booked"},
		},
	}
}

func (r *KontrakanRepository) GetAll() []model.Kontrakan { return r.Data }
func (r *KontrakanRepository) Insert(k model.Kontrakan)  { r.Data = append(r.Data, k) }
func (r *KontrakanRepository) UpdateStatus(id int, status string) bool {
	for i := range r.Data {
		if r.Data[i].ID == id {
			r.Data[i].Status = status
			return true
		}
	}
	return false
}
