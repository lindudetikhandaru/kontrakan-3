package usecase

import (
	"testing"
	"yayayaya/internal/repository"
)

func TestGetAllKontrakan(t *testing.T) {
	repo := repository.NewKontrakanRepository()
	uc := NewKontrakanUsecase(repo)

	result := uc.GetAllKontrakan()
	if len(result) == 0 {
		t.Errorf("Harusnya data kontrakan tidak boleh kosong")
	}
}
