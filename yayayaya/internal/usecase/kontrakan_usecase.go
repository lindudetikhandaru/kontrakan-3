package usecase

import (
	"errors"
	"yayayaya/internal/model"
	"yayayaya/internal/repository"
)

type KontrakanUsecase struct {
	repo *repository.KontrakanRepository
}

func NewKontrakanUsecase(repo *repository.KontrakanRepository) *KontrakanUsecase {
	return &KontrakanUsecase{repo: repo}
}

func (u *KontrakanUsecase) GetAllKontrakan() []model.Kontrakan { return u.repo.GetAll() }
func (u *KontrakanUsecase) AddKontrakan(k model.Kontrakan)     { u.repo.Insert(k) }
func (u *KontrakanUsecase) BookingKontrakan(id int) error {
	for _, k := range u.repo.GetAll() {
		if k.ID == id {
			if k.Status == "booked" {
				return errors.New("kamar penuh")
			}
			u.repo.UpdateStatus(id, "booked")
			return nil
		}
	}
	return errors.New("tidak ditemukan")
}
