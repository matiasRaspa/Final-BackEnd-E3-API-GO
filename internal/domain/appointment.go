package domain

type Appointment struct {
	Id          int64  `json:"id"`
	PatientId   int64  `json:"patient_id" binding:"required"`
	DentistId   int64  `json:"dentist_id" binding:"required"`
	DateTime    string `json:"date_time"`
	Description string `json:"description"`
}

type AppointmentRegister struct {
	DniPatient     string `json:"dni_patient" binding:"required"`
	LicenseDentist string `json:"license_dentist" binding:"required"`
	Description    string `json:"description"`
}

type AppointmentByDni struct {
	Id          int64   `json:"id"`
	Patient     Patient `json:"patient"`
	Dentist     Dentist `json:"dentist"`
	DateTime    string  `json:"date_time"`
	Description string  `json:"description"`
}
