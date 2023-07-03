package appointment

import (
	"fmt"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/appointment"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"
)

type IRepository interface {
	Create(appointment *domain.Appointment) error
	Read(id int) (*domain.Appointment, error)
	Update(appointment *domain.Appointment) error
	Patch(id int, dateTime string) error
	Delete(id int) error
	List() ([]*domain.Appointment, error)
	ListPatients() ([]*domain.Patient, error)
	ListDentists() ([]*domain.Dentist, error)
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) Create(appointment *domain.Appointment) error {
	err := r.Store.Create(appointment)
	if err != nil {
		return web.NewInternalServerApiError("failed to create appointment")
	}
	return nil
}

func (r *Repository) Read(id int) (*domain.Appointment, error) {
	appointment, err := r.Store.Read(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("appointment_id %d not found", id))
	}
	return appointment, nil
}

func (r *Repository) Update(appointment *domain.Appointment) error {
	err := r.Store.Update(appointment)
	if err != nil {
		return web.NewInternalServerApiError("failed to update appointment")
	}
	return nil
}

func (r *Repository) Patch(id int, dateTime string) error {
	appointment, err := r.Store.Read(id)
	if err != nil {
		return web.NewInternalServerApiError("failed to get appointment")
	}

	appointment.DateTime = dateTime

	err = r.Store.Update(appointment)
	if err != nil {
		return web.NewInternalServerApiError("failed to update appointment")
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	err := r.Store.Delete(id)
	if err != nil {
		return web.NewInternalServerApiError("failed to delete appointment")
	}
	return nil
}

func (r *Repository) List() ([]*domain.Appointment, error) {
	appointments, err := r.Store.List()
	if err != nil {
		return nil, web.NewInternalServerApiError("failed to get appointments list")
	}
	return appointments, nil
}

func (r *Repository) ListPatients() ([]*domain.Patient, error) {
	patients, err := r.Store.ListPatients()
	if err != nil {
		return nil, web.NewInternalServerApiError("failed to get patients list")
	}
	return patients, nil
}

func (r *Repository) ListDentists() ([]*domain.Dentist, error) {
	dentists, err := r.Store.ListDentists()
	if err != nil {
		return nil, web.NewInternalServerApiError("failed to get dentists list")
	}
	return dentists, nil
}
