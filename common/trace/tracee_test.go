package trace

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TraceTestSuite struct {
	suite.Suite
}

func TestTraceSuite(t *testing.T) {
	suite.Run(t, new(TraceTestSuite))
}

func (t *TraceTestSuite) TestFromContext_Success() {
	test := t.Assert()

	mockedContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx := context.WithValue(mockedContext, Key, "hello")

	expectedValue := ctx.Value(Key).(string)
	receivedValue := FromContext(ctx)

	test.Equal(expectedValue, receivedValue)
}

func (t *TraceTestSuite) TestFromContext_Fail() {
	test := t.Assert()

	mockedContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	receivedValue := FromContext(mockedContext)

	test.Equal("", receivedValue)
}

func (t *TraceTestSuite) TestNewContext_Success() {
	test := t.Assert()

	mockedString := "hello"
	mockedContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	expectedContext := context.WithValue(mockedContext, Key, mockedString)
	receivedContext := NewContext(mockedContext, mockedString)

	test.Equal(expectedContext, receivedContext)
}
