package healthcheck

import (
	"event-system-backend/internal/db"
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHealthCheckRoutes(r *gin.RouterGroup) {
	healthGroup := r.Group("/health")
	healthGroup.GET("", HealthCheck)
}

func HealthCheck(c *gin.Context) {
	sqlDB, err := db.GetDB().DB()
	if err != nil {
		c.Error(handler.NewError(http.StatusInternalServerError, "Failed to get DB connection"))
	}

	if err := sqlDB.Ping(); err != nil {
		c.Error(handler.NewError(http.StatusInternalServerError, "Database not reachable"))
	}
	c.JSON(http.StatusOK, dto.BaseResponse(true, "Service is healthy", nil))
}
