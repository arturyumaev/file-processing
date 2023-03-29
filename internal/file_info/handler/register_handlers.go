package handler

import (
	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, service file_info.Service) {
	h := New(service)

	userEndpoints := router.Group("/files")

	{
		userEndpoints.GET("/:name", h.GetFileInfo)
	}
}
