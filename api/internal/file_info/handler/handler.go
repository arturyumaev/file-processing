package handler

import (
	"context"
	"net/http"

	"github.com/arturyumaev/file-processing/api/internal/file_info"
	"github.com/gin-gonic/gin"
)

type handler struct {
	svc file_info.Service
}

func (h *handler) GetFileInfo(c *gin.Context) {
	filename := c.Param("name")
	ctx := context.Background()

	file, err := h.svc.GetFileInfo(ctx, filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, file)
}

func New(svc file_info.Service) file_info.Handler {
	return &handler{svc}
}
