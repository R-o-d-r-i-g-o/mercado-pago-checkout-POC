package query

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestQuery(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}

type QueryTestSuite struct {
	suite.Suite
}

type StructForTest struct {
	HelloWorld string `json:"hello_world"`
}

func (m *QueryTestSuite) TestReplaceOnlyOne() {
	test := m.Assert()

	expected := "select * from test where test = 1"
	query := New("select * from test where test = $query")

	test.Equal(expected, query.Replace(1).String())
}

func (m *QueryTestSuite) TestReplaceAll_OnlyOneParams() {
	test := m.Assert()

	expected := `select * from test where test = '1' and email = '1'`
	query := New("select * from test where test = $query and email = $query").
		Replace("1").String()

	test.Equal(expected, query)
}

func (m *QueryTestSuite) TestReplace_moreOneParams() {
	test := m.Assert()

	expected := `select * from test where test = '1' and email = 'dev@q2pay.com.br'`
	query := New("select * from test where test = $query and email = $query").
		Replace("1", "dev@q2pay.com.br").String()

	test.Equal(expected, query)
}

func (m *QueryTestSuite) TestReplace_WhenPassMoreParamsNeeded() {
	test := m.Assert()

	expected := `select * from test where test = '1' and email = 'dev@q2pay.com.br'`
	query := New("select * from test where test = $query and email = $query").
		Replace("1", "dev@q2pay.com.br", "3", "4").String()

	test.Equal(expected, query)
}

func (m *QueryTestSuite) TestReplace_TwoParamsAnotherType() {
	test := m.Assert()

	expected := `select * from test where test = 1 and email = 'dev@q2pay.com.br'`
	query := New("select * from test where test = $query and email = $query").
		Replace(1, "dev@q2pay.com.br", "3", 4).String()

	test.Equal(expected, query)
}
