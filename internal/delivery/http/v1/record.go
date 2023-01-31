package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/gin-gonic/gin"
)

func (h *v1Handler) initRecordRoutes(recordGroup *gin.RouterGroup) {
	recordGroup.POST("/", h.createRecord)
	recordGroup.GET("/:id", h.getRecordById)
}

func (h *v1Handler) createRecord(c *gin.Context) {
	var recordDTO *dto.RecordDTORequest
	if err := c.BindJSON(&recordDTO); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, fmt.Errorf("failed to bind JSON - %w", err))
		return
	}

	if err := recordDTO.Validate(); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, fmt.Errorf("validation error - %w", err))
		return
	}

	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to get user id - %w", err))
		return
	}

	if err := h.usecases.Record.CreateRecord(c.Request.Context(), userId, recordDTO); err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("create record error - %w", err))
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}

func (h *v1Handler) getRecordById(c *gin.Context) {
	id, err := h.getIdParam(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to get id from params - %w", err))
	}
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to get user id - %w", err))
		return
	}

	recordDTO, err := h.usecases.Record.GetRecordById(c.Request.Context(), userId, id)
	if err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) {
			status = http.StatusNotFound
		}

		h.newHttpErrorResponse(c, status, fmt.Errorf("get record by id error - %w", err))
		return
	}

	h.newHttpSuccessResponse(c, recordDTO)
}
