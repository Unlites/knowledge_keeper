package v1

import (
	"net/http"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/gin-gonic/gin"
)

func (h *v1Handler) initRecordRoutes(recordGroup *gin.RouterGroup) {
	recordGroup.POST("/", h.createRecord)
}

func (h *v1Handler) createRecord(c *gin.Context) {
	var recordDTO *dto.RecordDTORequest
	if err := c.BindJSON(&recordDTO); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := recordDTO.Validate(); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, &errs.ErrValidation{Message: err.Error()})
		return
	}

	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if err := h.usecases.Record.CreateRecord(c.Request.Context(), userId, recordDTO); err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}
