package domain

// Patient represents a patient.
type Patient struct {
	Id            int64  `json:"id"`
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	DNI           string `json:"dni" binding:"required"`
	AdmissionDate string `json:"admission_date"`
}
