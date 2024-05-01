package entities

type EmailMessageType string

const (
	HTML                  EmailMessageType = "text/html"
	PLAIN_TEXT            EmailMessageType = "text/plain"
	MULTIPART_MIXED       EmailMessageType = "multipart/mixed"
	MULTIPART_ALTERNATIVE EmailMessageType = "multipart/alternative"
)
