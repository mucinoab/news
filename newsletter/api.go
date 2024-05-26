package newsletter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Newsletter struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`

	Subject         string `json:"subject" binding:"required"`
	ContentFileName string `json:"content" binding:"required"`

	RecipientsFileName string `json:"recipients"`
	Recipient          string `json:"recipient"`
}

func Create(c *gin.Context) {
	var newsletter Newsletter

	if err := c.ShouldBindJSON(&newsletter); err != nil {
		c.String(http.StatusBadRequest, "Failed to deserialize the newsletter: %s", err)
		return
	}

	// TODO: Save the newsletter to the database

	c.String(http.StatusCreated, "Newsletter created successfully")
}
