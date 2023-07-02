package main

import (
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/cmd/server/handler"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/dentist"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/store"
	"database/sql"

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

	storage := store.SqlStore{DB: db}
	repo := dentist.Repository{Store: &storage}
	service := dentist.Service{Repository: &repo}
	dentistHandler := handler.DentistHandler{DentistService: &service}

	r := gin.Default()
	dentists := r.Group("/dentists")
	{
		dentists.GET("", dentistHandler.ListDentists)
		dentists.GET(":id", dentistHandler.GetDentistById)
		dentists.POST("", dentistHandler.MakeDentist)
		dentists.PUT(":id", dentistHandler.UpdateDentistById)
		dentists.PATCH(":id", dentistHandler.UpdateLicenseById)
		dentists.DELETE(":id", dentistHandler.DeleteDentistById)
	}

	r.Run(":8080")
}
