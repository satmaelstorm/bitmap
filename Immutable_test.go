package bitmap

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ImmutableTestSuite struct {
	suite.Suite
}

func TestImmutable(t *testing.T) {
	suite.Run(t, new(ImmutableTestSuite))
}

func checkUnlimited(s *ImmutableTestSuite, idx *Unlimited) {
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(4444))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{8, 4444, 2, 1}))
	s.True(idx.FindLeastOne([]int{3, 5, 4443, 4444}))
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(5555555))
	s.False(idx.FindAll([]int{0, 1, 2, 3, 4, 4444}))
	s.False(idx.FindLeastOne([]int{5, 6, 7, 9}))
}

func (s *ImmutableTestSuite) TestUnlimited() {
	bm, err := NewUnlimited([]int{-1})
	s.Nil(bm)
	s.NotNil(err)
	idx, err := NewUnlimited([]int{1, 8, 2, 20, 128, 4444, 0, 3})
	s.Nil(err)
	s.NotNil(idx)
	checkUnlimited(s, idx)

	idx2, _ := NewUnlimited([]int{0})
	s.True(idx2.FindOne(0))
	s.False(idx2.FindOne(1))
	s.True(idx2.FindAll([]int{0}))
	s.False(idx2.FindAll([]int{1}))
	s.False(idx2.FindAll([]int{}))
	s.True(idx2.FindLeastOne([]int{0}))
	s.False(idx2.FindLeastOne([]int{}))
	s.False(idx2.FindLeastOne([]int{1}))

	var idx3 *Unlimited
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(idx)
	s.Nil(err)
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&idx3)
	s.Nil(err)

	checkUnlimited(s, idx3)

	js, err := json.Marshal(idx)
	fmt.Println(js)
	s.Nil(err)
	s.NotNil(js)

	var idx4 *Unlimited
	err = json.Unmarshal(js, &idx4)
	s.Nil(err)

	checkUnlimited(s, idx4)
}

func (s *ImmutableTestSuite) TestIndex64() {
	bm, err := NewIndex64([]int{0, 1, 3, 7, 8, 12, 15, 16, 22, 24, 31, 32, 40, 50, 63, 64})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex64([]int{-1})
	s.Nil(bm)
	s.NotNil(err)
	idx, err := NewIndex64([]int{0, 1, 3, 7, 8, 12, 15, 16, 22, 24, 31, 32, 40, 50, 63})
	s.Nil(err)
	s.NotNil(idx)
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(63))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{3, 7, 8, 15, 16, 31, 32, 40, 50, 63}))
	s.True(idx.FindLeastOne([]int{2, 4, 50}))
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(64))
	s.False(idx.FindAll([]int{40, 50, 63, 64}))
	s.False(idx.FindLeastOne([]int{41, 51, 62, 64}))

	var idx2 Index64
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(idx)
	s.Nil(err)
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&idx2)
	s.Nil(err)

	s.True(idx2.FindOne(1))
	s.True(idx2.FindOne(0))
	s.True(idx2.FindOne(63))
	s.True(idx2.FindOne(3))
	s.True(idx2.FindAll([]int{3, 7, 8, 15, 16, 31, 32, 40, 50, 63}))
	s.True(idx2.FindLeastOne([]int{2, 4, 50}))
	s.False(idx2.FindOne(-1))
	s.False(idx2.FindOne(4))
	s.False(idx2.FindOne(64))
	s.False(idx2.FindOne(62))
	s.False(idx2.FindAll([]int{40, 50, 63, 64}))
	s.False(idx2.FindLeastOne([]int{41, 51, 62, 64}))
}

func (s *ImmutableTestSuite) TestIndex32() {
	bm, err := NewIndex32([]int{0, 1, 3, 7, 8, 12, 15, 16, 22, 24, 31, 32})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex32([]int{-1})
	s.Nil(bm)
	s.NotNil(err)
	idx, err := NewIndex32([]int{0, 1, 3, 7, 8, 12, 15, 16, 22, 24, 31})
	s.Nil(err)
	s.NotNil(idx)
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(31))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{3, 7, 8, 15, 16, 31}))
	s.True(idx.FindLeastOne([]int{2, 31, 50}))
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(32))
	s.False(idx.FindAll([]int{1, 3, 8, 30}))
	s.False(idx.FindLeastOne([]int{25, 30, 20, 32}))

	var idx2 Index32
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(idx)
	s.Nil(err)
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&idx2)
	s.Nil(err)

	s.True(idx2.FindOne(1))
	s.True(idx2.FindOne(0))
	s.True(idx2.FindOne(31))
	s.True(idx2.FindOne(3))
	s.True(idx2.FindAll([]int{3, 7, 8, 15, 16, 31}))
	s.True(idx2.FindLeastOne([]int{2, 31, 50}))
	s.False(idx2.FindOne(-1))
	s.False(idx2.FindOne(4))
	s.False(idx2.FindOne(30))
	s.False(idx2.FindOne(32))
	s.False(idx2.FindAll([]int{1, 3, 8, 30}))
	s.False(idx2.FindLeastOne([]int{25, 30, 20, 32}))
}

