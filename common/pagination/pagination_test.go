package pagination

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PaginationTestSuite struct {
	suite.Suite
}

func TestPaginationSuite(t *testing.T) {
	suite.Run(t, new(PaginationTestSuite))
}

func (p *PaginationTestSuite) TestCurrentOn_Success() {
	test := p.Assert()
	mockedPageNumber := 1

	expectedPage := Page[string]{CurrentPage: mockedPageNumber}
	receivedPage := expectedPage.CurrentOn(mockedPageNumber)

	test.Equal(expectedPage.CurrentPage, receivedPage.CurrentPage)
}

func (p *PaginationTestSuite) TestLimitedTo_Success() {
	test := p.Assert()
	mockedLimitPage := 1

	expectedPage := Page[string]{Limit: mockedLimitPage}
	receivedPage := expectedPage.LimitedTo(mockedLimitPage)

	test.Equal(expectedPage.Limit, receivedPage.Limit)
}

func (p *PaginationTestSuite) TestTotalOf_Success() {
	test := p.Assert()
	mockedTotalCount := 1

	expectedPage := Page[string]{TotalCount: int64(mockedTotalCount)}
	receivedPage := expectedPage.TotalOf(int64(mockedTotalCount))

	test.Equal(expectedPage.TotalCount, receivedPage.TotalCount)
}

func (p *PaginationTestSuite) TestWith_Success() {
	test := p.Assert()
	mockedPageData := []string{"hello", " world"}

	expectedPage := Page[string]{
		Data:  mockedPageData,
		Count: len(mockedPageData),
	}
	expectedPage.TotalPages = int64(math.Ceil(float64(expectedPage.TotalCount) / float64(expectedPage.Limit)))
	receivedPage := expectedPage.With(mockedPageData)

	test.Equal(expectedPage, receivedPage)
}
