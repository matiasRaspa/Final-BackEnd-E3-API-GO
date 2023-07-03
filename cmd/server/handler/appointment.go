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
