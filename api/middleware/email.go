package middleware

import (
	"bytes"
	"code-space-backend-api/app/Mail/entities"
	"code-space-backend-api/app/Mail/interfaces/dto"
	"code-space-backend-api/app/Mail/usecases"
	"code-space-backend-api/env"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type credentials struct {
	Email    string `json:"email"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type emailMiddleware struct {
	sendEmailWithTemplate usecases.SendEmailWithTemplate
}

func NewEmailMiddleware(sendEmailWithTemplate usecases.SendEmailWithTemplate) Middleware {
	return &emailMiddleware{
		sendEmailWithTemplate: sendEmailWithTemplate,
	}
}

func (e *emailMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		creds, err := e.saveRequestBodyAndParseToUserCredentials(c)
		if err != nil {
			c.Next()
			return
		}

		c.Next()

		var emailMessage = e.createEmailMessageWithEnvironmentCredentials(creds)
		if err := e.sendEmailWithTemplate.Execute(emailMessage); err != nil {
			c.Abort()
			return
		}
	}
}

func (e *emailMiddleware) saveRequestBodyAndParseToUserCredentials(c *gin.Context) (credentials, error) {
	var creds credentials

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return creds, err
	}

	if err := json.Unmarshal(payload, &creds); err != nil {
		return creds, err
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
	return creds, nil
}

func (e *emailMiddleware) createEmailMessageWithEnvironmentCredentials(userCredentials credentials) dto.LoginEmailTemplateInputDTO {
	var (
		receivers      = []string{userCredentials.Email}
		emailType      = string(entities.HTML)
		subjectMessage = fmt.Sprintf("Welcome to your Code Academy, %s", userCredentials.Username)
	)

	return dto.LoginEmailTemplateInputDTO{
		Receivers:   receivers,
		ContentType: emailType,
		Subject:     subjectMessage,
		Sender:      env.EmailService.SenderAddress,
		Login:       userCredentials.Email,
		Password:    userCredentials.Password,
	}
}
