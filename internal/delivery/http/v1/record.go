package v1

import (
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/gin-gonic/gin"
)

type recordHandler struct {
	log           logger.Logger
	group         *gin.RouterGroup
	recordUsecase usecases.Record
}

func NewRecordHandler(log logger.Logger, group *gin.RouterGroup, recordUsecase usecases.Record) *recordHandler {
	return &recordHandler{
		log:           log,
		group:         group,
		recordUsecase: recordUsecase,
	}
}

func (a *recordHandler) InitRoutes() {

}
