package v1

import (
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/gin-gonic/gin"
)

type recordHandler struct {
	group         *gin.RouterGroup
	recordUsecase usecases.Record
}

func NewRecordHandler(group *gin.RouterGroup, recordUsecase usecases.Record) *recordHandler {
	return &recordHandler{
		group:         group,
		recordUsecase: recordUsecase,
	}
}

func (a *recordHandler) InitRoutes() {

}
