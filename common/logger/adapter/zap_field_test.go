package adapter

import (
	"testing"

	"code-space-backend-api/common/logger/field"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zapcore"
)

type ZapFieldTestSuite struct {
	suite.Suite

	givenKey   string
	givenValue string

	givenZapFunction  ZapFieldFunc[string]
	givenZapCoreField zapcore.Field

	givenFieldArray []field.Field
}

func TestZapField(t *testing.T) {
	suite.Run(t, new(ZapFieldTestSuite))
}

func (z *ZapFieldTestSuite) SetupTest() {
	z.givenKey = "KEY"
	z.givenValue = "VALUE"

	z.givenZapFunction = func(key, value string) zapcore.Field {
		return zapcore.Field{
			Key:    key,
			String: value,
		}
	}

	z.givenZapCoreField = zapcore.Field{
		Key:    z.givenKey,
		String: z.givenValue,
	}

	z.givenFieldArray = []field.Field{
		{ZapField: z.givenZapCoreField},
		{ZapField: z.givenZapCoreField},
		{ZapField: z.givenZapCoreField},
	}
}

func (z *ZapFieldTestSuite) TestZapToFieldFunc_Success() {
	test := z.Assert()

	receivedFunction := ZapToFieldFunc(z.givenZapFunction)
	test.NotEmpty(receivedFunction)

	receivedField := receivedFunction(z.givenKey, z.givenValue)
	test.Equal(z.givenZapCoreField, receivedField.ZapField)
}

func (z *ZapFieldTestSuite) TestFieldsToZap_Success() {
	test := z.Assert()

	receivedZapFields := FieldsToZap(z.givenFieldArray)

	test.NotEmpty(receivedZapFields)
	test.Equal(cap(z.givenFieldArray), cap(receivedZapFields))
}
