package newsletter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.SaveUploadedFile(file, "/tmp/"+file.Filename)
	c.String(http.StatusOK, "OK")
}
