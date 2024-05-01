package uuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UuidTestSuite struct {
	suite.Suite

	givenUuid V4
}

func TestUuidSuite(t *testing.T) {
	suite.Run(t, new(UuidTestSuite))
}

func (u *UuidTestSuite) SetupTest() {
	u.givenUuid = V4("hello World")
}

func (u *UuidTestSuite) TestNewUUID_Success() {
	test := u.Assert()

	receivedId := New()
	_, err := uuid.Parse(receivedId)

	test.Nil(err)
}

func (u *UuidTestSuite) TestIsValid_Success() {
	test := u.Assert()

	receivedBoolean := IsValid(uuid.NewString())

	test.True(receivedBoolean)
}

func (u *UuidTestSuite) TestV4String() {
	test := u.Assert()

	receivedValue := u.givenUuid.String()
	expectedValue := "hello World"

	test.NotEmpty(receivedValue)
	test.Equal(expectedValue, receivedValue)
}
