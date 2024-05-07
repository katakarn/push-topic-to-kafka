package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VersionHandler godoc
//
//	@summary		Version
//	@description	Version for the service
//	@id				VersionHandler
//	@produce		json
//	@response		200	{object}	Response	"OK"
//	@router			/version [get]
func VersionHandler(version, buildTime, build string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{Version: version, BuildTime: buildTime, Build: build})
	}
}
