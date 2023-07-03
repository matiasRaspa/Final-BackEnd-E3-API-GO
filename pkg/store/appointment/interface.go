package store

import "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"

type StoreInterface interface {
	Create(appointment *domain.Appointment) error
	Read(id int) (*domain.Appointment, error)
	Update(appointment *domain.Appointment) error
	Patch(id int, dateTime string) error
	Delete(id int) error
	List() ([]*domain.Appointment, error)
	ListPatients() ([]*domain.Patient, error)
	ListDentists() ([]*domain.Dentist, error)
}
