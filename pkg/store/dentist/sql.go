package store

import (
	"database/sql"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Create(dentist *domain.Dentist) error {
	query := "INSERT INTO dentists (last_name, name, license) VALUES (?, ?, ?);"
	result, err := s.DB.Exec(query, dentist.LastName, dentist.Name, dentist.License)
	if err != nil {
		return err
	}
	dentist.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Read(id int) (*domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentist.Id, &dentist.LastName, &dentist.Name, &dentist.License)
	if err != nil {
		return nil, err
	}
	return &dentist, nil
}

func (s *SqlStore) Update(dentist *domain.Dentist) error {
	query := "UPDATE dentists SET last_name = ?, name = ?, license = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, dentist.LastName, dentist.Name, dentist.License, dentist.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Patch(id int, license string) error {
	query := "UPDATE dentists SET license = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, license, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Delete(id int) error {
	query := "DELETE FROM dentists WHERE id = ?;"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) List() ([]*domain.Dentist, error) {
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
