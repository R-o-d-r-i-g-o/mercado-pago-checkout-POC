package entities

type Email struct {
	ContentType EmailMessageType
	Receivers   []string
	Message     string
	Subject     string
	Sender      string
}

func (e Email) WithReceivers(receivers ...string) Email {
	e.Receivers = receivers
	return e
}

func (e Email) WithMessage(message string) Email {
	e.Message = message
	return e
}

func (e Email) WithSubject(subject string) Email {
	e.Subject = subject
	return e
}

func (e Email) WithSender(sender string) Email {
	e.Sender = sender
	return e
}

func (e Email) WithContentType(contentType EmailMessageType) Email {
	e.ContentType = contentType
	return e
}

func (e Email) WithDefaultContentType() Email {
	e.ContentType = HTML
	return e
}

func (e Email) IsValidContentType() bool {
	var contentTypes = map[EmailMessageType]bool{
		HTML:                  true,
		PLAIN_TEXT:            true,
		MULTIPART_MIXED:       true,
		MULTIPART_ALTERNATIVE: true,
	}

	return contentTypes[e.ContentType]
}
