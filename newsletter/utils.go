package newsletter

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wneessen/go-mail"
)

func getHostUrl(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.SaveUploadedFile(file, "/tmp/"+file.Filename)
	c.String(http.StatusOK, "OK")
}

func BuildMessages(newsletter Newsletter, emails []string, host string) ([]*mail.Msg, error) {
	messages := make([]*mail.Msg, 0, len(emails))

	for _, email := range emails {
		m := mail.NewMsg()
		m.Subject(newsletter.Subject)

		username := os.Getenv("USERNAME")
		if username == "" {
			return nil, errors.New("missing sender email(username)")
		}

		err := m.From(username)
		if err != nil {
			return nil, err
		}

		err = m.To(email)

		if err != nil {
			return nil, fmt.Errorf("failed to add receipient email (%s): %s", email, err)
		}

		body := fmt.Sprintf("<a href='%s/newsletter/%d/unsubscribe/%s'>Unsubscribe</a>", host, newsletter.Id, url.QueryEscape(email))
		m.SetBodyString(mail.TypeTextHTML, body)
		m.AttachFile("/tmp/" + newsletter.ContentFileName)

		messages = append(messages, m)
	}

	return messages, nil
}

func SendEmails(newsletter Newsletter, emails []string, host string) error {
	username := os.Getenv("USERNAME")
	if username == "" {
		return errors.New("missing username for mail server")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		return errors.New("missing password for mail server")
	}

	client, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(username), mail.WithPassword(password))
	if err != nil {
		return err
	}

	messaes, err := BuildMessages(newsletter, emails, host)
	if err != nil {
		return err
	}

	if err := client.DialAndSend(messaes...); err != nil {
		return fmt.Errorf("failed send mail: %s", err)
	}

	return nil
}
