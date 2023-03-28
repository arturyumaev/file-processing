package handler

import (
	"net/http"

	"github.com/arturyumaev/file-processing/api/internal/file_info"
	"github.com/gin-gonic/gin"
)

type handler struct {
	svc file_info.Service
}

func (h *handler) GetFileInfo(c *gin.Context) {
	filename := c.Param("name")

	file, err := h.svc.GetFileInfo(c.Request.Context(), filename)
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
