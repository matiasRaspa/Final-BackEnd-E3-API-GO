package dentist

import (
	"fmt"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/dentist"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"
)

type IRepository interface {
	Create(dentist *domain.Dentist) error
	Read(id int) (*domain.Dentist, error)
	Update(dentist *domain.Dentist) error
	Patch(id int, license string) error
	Delete(id int) error
	List() ([]*domain.Dentist, error)
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) Create(dentist *domain.Dentist) error {
	err := r.Store.Create(dentist)
	if err != nil {
		return web.NewInternalServerApiError("failed to create dentist")
	}
	return nil
}

func (r *Repository) Read(id int) (*domain.Dentist, error) {
	dentist, err := r.Store.Read(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentist_id %d not found", id))
	}
	return dentist, nil
}

func (r *Repository) Update(dentist *domain.Dentist) error {
	err := r.Store.Update(dentist)
	if err != nil {
		return web.NewInternalServerApiError("failed to update dentist")
	}
	return nil
}

func (r *Repository) Patch(id int, license string) error {
	dentist, err := r.Store.Read(id)
	if err != nil {
		return web.NewInternalServerApiError("failed to get dentist")
	}

	dentist.License = license

	err = r.Store.Update(dentist)
	if err != nil {
		return web.NewInternalServerApiError("failed to update dentist")
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	err := r.Store.Delete(id)
	if err != nil {
		return web.NewInternalServerApiError("failed to delete dentist")
	}
	return nil
}

func (r *Repository) List() ([]*domain.Dentist, error) {
	dentists, err := r.Store.List()
	if err != nil {
		return nil, web.NewInternalServerApiError("failed to get dentists list")
	}
	return dentists, nil
}
