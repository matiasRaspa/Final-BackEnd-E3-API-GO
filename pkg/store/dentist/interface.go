package store

import "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"

type StoreInterface interface {
	Create(dentist *domain.Dentist) error
	Read(id int) (*domain.Dentist, error)
	Update(dentist *domain.Dentist) error
	Patch(id int, license string) error
	Delete(id int) error
	List() ([]*domain.Dentist, error)
}
