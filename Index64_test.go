package bitmap

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Index64TestSuite struct {
	suite.Suite
	testCase      []int
	wrongTestCase []int
}

func TestIndex64(t *testing.T) {
	suite.Run(t, new(Index64TestSuite))
}

func (s *Index64TestSuite) SetupSuite() {
	s.testCase = []int{1, 2, 4, 8, 9, 10, 11, 12, 23, 25, 30, 31, 40, 50, 60, 63}
	s.wrongTestCase = []int{1, 2, 4, 8, 9, 10, 11, 12, 23, 25, 30, 31, 32, 64}
}

func (s *Index64TestSuite) TestNew() {
	idx, err := NewIndex64(s.wrongTestCase)
	s.NotNil(err)
	s.Nil(idx)
}

func (s *Index64TestSuite) TestFindOne() {
	index, err := NewIndex64(s.testCase)
	s.Nil(err)
	s.True(index.FindOne(1))
	s.True(index.FindOne(40))
	s.False(index.FindOne(5))
	s.False(index.FindOne(66))
}

func (s *Index64TestSuite) TestFindAll() {
	index, err := NewIndex64(s.testCase)
	s.Nil(err)
	s.True(index.FindAll([]int{1, 4, 8}))
	s.True(index.FindAll(s.testCase))
	s.False(index.FindAll([]int{1, 4, 8, 0}))
	s.False(index.FindAll(s.wrongTestCase))
}

func (s *Index64TestSuite) TestFindLeastOne() {
	index, err := NewIndex64(s.testCase)
	s.Nil(err)
	s.True(index.FindLeastOne([]int{1, 3, 5}))
	s.True(index.FindLeastOne(s.wrongTestCase))
	s.False(index.FindLeastOne([]int{3, 5, 7, 0}))
}
