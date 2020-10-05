package bitmap

import (
	"errors"
)

//Bitmap, which can contain numbers from 0 to 63
//Immutable. Thread-safe
type Index64 struct {
	set uint64
}

func posFor64(val int) uint64 {
	return 1 << (val & bits64mask)
}

//Create new bitmap index
func NewIndex64(values []int) (*Index64, error) {
	r := new(Index64)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

func (i *Index64) build(values []int) error {
	for _, val := range values {
		if val < 0 {
			return errors.New("types Index64 or Atomic64 can't contain values less then 0")
		}
		if val > 63 {
			return errors.New("types Index64 or Atomic32 can't contain values more then 63")
		}
		bit := posFor64(val)
		i.set |= bit
	}
	return nil
}

//Return true if val contains in bitmap
func (i *Index64) FindOne(val int) bool {
	if val < 0 {
		return false
	}
	if val > 63 {
		return false
	}
	bit := posFor64(val)
	return i.set&bit != 0
}

//Returns true if bitmap contains all values
func (i *Index64) FindAll(values []int) bool {
	for _, val := range values {
		if !i.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (i *Index64) FindLeastOne(values []int) bool {
	for _, val := range values {
		if i.FindOne(val) {
			return true
		}
	}
	return false
}
