package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/cmd/server/handler"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/appointment"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/dentist"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/patient"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/middleware"
	appointmentStore "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/appointment"
	dentistStore "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/dentist"
	patientStore "github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store/patient"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := godotenv.Load("./cmd/.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	password := os.Getenv("TOKEN")
	fmt.Println(password)

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
		dentists.POST("", middleware.Authentication(), dentistHandler.MakeDentist)
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.UpdateDentistById)
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.UpdateLicenseById)
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.DeleteDentistById)
	}

	patients := r.Group("/patients")
	{
		patients.GET("", patientHandler.ListPatients)
		patients.GET(":id", patientHandler.GetPatientById)
		patients.POST("", middleware.Authentication(), patientHandler.MakePatient)
		patients.PUT(":id", middleware.Authentication(), patientHandler.UpdatePatientById)
		patients.PATCH(":id", middleware.Authentication(), patientHandler.UpdateAddressById)
		patients.DELETE(":id", middleware.Authentication(), patientHandler.DeletePatientById)
	}

	appointments := r.Group("/appointments")
	{
		appointments.GET("", appointmentHandler.ListAppointments)
		appointments.GET(":id", appointmentHandler.GetAppointmentById)
		appointments.GET("/patient", appointmentHandler.GetAppointmentByDni)
		appointments.POST("", middleware.Authentication(), appointmentHandler.MakeAppointment)
		appointments.POST("/register", middleware.Authentication(), appointmentHandler.MakeAppointmentByDniAndLicense)
		appointments.PUT(":id", middleware.Authentication(), appointmentHandler.UpdateAppointmentById)
		appointments.PATCH(":id", middleware.Authentication(), appointmentHandler.UpdateDateTimeById)
		appointments.DELETE(":id", middleware.Authentication(), appointmentHandler.DeleteAppointmentById)
	}

	r.Run(":8080")
}
