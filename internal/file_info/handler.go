package file_info

import "github.com/gin-gonic/gin"

type Handler interface {
	GetFileInfo(c *gin.Context)
}
