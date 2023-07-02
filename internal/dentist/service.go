package dentist

import "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"

type IService interface {
	Create(dentist *domain.Dentist) error
	Read(id int) (*domain.Dentist, error)
	Update(dentist *domain.Dentist) error
	Patch(id int, license string) error
	Delete(id int) error
	List() ([]*domain.Dentist, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) Create(dentist *domain.Dentist) error {
	err := s.Repository.Create(dentist)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Read(id int) (*domain.Dentist, error) {
	dentist, err := s.Repository.Read(id)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

func (s *Service) Update(dentist *domain.Dentist) error {
	err := s.Repository.Update(dentist)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Patch(id int, license string) error {
	err := s.Repository.Patch(id, license)
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

func (s *Service) List() ([]*domain.Dentist, error) {
	dentists, err := s.Repository.List()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}
