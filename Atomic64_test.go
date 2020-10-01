package bitmap

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Atomic64TestSuite struct {
	suite.Suite
	testCase []int
	wrongTestCase []int
}

func TestAtomic64(t *testing.T) {
	suite.Run(t, new(Atomic64TestSuite))
}

func (s *Atomic64TestSuite) SetupSuite() {
	s.testCase = []int{1,2,4,8,9,10,11,12,23,25,30,31,40,50,60,63}
	s.wrongTestCase = []int{1,2,4,8,9,10,11,12,23,25,30,31,32,64}
}

func (s *Atomic64TestSuite) TestNew() {
	idx, err := NewAtomic64(s.wrongTestCase)
	s.NotNil(err)
	s.Nil(idx)
}

func (s *Atomic64TestSuite) TestFindOne() {
	index, err := NewAtomic64(s.testCase)
	s.Nil(err)
	s.True(index.FindOne(1))
	s.True(index.FindOne(40))
	s.False(index.FindOne(5))
	s.False(index.FindOne(66))
}

func (s *Atomic64TestSuite) TestFindAll() {
	index, err := NewAtomic64(s.testCase)
	s.Nil(err)
	s.True(index.FindAll([]int{1,4,8}))
	s.True(index.FindAll(s.testCase))
	s.False(index.FindAll([]int{1,4,8,0}))
	s.False(index.FindAll(s.wrongTestCase))
}

func (s *Atomic64TestSuite) TestFindLeastOne() {
	index, err := NewAtomic64(s.testCase)
	s.Nil(err)
	s.True(index.FindLeastOne([]int{1,3,5}))
	s.True(index.FindLeastOne(s.wrongTestCase))
	s.False(index.FindLeastOne([]int{3,5,7,0}))
}

func (s *Atomic64TestSuite) TestAdd() {
	index, err := NewAtomic64(s.testCase)
	s.Nil(err)
	s.NotNil(index)
	s.NotNil(index.Add(65))
	s.Nil(index.Add(42))
	s.True(index.FindOne(42))
}

func (s *Atomic64TestSuite) TestDelete() {
	index, err := NewAtomic64(s.testCase)
	s.Nil(err)
	s.NotNil(index)
	index.Delete(64)
	s.True(index.FindOne(40))
	index.Delete(40)
	s.False(index.FindOne(40))
}

