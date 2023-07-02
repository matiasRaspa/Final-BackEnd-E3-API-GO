package patient

import (
	"fmt"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/patient"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"
)

type IRepository interface {
	Create(patient *domain.Patient) error
	Read(id int) (*domain.Patient, error)
	Update(patient *domain.Patient) error
	Patch(id int, address string) error
	Delete(id int) error
	List() ([]*domain.Patient, error)
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) Create(patient *domain.Patient) error {
	err := r.Store.Create(patient)
	if err != nil {
		return web.NewInternalServerApiError("failed to create patient")
	}
	return nil
}

func (r *Repository) Read(id int) (*domain.Patient, error) {
	patient, err := r.Store.Read(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("patient_id %d not found", id))
	}
	return patient, nil
}

func (r *Repository) Update(patient *domain.Patient) error {
	err := r.Store.Update(patient)
	if err != nil {
		return web.NewInternalServerApiError("failed to update patient")
	}
	return nil
}

func (r *Repository) Patch(id int, address string) error {
	patient, err := r.Store.Read(id)
	if err != nil {
		return web.NewInternalServerApiError("failed to get patient")
	}

	patient.Address = address

	err = r.Store.Update(patient)
	if err != nil {
		return web.NewInternalServerApiError("failed to update patient")
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	err := r.Store.Delete(id)
	if err != nil {
		return web.NewInternalServerApiError("failed to delete patient")
	}
	return nil
}

func (r *Repository) List() ([]*domain.Patient, error) {
	patients, err := r.Store.List()
	if err != nil {
		return nil, web.NewInternalServerApiError("failed to get patients list")
	}
	return patients, nil
}
