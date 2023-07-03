package store

import (
	"database/sql"
	"time"

	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Create(appointment *domain.Appointment) error {

	dateTime := time.Now()
	formatDate := "2006-01-02 15:04:05"
	appointment.DateTime = dateTime.Format(formatDate)
	query := "INSERT INTO appointments (patient_id, dentist_id, date_time, description) VALUES (?, ?, ?, ?);"
	result, err := s.DB.Exec(query, appointment.PatientId, appointment.DentistId, appointment.DateTime, appointment.Description)
	if err != nil {
		return err
	}
	appointment.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Read(id int) (*domain.Appointment, error) {
	var appointment domain.Appointment
	query := "SELECT * FROM appointments WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&appointment.Id, &appointment.PatientId, &appointment.DentistId, &appointment.DateTime, &appointment.Description)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (s *SqlStore) Update(appointment *domain.Appointment) error {
	query := "UPDATE appointments SET patient_id = ?, dentist_id = ?, date_time = ?, description = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, appointment.PatientId, appointment.DentistId, appointment.DateTime, appointment.Description, appointment.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Patch(id int, dateTime string) error {
	query := "UPDATE appointments SET date_time = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, dateTime, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Delete(id int) error {
	query := "DELETE FROM appointments WHERE id = ?;"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) List() ([]*domain.Appointment, error) {
	var appointments []*domain.Appointment
	query := "SELECT * FROM appointments;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.Id, &appointment.PatientId, &appointment.DentistId, &appointment.DateTime, &appointment.Description)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, &appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (s *SqlStore) ListPatients() ([]*domain.Patient, error) {
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

func (s *SqlStore) ListDentists() ([]*domain.Dentist, error) {
	var dentists []*domain.Dentist
	query := "SELECT * FROM dentists;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(&dentist.Id, &dentist.LastName, &dentist.Name, &dentist.License)
		if err != nil {
			return nil, err
		}
		dentists = append(dentists, &dentist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dentists, nil
}
