package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
//
//	@summary		Health Check
//	@description	Health checking for the service
//	@id				HealthCheckHandler
//	@produce		json
//	@response		200	{object}	Response	"OK"
//	@router			/health [get]
func HealthCheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{Status: "It's All Good, API is Working :)"})
	}
}
