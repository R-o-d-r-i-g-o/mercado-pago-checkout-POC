package errors

import "net/http"

var statusCodeMap = map[ErrorType]int{
	TypeBusinessRule:  http.StatusUnprocessableEntity,
	TypeValidation:    http.StatusBadRequest,
	TypeNotFound:      http.StatusNotFound,
	TypeUnathorizated: http.StatusUnauthorized,
}

func HttpStatusCode(err Error) int {
	if code, ok := statusCodeMap[err.Type]; ok {
		return code
	}

	return http.StatusInternalServerError
}
