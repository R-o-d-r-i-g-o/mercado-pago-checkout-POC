package errors

import (
	"net/http"
	"testing"
)

func TestHttpStatusCode(t *testing.T) {
	tests := []struct {
		err  Error
		want int
	}{
		{
			err:  BusinessRule,
			want: http.StatusUnprocessableEntity,
		},
		{
			err:  Validation,
			want: http.StatusBadRequest,
		},
		{
			err:  NotFound,
			want: http.StatusNotFound,
		},
		{
			err:  Unknown,
			want: http.StatusInternalServerError,
		},
		{
			err:  Unathorizated,
			want: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.err.Type), func(t *testing.T) {
			if got := HttpStatusCode(tt.err); got != tt.want {
				t.Errorf("HttpStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
