package handler

import (
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/dentist"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/internal/domain"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	DentistService dentist.IService
}

// MakeDentist creates a new dentist.
// @Summary Create a new dentist
// @Description Create a new dentist
// @Accept json
// @Produce json
// @Param dentist body domain.Dentist true "Dentist object"
// @Success 200 {object} domain.Dentist
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /dentists [post]
func (h *DentistHandler) MakeDentist(ctx *gin.Context) {
	var dentist domain.Dentist
	if err := ctx.ShouldBindJSON(&dentist); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err := h.DentistService.Create(&dentist)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to create dentist"))
		return
	}

	ctx.JSON(http.StatusOK, dentist)
}

// GetDentistById retrieves a dentist by ID.
// @Summary Get a dentist by ID
// @Description Get a dentist by ID
// @Produce json
// @Param id path int true "Dentist ID"
// @Success 200 {object} domain.Dentist
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /dentists/{id} [get]
func (h *DentistHandler) GetDentistById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	dentistFound, errFound := h.DentistService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}
	ctx.JSON(http.StatusOK, dentistFound)
}

// UpdateDentistById updates a dentist by ID.
// @Summary Update a dentist by ID
// @Description Update a dentist by ID
// @Accept json
// @Produce json
// @Param id path int true "Dentist ID"
// @Param dentist body domain.Dentist true "Dentist object"
// @Success 200 {object} domain.Dentist
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /dentists/{id} [put]
func (h *DentistHandler) UpdateDentistById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	dentistFound, errFound := h.DentistService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}

	if err := ctx.ShouldBindJSON(&dentistFound); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err = h.DentistService.Update(dentistFound)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to update dentist"))
		return
	}

	ctx.JSON(http.StatusOK, dentistFound)
}

// UpdateLicenseById updates the license of a dentist by ID.
// @Summary Update the license of a dentist by ID
// @Description Update the license of a dentist by ID
// @Accept json
// @Produce json
// @Param id path int true "Dentist ID"
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /dentists/{id} [patch]
func (h *DentistHandler) UpdateLicenseById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	dentistFound, errFound := h.DentistService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}

	var patchData struct {
		License string `json:"license"`
	}
	if err := ctx.ShouldBindJSON(&patchData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	err = h.DentistService.Patch(int(dentistFound.Id), patchData.License)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to patch dentist"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dentist patched successfully"})
}

// DeleteDentistById deletes a dentist by ID.
// @Summary Delete a dentist by ID
// @Description Delete a dentist by ID
// @Produce json
// @Param id path int true "Dentist ID"
// @Failure 400 {object} web.ErrorApi
// @Failure 500 {object} web.ErrorApi
// @Router /dentists/{id} [delete]
func (h *DentistHandler) DeleteDentistById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	dentistFound, errFound := h.DentistService.Read(id)
	if errFound != nil {
		if apiErr, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentist"))
		return
	}

	err = h.DentistService.Delete(int(dentistFound.Id))
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to delete dentist"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dentist deleted successfully"})
}

// ListDentists retrieves a list of dentists.
// @Summary Get a list of dentists
// @Description Get a list of dentists
// @Produce json
// @Success 200 {array} domain.Dentist
// @Failure 500 {object} web.ErrorApi
// @Router /dentists [get]
func (h *DentistHandler) ListDentists(ctx *gin.Context) {
	dentists, err := h.DentistService.List()
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError("failed to get dentists list"))
		return
	}
	ctx.JSON(http.StatusOK, dentists)
}
