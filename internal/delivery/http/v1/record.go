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
	recordGroup.GET("/subtopics", h.getAllSubtopicsByTopic)
	recordGroup.PUT("/:id", h.updateRecord)
	recordGroup.DELETE("/:id", h.deleteRecord)
}

func (h *v1Handler) createRecord(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	var recordDTO *dto.RecordDTORequest
	if err := c.BindJSON(&recordDTO); err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("failed to bind JSON - %w", err),
		)
		return
	}

	if err := recordDTO.Validate(); err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("validation error - %w", err),
		)
		return
	}

	if err := h.usecases.Record.CreateRecord(
		c.Request.Context(),
		userId,
		recordDTO,
	); err != nil {

		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("create record error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}

func (h *v1Handler) getRecordById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	id, err := h.getIdParam(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("invalid id param - %w", err),
		)
	}

	recordDTO, err := h.usecases.Record.GetRecordById(
		c.Request.Context(),
		userId,
		id,
	)
	if err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) {
			status = http.StatusNotFound
		}

		h.newHttpErrorResponse(
			c,
			status,
			fmt.Errorf("get record by id error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, recordDTO)
}

func (h *v1Handler) getAllRecords(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	topic := c.DefaultQuery("topic", "")
	subtopic := c.DefaultQuery("subtopic", "")
	title := c.DefaultQuery("title", "")
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			errors.New("query param error - 'offset' must be integer"),
		)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			errors.New("query param error - 'limit' must be integer"),
		)
		return
	}

	recordDTOs, err := h.usecases.Record.GetAllRecords(
		c.Request.Context(),
		userId,
		topic,
		subtopic,
		title,
		offset,
		limit,
	)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("get all records error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, recordDTOs)
}

func (h *v1Handler) getAllTopics(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	topics, err := h.usecases.GetAllTopics(c, userId)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("get all topics error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, topics)
}

func (h *v1Handler) getAllSubtopicsByTopic(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	topic := c.Query("topic")
	if topic == "" {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			errors.New("query param topic is required"),
		)
		return
	}

	subtopics, err := h.usecases.GetAllSubtopicsByTopic(c, userId, topic)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("get all subtopics error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, subtopics)
}

func (h *v1Handler) updateRecord(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	id, err := h.getIdParam(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("invalid id param - %w", err),
		)
	}

	var recordDTO *dto.RecordDTORequest
	if err := c.BindJSON(&recordDTO); err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("failed to bind JSON - %w", err),
		)
		return
	}

	if err := recordDTO.Validate(); err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("validation error - %w", err),
		)
		return
	}

	if err := h.usecases.UpdateRecord(c, userId, id, recordDTO); err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) {
			status = http.StatusNotFound
		}

		h.newHttpErrorResponse(
			c,
			status,
			fmt.Errorf("update record error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}

func (h *v1Handler) deleteRecord(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("failed to get user id - %w", err),
		)
		return
	}

	id, err := h.getIdParam(c)
	if err != nil {
		h.newHttpErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Errorf("invalid id param - %w", err),
		)
	}
	if err := h.usecases.DeleteRecord(c, userId, id); err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) {
			status = http.StatusNotFound
		}

		h.newHttpErrorResponse(
			c,
			status,
			fmt.Errorf("delete record error - %w", err),
		)
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}
