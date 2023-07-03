package handler

import (
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/patient"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	PatientService patient.IService
}

// MakePatient creates a new patient.
// @Summary Create a new patient
// @Description Create a new patient
// @Accept json
// @Produce json
// @Param patient body domain.Patient true "Patient object"
// @Success 200 {object} domain.Patient
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /patients [post]
func (h *PatientHandler) MakePatient(ctx *gin.Context) {
	var patient domain.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err := h.PatientService.Create(&patient)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to create patient"))
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

// GetPatientById retrieves a patient by ID.
// @Summary Get a patient by ID
// @Description Get a patient by ID
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} domain.Patient
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /patients/{id} [get]
func (h *PatientHandler) GetPatientById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	patientFound, errFound := h.PatientService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}
	ctx.JSON(http.StatusOK, patientFound)
}

// UpdatePatientById updates a patient by ID.
// @Summary Update a patient by ID
// @Description Update a patient by ID
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body domain.Patient true "Patient object"
// @Success 200 {object} domain.Patient
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /patients/{id} [put]
func (h *PatientHandler) UpdatePatientById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	patientFound, errFound := h.PatientService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	if err := ctx.ShouldBindJSON(&patientFound); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err = h.PatientService.Update(patientFound)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to update patient"))
		return
	}

	ctx.JSON(http.StatusOK, patientFound)
}

// UpdateAddressById updates the address of a patient by ID.
// @Summary Update the address of a patient by ID
// @Description Update the address of a patient by ID
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /patients/{id} [patch]
func (h *PatientHandler) UpdateAddressById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	patientFound, errFound := h.PatientService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	var patchData struct {
		Address string `json:"address"`
	}
	if err := ctx.ShouldBindJSON(&patchData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err = h.PatientService.Patch(int(patientFound.Id), patchData.Address)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to patch patient"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "patient patched successfully"})
}

// DeletePatientById deletes a patient by ID.
// @Summary Delete a patient by ID
// @Description Delete a patient by ID
// @Produce json
// @Param id path int true "Patient ID"
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /patients/{id} [delete]
func (h *PatientHandler) DeletePatientById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	patientFound, errFound := h.PatientService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patient"))
		return
	}

	err = h.PatientService.Delete(int(patientFound.Id))
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to delete patient"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "patient deleted successfully"})
}

// ListPatients retrieves the list of patients.
// @Summary Get the list of patients
// @Description Get the list of patients
// @Produce json
// @Success 200 {array} domain.Patient
// @Failure 500 {object} web.ErrorApi
// @Router /patients [get]
func (h *PatientHandler) ListPatients(ctx *gin.Context) {
	patients, err := h.PatientService.List()
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get patients list"))
		return
	}
	ctx.JSON(http.StatusOK, patients)
}
