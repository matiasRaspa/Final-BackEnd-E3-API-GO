package store

import (
	"database/sql"
	"time"

	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Create(patient *domain.Patient) error {
	admissionDate := time.Now()
	formatDate := "2006-01-02"
	patient.AdmissionDate = admissionDate.Format(formatDate)
	query := "INSERT INTO patients (name, last_name, address, dni, admission_date) VALUES (?, ?, ?, ?, ?);"
	result, err := s.DB.Exec(query, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.AdmissionDate)
	if err != nil {
		return err
	}
	patient.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Read(id int) (*domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.AdmissionDate)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (s *SqlStore) Update(patient *domain.Patient) error {
	query := "UPDATE patients SET name = ?, last_name = ?, address = ?, dni = ?, admission_date = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.AdmissionDate, patient.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Patch(id int, address string) error {
	query := "UPDATE patients SET address = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, address, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Delete(id int) error {
	query := "DELETE FROM patients WHERE id = ?;"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) List() ([]*domain.Patient, error) {
	var patients []*domain.Patient
	query := "SELECT * FROM patients;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(&patient.Id, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.AdmissionDate)
		if err != nil {
			return nil, err
		}
		patients = append(patients, &patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}
