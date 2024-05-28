package newsletter

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

func registerReceipients(nlID int64, nl Newsletter, db *Database) error {
	var emails []string

	if nl.Recipient != "" {
		emails = append(emails, strings.TrimSpace(nl.Recipient))
	}

	if nl.RecipientsFileName != "" {
		fileExtension := filepath.Ext(nl.RecipientsFileName)

		if fileExtension != ".csv" && fileExtension != ".txt" {
			return errors.New("invalid file format. Please provide a .csv or .txt file")
		}

		if fileExtension == ".csv" {
			b, err := os.Open("/tmp/" + nl.RecipientsFileName)
			if err != nil {
				return errors.New("failed to open the file")
			}
			csvReader := csv.NewReader(b)
			records, err := csvReader.ReadAll()
			if err != nil {
				return errors.New("unable to parse file as CSV. Make sure it is an email per row")
			}

			for _, record := range records {
				emails = append(emails, record[0])
			}

		} else if fileExtension == ".txt" {
			b, err := os.ReadFile("/tmp/" + nl.RecipientsFileName)

			if err != nil {
				return errors.New("failed to open the file")
			}

			emails = strings.Split(string(b), ",")
		}
	}

	valueStrings := make([]string, 0, len(emails))
	valueArgs := make([]interface{}, 0, len(emails)*2)

	for _, email := range emails {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, strings.TrimSpace(email), nlID)
	}

	stmt := fmt.Sprintf("INSERT INTO RECEIPIENT (EMAIL, NEWSLETTER_ID) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.driver.Exec(stmt, valueArgs...)

	if err != nil {
		return errors.New("failed to register recipients")
	}

	return nil
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

	tx, err := db.driver.BeginTx(context.Background(), nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create the newsletter: %s", err)
		return
	}

	query := `INSERT INTO NEWSLETTER
              (NAME, DESCRIPTION, SUBJECT, CONTENT_ATTACHMENT_PATH)
            VALUES (?, ?, ?, ?)`
	queryResult, err := db.driver.Exec(query, newsletter.Name, newsletter.Description, newsletter.Subject, newsletter.ContentFileName)

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create the newsletter: %s", err)
		tx.Rollback()
		return
	}

	id, err := queryResult.LastInsertId()

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get the newsletter ID: %s", err)
		tx.Rollback()
		return
	}

	err = registerReceipients(id, newsletter, db)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		c.String(http.StatusInternalServerError, "Failed to commit transaction", err)
		tx.Rollback()
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
