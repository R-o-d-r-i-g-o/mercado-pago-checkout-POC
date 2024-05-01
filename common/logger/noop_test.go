package logger

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"code-space-backend-api/common/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type NewLoggerTestSuite struct {
	suite.Suite

	mockedContext *gin.Context
	mockError     error

	methods []string
}

func TestNewLoggerSuite(t *testing.T) {
	suite.Run(t, new(NewLoggerTestSuite))
}

func (l *NewLoggerTestSuite) SetupTest() {
	l.mockedContext, _ = gin.CreateTestContext(httptest.NewRecorder())
	l.mockError = errors.New().WithMessage("some error")

	l.methods = []string{"Debug", "Error", "Info", "SyncLogs", "WithContext"}
}

func (l *NewLoggerTestSuite) TestNewNoopLogger() {
	test := l.Assert()

	receivedValue := NewNoopLogger()
	expectedValue := noopLogger{}

	test.Equal(expectedValue, receivedValue)
}

func (l *NewLoggerTestSuite) TestWithContext() {
	test := l.Assert()

	receivedValue := NewNoopLogger().WithContext(l.mockedContext)
	expectedValue := noopLogger{}

	test.Equal(expectedValue, receivedValue)
}

func (l *NewLoggerTestSuite) TestError() {
	test := l.Assert()

	structWithMethods := reflect.TypeOf(noopLogger{})
	for i := 0; i < structWithMethods.NumMethod(); i++ {
		method := structWithMethods.Method(i)
		test.Equal(l.methods[i], method.Name)
	}
}
