package patient

import "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"

type IService interface {
	Create(patient *domain.Patient) error
	Read(id int) (*domain.Patient, error)
	Update(patient *domain.Patient) error
	Patch(id int, address string) error
	Delete(id int) error
	List() ([]*domain.Patient, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) Create(patient *domain.Patient) error {
	err := s.Repository.Create(patient)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Read(id int) (*domain.Patient, error) {
	patient, err := s.Repository.Read(id)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (s *Service) Update(patient *domain.Patient) error {
	err := s.Repository.Update(patient)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Patch(id int, address string) error {
	err := s.Repository.Patch(id, address)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) List() ([]*domain.Patient, error) {
	patients, err := s.Repository.List()
	if err != nil {
		return nil, err
	}
	return patients, nil
}
