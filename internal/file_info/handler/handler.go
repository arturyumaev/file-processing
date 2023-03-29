package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arturyumaev/file-processing/internal/file_info"
)

type handler struct {
	svc file_info.Service
}

// GetFileInfo godoc
// @Summary      Get meta information about a file
// @Description  get file info by its name
// @Tags         files
// @Produce      json
// @Param        name path string false "File name"
// @Success      200  {object}  models.FileInfo
// @Failure      400  {object}  file_info.HttpResponseErr
// @Failure      404  {object}  file_info.HttpResponseErr
// @Failure      408  {object}  file_info.HttpResponseErr
// @Failure      500  {object}  file_info.HttpResponseErr
// @Router       /files/{name} [get]
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
