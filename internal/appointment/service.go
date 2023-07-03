package appointment

import "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"

type IService interface {
	Create(appointment *domain.Appointment) error
	Read(id int) (*domain.Appointment, error)
	Update(appointment *domain.Appointment) error
	Patch(id int, dateTime string) error
	Delete(id int) error
	List() ([]*domain.Appointment, error)
	ListPatients() ([]*domain.Patient, error)
	ListDentists() ([]*domain.Dentist, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) Create(appointment *domain.Appointment) error {
	err := s.Repository.Create(appointment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Read(id int) (*domain.Appointment, error) {
	appointment, err := s.Repository.Read(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *Service) Update(appointment *domain.Appointment) error {
	err := s.Repository.Update(appointment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Patch(id int, dateTime string) error {
	err := s.Repository.Patch(id, dateTime)
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

func (s *Service) List() ([]*domain.Appointment, error) {
	appointments, err := s.Repository.List()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *Service) ListPatients() ([]*domain.Patient, error) {
	patients, err := s.Repository.ListPatients()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (s *Service) ListDentists() ([]*domain.Dentist, error) {
	dentists, err := s.Repository.ListDentists()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}
