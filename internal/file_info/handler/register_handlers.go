package handler

import "github.com/gin-gonic/gin"

func RegisterHandlers(router *gin.Engine, service Service) {
	h := New(service)

	userEndpoints := router.Group("/files")

	{
		userEndpoints.GET("/:name", h.GetFileInfo)
	}
}
