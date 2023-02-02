package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/gin-gonic/gin"
)

func (h *v1Handler) initRecordRoutes(recordGroup *gin.RouterGroup) {
	recordGroup.POST("", h.createRecord)
	recordGroup.GET("/:id", h.getRecordById)
	recordGroup.GET("", h.getAllRecords)
	recordGroup.GET("/topics", h.getAllTopics)
	recordGroup.GET("/search", h.searchRecordsByTitle)
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
		h.newHttpErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid id param - %w", err))
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

func (h *v1Handler) getAllRecords(c *gin.Context) {
	topic := c.DefaultQuery("topic", "")
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, errors.New("query param error - 'offset' must be integer"))
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, errors.New("query param error - 'limit' must be integer"))
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to get user id - %w", err))
		return
	}

	recordDTOs, err := h.usecases.Record.GetAllRecords(c.Request.Context(), userId, topic, offset, limit)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("get all records error - %w", err))
		return
	}

	h.newHttpSuccessResponse(c, recordDTOs)
}

func (h *v1Handler) getAllTopics(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to get user id - %w", err))
		return
	}

	topics, err := h.usecases.GetAllTopics(c, userId)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("get all topics error - %w", err))
		return
	}

	h.newHttpSuccessResponse(c, topics)
}

func (h *v1Handler) searchRecordsByTitle(c *gin.Context) {
	title, exists := c.GetQuery("title")
	if !exists || title == "" {
		h.newHttpErrorResponse(c, http.StatusBadRequest, errors.New("query param error - missing 'title'"))
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, errors.New("query param error - 'offset' must be integer"))
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, errors.New("query param error - 'limit' must be integer"))
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to get user id - %w", err))
		return
	}

	recordDTOs, err := h.usecases.Record.SearchRecordsByTitle(c.Request.Context(), userId, title, offset, limit)
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("search records by title error - %w", err))
		return
	}

	h.newHttpSuccessResponse(c, recordDTOs)
}
