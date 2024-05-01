package adapter

import (
	defaultError "errors"
	"testing"

	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/logger/adapter/mocks"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zapcore"
)

type ZapErrorTestSuite struct {
	suite.Suite

	givenObjectEncoder          *mocks.ObjectEncoder
	givenArrayEncoder           *mocks.ArrayEncoder
	givenDefaulError            error
	givenError                  errors.Error
	givenDefaultErrorMarshaller zapErrorMarshaller
	givenMarshaller             zapErrorMarshaller
}

func TestZapError(t *testing.T) {
	suite.Run(t, new(ZapErrorTestSuite))
}

func (z *ZapErrorTestSuite) SetupTest() {
	z.givenObjectEncoder = mocks.NewObjectEncoder(z.T())
	z.givenArrayEncoder = mocks.NewArrayEncoder(z.T())

	z.givenDefaulError = defaultError.New("I am an error")
	z.givenDefaultErrorMarshaller = zapErrorMarshaller{z.givenDefaulError}

	z.givenError = errors.New().Add(z.givenDefaulError)
	z.givenMarshaller = zapErrorMarshaller{z.givenError}
}

func (z *ZapErrorTestSuite) TestZapError_Success() {
	test := z.Assert()

	var expectedObjectMarshaler zapcore.ObjectMarshaler = &zapErrorMarshaller{z.givenDefaulError}

	receivedObjectMarshaler := ZapError(z.givenDefaulError)

	test.Equal(expectedObjectMarshaler, receivedObjectMarshaler)
}

func (z *ZapErrorTestSuite) TestMarshalLogObject_TypeAssertion_Success() {
	assert := z.Assert()

	z.givenObjectEncoder.On("AddString", "type", string(z.givenError.Type)).Maybe()

	z.givenObjectEncoder.On("AddArray", "errors", &z.givenMarshaller).Return(defaultError.New("djfhsk"))

	assert.Nil(z.givenMarshaller.MarshalLogObject(z.givenObjectEncoder))
}

func (z *ZapErrorTestSuite) TestMarshalLogObject_Success() {
	assert := z.Assert()

	z.givenObjectEncoder.On("AddString", "message", z.givenDefaultErrorMarshaller.err.Error()).Times(1)

	assert.Nil(z.givenDefaultErrorMarshaller.MarshalLogObject(z.givenObjectEncoder))
}

func (z *ZapErrorTestSuite) TestMarshalLogArray() {
	err, _ := z.givenMarshaller.err.(errors.Error)
	z.givenArrayEncoder.On("AppendObject", &zapErrorMarshaller{err.Errors[0]}).Return(nil)

	z.givenMarshaller.MarshalLogArray(z.givenArrayEncoder)
}

func (z *ZapErrorTestSuite) TestMarshalLogArray_NoAppend() {
	assert := z.Assert()

	assert.Nil(z.givenDefaultErrorMarshaller.MarshalLogArray(z.givenArrayEncoder))
}

func (z *ZapErrorTestSuite) TestAddStringOmitEmpty_Success() {
	mockedString := "hello"

	z.givenObjectEncoder.On("AddString", mockedString, mockedString).Maybe()

	addStringOmitEmpty(z.givenObjectEncoder, mockedString, mockedString)
}