func (s *ImmutableTestSuite) TestIndex16() {
	bm, err := NewIndex16([]int{0, 1, 3, 7, 8, 12, 15, 16})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex16([]int{-1})
	s.Nil(bm)
	s.NotNil(err)
	idx, err := NewIndex16([]int{0, 1, 3, 7, 8, 12, 15})
	s.Nil(err)
	s.NotNil(idx)
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(15))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{3, 7, 8, 15}))
	s.True(idx.FindLeastOne([]int{2, 15, 50}))
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(16))
	s.False(idx.FindAll([]int{1, 3, 8, 16}))
	s.False(idx.FindLeastOne([]int{2, 6, 14, 16}))

	var idx2 Index16
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(idx)
	s.Nil(err)
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&idx2)
	s.Nil(err)

	s.True(idx2.FindOne(1))
	s.True(idx2.FindOne(0))
	s.True(idx2.FindOne(15))
	s.True(idx2.FindOne(3))
	s.True(idx2.FindAll([]int{3, 7, 8, 15}))
	s.True(idx2.FindLeastOne([]int{2, 15, 50}))
	s.False(idx2.FindOne(-1))
	s.False(idx2.FindOne(4))
	s.False(idx2.FindOne(14))
	s.False(idx2.FindOne(16))
	s.False(idx2.FindAll([]int{1, 3, 8, 16}))
	s.False(idx2.FindLeastOne([]int{2, 6, 14, 16}))
}

func (s *ImmutableTestSuite) TestIndex8() {
	bm, err := NewIndex8([]int{0, 1, 3, 7, 8})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex8([]int{-1})
	s.Nil(bm)
	s.NotNil(err)
	idx, err := NewIndex8([]int{0, 1, 3, 7})
	s.Nil(err)
	s.NotNil(idx)
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(7))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{1, 7}))
	s.True(idx.FindLeastOne([]int{2, 1}))
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(8))
	s.False(idx.FindAll([]int{0, 1, 2}))
	s.False(idx.FindLeastOne([]int{2, 4}))

	var idx2 Index8
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(idx)
	s.Nil(err)
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&idx2)
	s.Nil(err)

	s.True(idx.FindOne(1))
	s.True(idx.FindOne(0))
	s.True(idx.FindOne(7))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{1, 7}))
	s.True(idx.FindLeastOne([]int{2, 1}))
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(6))
	s.False(idx.FindOne(8))
	s.False(idx.FindAll([]int{0, 1, 2}))
	s.False(idx.FindLeastOne([]int{2, 4}))
}

func (s *ImmutableTestSuite) TestSmart() {
	idx := NewSmart([]int{-1, 1, 3, 7, 8, -2})
	s.True(idx.FindOne(1))
	s.True(idx.FindOne(7))
	s.True(idx.FindOne(3))
	s.True(idx.FindAll([]int{1, 7}))
	s.True(idx.FindLeastOne([]int{2, 1}))
	s.True(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(9))
	s.False(idx.FindAll([]int{-1, 1, 2}))
	s.False(idx.FindLeastOne([]int{2, 4}))
	s.Equal(-2, idx.GetMin())
	s.Equal(8, idx.GetMax())
	idx = NewSmart([]int{-1, 1})
	s.Equal(-1, idx.GetMin())
	s.Equal(1, idx.GetMax())
	idx = NewSmart([]int{-1, 15})
	s.Equal(-1, idx.GetMin())
	s.Equal(15, idx.GetMax())
	idx = NewSmart([]int{-1, 32})
	s.Equal(-1, idx.GetMin())
	s.Equal(32, idx.GetMax())
	idx = NewSmart([]int{-1, 64})
	s.Equal(-1, idx.GetMin())
	s.Equal(64, idx.GetMax())
}

func (s *ImmutableTestSuite) TestImmutable() {
	bm, err := NewIndex8([]int{0, 1, 3, 7, 8})
	s.Nil(bm)
	s.NotNil(err)
	bm, err = NewIndex8([]int{-1})
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
	s.False(idx.FindOne(-1))
	s.False(idx.FindOne(4))
	s.False(idx.FindOne(8))
	s.False(idx.FindAll([]int{0, 1, 2}))
	s.False(idx.FindLeastOne([]int{2, 4}))
}
