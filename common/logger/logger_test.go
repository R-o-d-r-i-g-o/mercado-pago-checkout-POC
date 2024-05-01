package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

func (l *LoggerTestSuite) TestMountLogger_Success() {
	test := l.Assert()

	engineLogger, err := New("development")

	test.Nil(err)
	test.NotEmpty(engineLogger)
}

func (l *LoggerTestSuite) TestMountLogger_Fail() {
	test := l.Assert()

	engineLogger, err := New("isInvalidMode")

	message := err.Error()
	test.Equal("cant initialize logger: invalid mode", message)
	test.Nil(engineLogger)
}

func (l *LoggerTestSuite) TestGetLogLevel_Fail() {
	test := l.Assert()

	_, err := getCurrentModeLog("isInvalidMode")
	message := err.Error()
	test.Equal("invalid mode", message)
}

func (l *LoggerTestSuite) TestGetLogLevel_Success() {
	test := l.Assert()

	_, errDev := getCurrentModeLog("development")
	test.Nil(errDev)

	_, errProd := getCurrentModeLog("production")
	test.Nil(errProd)
}
