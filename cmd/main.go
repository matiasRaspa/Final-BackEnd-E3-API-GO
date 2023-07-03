package main

import (
	"database/sql"
	"net/http"

	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/cmd/server/handler"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/appointment"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/dentist"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/patient"
	appointmentStore "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/appointment"
	dentistStore "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/dentist"
	patientStore "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/patient"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	connectionString := "root:root@tcp(localhost:3306)/dental_clinic_db"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing)
	}

	storageDentist := dentistStore.SqlStore{DB: db}
	repoDentist := dentist.Repository{Store: &storageDentist}
	serviceDentist := dentist.Service{Repository: &repoDentist}
	dentistHandler := handler.DentistHandler{DentistService: &serviceDentist}

	storagePatient := patientStore.SqlStore{DB: db}
	repoPatient := patient.Repository{Store: &storagePatient}
	servicePatient := patient.Service{Repository: &repoPatient}
	patientHandler := handler.PatientHandler{PatientService: &servicePatient}

	storageAppointment := appointmentStore.SqlStore{DB: db}
	repoAppointment := appointment.Repository{Store: &storageAppointment}
	serviceAppointment := appointment.Service{Repository: &repoAppointment}
	appointmentHandler := handler.AppointmentHandler{AppointmentService: &serviceAppointment}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome API!")
	})

	dentists := r.Group("/dentists")
	{
		dentists.GET("", dentistHandler.ListDentists)
		dentists.GET(":id", dentistHandler.GetDentistById)
		dentists.POST("", dentistHandler.MakeDentist)
		dentists.PUT(":id", dentistHandler.UpdateDentistById)
		dentists.PATCH(":id", dentistHandler.UpdateLicenseById)
		dentists.DELETE(":id", dentistHandler.DeleteDentistById)
	}

	patients := r.Group("/patients")
	{
		patients.GET("", patientHandler.ListPatients)
		patients.GET(":id", patientHandler.GetPatientById)
		patients.POST("", patientHandler.MakePatient)
		patients.PUT(":id", patientHandler.UpdatePatientById)
		patients.PATCH(":id", patientHandler.UpdateAddressById)
		patients.DELETE(":id", patientHandler.DeletePatientById)
	}

	appointments := r.Group("/appointments")
	{
		appointments.GET("", appointmentHandler.ListAppointments)
		appointments.GET(":id", appointmentHandler.GetAppointmentById)
		appointments.GET("/patient", appointmentHandler.GetAppointmentByDni)
		appointments.POST("", appointmentHandler.MakeAppointment)
		appointments.POST("/register", appointmentHandler.MakeAppointmentByDniAndLicense)
		appointments.PUT(":id", appointmentHandler.UpdateAppointmentById)
		appointments.PATCH(":id", appointmentHandler.UpdateDateTimeById)
		appointments.DELETE(":id", appointmentHandler.DeleteAppointmentById)
	}

	r.Run(":8080")
}
