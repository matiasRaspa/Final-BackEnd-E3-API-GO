package handler

import (
	"net/http"
	"strconv"

	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/appointment"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	AppointmentService appointment.IService
}

// MakeAppointment godoc
// @Summary Create a new appointment
// @Description Create a new appointment with the provided data
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointment body domain.Appointment true "Appointment object"
// @Success 200 {object} domain.Appointment
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments [post]
func (h *AppointmentHandler) MakeAppointment(ctx *gin.Context) {
	var appointment domain.Appointment
	if err := ctx.ShouldBindJSON(&appointment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	patientList, errFound := h.AppointmentService.ListPatients()
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}
	for _, patient := range patientList {
		if patient.Id == appointment.PatientId {
			appointment.PatientId = patient.Id
			break
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	dentistList, errFound := h.AppointmentService.ListDentists()
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}
	for _, dentist := range dentistList {
		if dentist.Id == appointment.DentistId {
			appointment.DentistId = dentist.Id
			break
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}

	err := h.AppointmentService.Create(&appointment)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to create appointment"))
		return
	}

	ctx.JSON(http.StatusOK, appointment)
}

// MakeAppointmentByDniAndLicense godoc
// @Summary Create a new appointment using patient's DNI and dentist's license
// @Description Create a new appointment with the provided patient's DNI and dentist's license
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointmentQuery body domain.AppointmentRegister true "Appointment query object"
// @Success 200 {object} domain.Appointment
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments/dni-license [post]
func (h *AppointmentHandler) MakeAppointmentByDniAndLicense(ctx *gin.Context) {
	var appointmentQuery domain.AppointmentRegister
	if err := ctx.ShouldBindJSON(&appointmentQuery); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	var appointment domain.Appointment
	appointment.Description = appointmentQuery.Description

	patientList, errFound := h.AppointmentService.ListPatients()
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}
	for _, patient := range patientList {
		if patient.DNI == appointmentQuery.DniPatient {
			appointment.PatientId = patient.Id
			break
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	dentistList, errFound := h.AppointmentService.ListDentists()
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}
	for _, dentist := range dentistList {
		if dentist.License == appointmentQuery.LicenseDentist {
			appointment.DentistId = dentist.Id
			break
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}

	err := h.AppointmentService.Create(&appointment)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to create appointment"))
		return
	}

	ctx.JSON(http.StatusOK, appointment)
}

// GetAppointmentById godoc
// @Summary Get an appointment by ID
// @Description Get details of a specific appointment based on its ID
// @Tags appointments
// @Param id path int true "Appointment ID"
// @Success 200 {object} domain.Appointment
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments/{id} [get]
func (h *AppointmentHandler) GetAppointmentById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	appointmentFound, errFound := h.AppointmentService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get appointment"))
		return
	}
	ctx.JSON(http.StatusOK, appointmentFound)
}

// GetAppointmentByDni godoc
// @Summary Get an appointment by patient's DNI
// @Description Get details of a specific appointment based on the patient's DNI
// @Tags appointments
// @Param dni query string true "Patient's DNI"
// @Success 200 {object} domain.AppointmentByDni
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments/dni [get]
func (h *AppointmentHandler) GetAppointmentByDni(ctx *gin.Context) {
	dniQueryParam := ctx.Query("dni")

	var appointmentByDni domain.AppointmentByDni

	patientList, errFound := h.AppointmentService.ListPatients()
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}
	for _, patient := range patientList {
		if patient.DNI == dniQueryParam {
			appointmentByDni.Patient = *patient
			break
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	appointmentsList, err := h.AppointmentService.List()
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get appointments list"))
		return
	}
	for _, appointment := range appointmentsList {
		if appointment.PatientId == appointmentByDni.Patient.Id {
			appointmentByDni.Id = appointment.Id
			appointmentByDni.DateTime = appointment.DateTime
			appointmentByDni.Description = appointment.Description

			dentistList, errFound := h.AppointmentService.ListDentists()
			if errFound != nil {
				if apiErr, ok := errFound.(*web.ErrorApi); ok {
					ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
					return
				}
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
				return
			}
			for _, dentist := range dentistList {
				if dentist.Id == appointment.DentistId {
					appointmentByDni.Dentist = *dentist
					break
				}
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
				return
			}
			break
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	ctx.JSON(http.StatusOK, appointmentByDni)
}

// UpdateAppointmentById godoc
// @Summary Update an appointment by ID
// @Description Update a specific appointment based on its ID
// @Tags appointments
// @Param id path int true "Appointment ID"
// @Success 200 {object} domain.Appointment
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments/{id} [put]
func (h *AppointmentHandler) UpdateAppointmentById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	appointmentFound, errFound := h.AppointmentService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get appointment"))
		return
	}

	if err := ctx.ShouldBindJSON(&appointmentFound); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err = h.AppointmentService.Update(appointmentFound)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to update appointment"))
		return
	}

	ctx.JSON(http.StatusOK, appointmentFound)
}

// UpdateDateTimeById godoc
// @Summary Update the date and time of an appointment by ID
// @Description Update the date and time of a specific appointment based on its ID
// @Tags appointments
// @Param id path int true "Appointment ID"
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments/{id} [patch]
func (h *AppointmentHandler) UpdateDateTimeById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	appointmentFound, errFound := h.AppointmentService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get appointment"))
		return
	}

	var patchData struct {
		DateTime string `json:"date_time"`
	}
	if err := ctx.ShouldBindJSON(&patchData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err = h.AppointmentService.Patch(int(appointmentFound.Id), patchData.DateTime)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to patch appointment"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "appointment patched successfully"})
}

// DeleteAppointmentById godoc
// @Summary Delete an appointment by ID
// @Description Delete an appointment based on its ID
// @Tags appointments
// @Param id path int true "Appointment ID"
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /appointments/{id} [delete]
func (h *AppointmentHandler) DeleteAppointmentById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	appointmentFound, errFound := h.AppointmentService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get appointment"))
		return
	}

	err = h.AppointmentService.Delete(int(appointmentFound.Id))
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to delete appointment"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "appointment deleted successfully"})
}

// ListAppointments godoc
// @Summary Get a list of appointments
// @Description Retrieve a list of all appointments
// @Tags appointments
// @Produce json
// @Success 200 {array} domain.Appointment
// @Failure 500 {object} web.ErrorApi
// @Router /appointments [get]
func (h *AppointmentHandler) ListAppointments(ctx *gin.Context) {
	appointments, err := h.AppointmentService.List()
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get appointments list"))
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}
