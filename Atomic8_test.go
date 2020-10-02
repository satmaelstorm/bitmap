package bitmap

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Atomic8TestSuite struct {
	suite.Suite
	testCase      []int
	wrongTestCase []int
}

func TestAtomic8(t *testing.T) {
	suite.Run(t, new(Atomic8TestSuite))
}

func (s *Atomic8TestSuite) SetupSuite() {
	s.testCase = []int{1, 2, 4}
	s.wrongTestCase = []int{1, 2, 4, 8}
}

func (s *Atomic8TestSuite) TestNew() {
	idx, err := NewAtomic8(s.wrongTestCase)
	s.NotNil(err)
	s.Nil(idx)
}

func (s *Atomic8TestSuite) TestFindOne() {
	index, err := NewAtomic8(s.testCase)
	s.Nil(err)
	s.True(index.FindOne(1))
	s.True(index.FindOne(4))
	s.False(index.FindOne(5))
	s.False(index.FindOne(12))
}

func (s *Atomic8TestSuite) TestFindAll() {
	index, err := NewAtomic8(s.testCase)
	s.Nil(err)
	s.True(index.FindAll([]int{1, 4}))
	s.True(index.FindAll(s.testCase))
	s.False(index.FindAll([]int{1, 4, 0}))
	s.False(index.FindAll(s.wrongTestCase))
}

func (s *Atomic8TestSuite) TestFindLeastOne() {
	index, err := NewAtomic8(s.testCase)
	s.Nil(err)
	s.True(index.FindLeastOne([]int{1, 3, 5}))
	s.True(index.FindLeastOne(s.wrongTestCase))
	s.False(index.FindLeastOne([]int{3, 5, 7, 0}))
}

func (s *Atomic8TestSuite) TestAdd() {
	index, err := NewAtomic8(s.testCase)
	s.Nil(err)
	s.NotNil(index)
	s.NotNil(index.Add(8))
	s.Nil(index.Add(3))
	s.True(index.FindOne(3))
}

func (s *Atomic8TestSuite) TestDelete() {
	index, err := NewAtomic8(s.testCase)
	s.Nil(err)
	s.NotNil(index)
	index.Delete(8)
	s.True(index.FindOne(4))
	index.Delete(4)
	s.False(index.FindOne(4))
	s.True(index.FindOne(2))
}
