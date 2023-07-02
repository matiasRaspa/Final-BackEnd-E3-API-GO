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
