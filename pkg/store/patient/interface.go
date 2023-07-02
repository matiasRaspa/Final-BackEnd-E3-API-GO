package store

import "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"

type StoreInterface interface {
	Create(patient *domain.Patient) error
	Read(id int) (*domain.Patient, error)
	Update(patient *domain.Patient) error
	Patch(id int, address string) error
	Delete(id int) error
	List() ([]*domain.Patient, error)
}
