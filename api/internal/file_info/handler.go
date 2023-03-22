package file_info

import "github.com/gin-gonic/gin"

type Handler interface {
	GetFilename(c *gin.Context)
}
