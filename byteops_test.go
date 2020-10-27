package bitmap

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestByteOpsSuite struct {
	suite.Suite
}

func TestByteOps(t *testing.T) {
	suite.Run(t, new(TestByteOpsSuite))
}

func (s TestByteOpsSuite) TestInt64ToBytes() {
	r := Int64ToBytes(int64(^uint64(0) >> 1))
	for i, v := range r {
		if i == 7 {
			s.Equal(byte(127), v)
		} else {
			s.Equal(byte(255), v)
		}
	}
	r2 := BytesToInt64(r)
	s.Equal(int64(^uint64(0) >> 1), r2)
}

func (s TestByteOpsSuite) TestInt32ToBytes() {
	r := Int32ToBytes(int32(^uint32(0) >> 1))
	for i, v := range r {
		if i == 3 {
			s.Equal(byte(127), v)
		} else {
			s.Equal(byte(255), v)
		}
	}
	r2 := BytesToInt32(r)
	s.Equal(int32(^uint32(0) >> 1), r2)
}

func (s TestByteOpsSuite) TestInt16ToBytes() {
	r := Int16ToBytes(int16(^uint16(0) >> 1))
	for i, v := range r {
		if i == 1 {
			s.Equal(byte(127), v)
		} else {
			s.Equal(byte(255), v)
		}
	}
	r2 := BytesToInt16(r)
	s.Equal(int16(^uint16(0) >> 1), r2)
}