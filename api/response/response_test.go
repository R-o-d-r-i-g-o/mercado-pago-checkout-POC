package response

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ResponseSuite struct {
	suite.Suite

	responseError Response
}

func TestResponseSuite(t *testing.T) {
	suite.Run(t, new(ResponseSuite))
}

func (d *ResponseSuite) SetupTest() {
	d.responseError = Response{}
}

func (d *ResponseSuite) TestError_Sucess() {
	assert := d.Assert()

	response := Error(d.responseError)

	assert.Equal(response, Response{
		Error: d.responseError,
	})
}

func (d *ResponseSuite) TestData_Sucess() {
	assert := d.Assert()

	response := Data(d.responseError)

	assert.Equal(response, Response{
		Data: d.responseError,
	})
}
