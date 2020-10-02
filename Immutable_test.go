package bitmap

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ImmutableTestSuite struct {
	suite.Suite
}

func TestImmutable(t *testing.T) {
	suite.Run(t, new(ImmutableTestSuite))
}

func (s *ImmutableTestSuite) TestUnlimited() {
	idx := NewImmutable(NewUnlimited([]int{1, 8, 2, 20, 128, 4444, 0, 3}))
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(4444))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{8, 4444, 2, 1}))
	s.True(idx.FindLeastOne([]int{3, 5, 4443, 4444}))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(5555555))
	s.False(idx.FindAll([]int{0, 1, 2, 3, 4, 4444}))
	s.False(idx.FindLeastOne([]int{5, 6, 7, 9}))

	idx2 := NewImmutable(NewUnlimited([]int{0}))
	s.True(idx2.FindOne(0))
	s.False(idx2.FindOne(1))
	s.True(idx2.FindAll([]int{0}))
	s.False(idx2.FindAll([]int{1}))
	s.False(idx2.FindAll([]int{}))
	s.True(idx2.FindLeastOne([]int{0}))
	s.False(idx2.FindLeastOne([]int{}))
	s.False(idx2.FindLeastOne([]int{1}))
}

func (s *ImmutableTestSuite) TestIndex64() {
	bm, err := NewIndex64([]int{0, 1, 3, 7, 8, 12, 15, 16, 22, 24, 31, 32, 40, 50, 63, 64})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex64([]int{0, 1, 3, 7, 8, 12, 15, 16, 22, 24, 31, 32, 40, 50, 63})
	s.Nil(err)
	s.NotNil(bm)
	idx := NewImmutable(bm)
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(63))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{3, 7, 8, 15, 16, 31, 32, 40, 50, 63}))
	s.True(idx.FindLeastOne([]int{2, 4, 50}))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(64))
	s.False(idx.FindAll([]int{40, 50, 63, 64}))
	s.False(idx.FindLeastOne([]int{41, 51, 62, 64}))
}

func (s *ImmutableTestSuite) TestIndex8() {
	bm, err := NewIndex8([]int{0, 1, 3, 7, 8})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex8([]int{0, 1, 3, 7})
	s.Nil(err)
	s.NotNil(bm)
	idx := NewImmutable(bm)
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(7))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{1, 7}))
	s.True(idx.FindLeastOne([]int{2, 1}))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(8))
	s.False(idx.FindAll([]int{0, 1, 2}))
	s.False(idx.FindLeastOne([]int{2, 4}))
}
