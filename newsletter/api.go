package newsletter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Newsletter struct {
	Id          int    `json:"id"`
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

	db, err := NewDatabase()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create the newsletter: %s", err)
		return
	}
	defer db.Close()

	query := `INSERT INTO NEWSLETTER
              (NAME, DESCRIPTION, SUBJECT, CONTENT_ATTACHMENT_PATH )
            VALUES (?, ?, ?, ?)`
	_, err = db.driver.Exec(query, newsletter.Name, newsletter.Description, newsletter.Subject, newsletter.ContentFileName)

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create the newsletter: %s", err)
		return
	}

	c.String(http.StatusCreated, "Newsletter created successfully")
}

func GetAll(c *gin.Context) {
	db, err := NewDatabase()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get the newsletters: %s", err)
		return
	}
	defer db.Close()

	result, err := db.driver.Query("SELECT * FROM NEWSLETTER")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get the newsletters: %s", err)
		return
	}

	newsletters := []Newsletter{}

	for result.Next() {
		var newsletter Newsletter
		err = result.Scan(&newsletter.Id, &newsletter.Name, &newsletter.Description, &newsletter.Subject, &newsletter.ContentFileName)

		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get the newsletters: %s", err)
			return
		}

		newsletters = append(newsletters, newsletter)
	}

	c.JSON(http.StatusOK, newsletters)
}

func Send(c *gin.Context) {
	c.JSON(http.StatusOK, "Newsletter sent successfully")
}
