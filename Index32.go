package bitmap

import (
	"errors"
)

//Bitmap, which can contain numbers from 0 to 31
//Immutable. Thread-safe
type Index32 struct {
	set uint32
}

func posFor32(val int) uint32 {
	return 1 << (val & bits32mask)
}

//Create new bitmap index
func NewIndex32(values []int) (*Index32, error) {
	r := new(Index32)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

func (i *Index32) build(values []int) error {
	for _, val := range values {
		if val > 31 {
			return errors.New("types Index32 or Atomic32 can't contain values more then 31")
		}
		bit := posFor32(val)
		i.set |= bit
	}
	return nil
}

//Return true if val contains in bitmap
func (i *Index32) FindOne(val int) bool {
	if val > 31 {
		return false
	}
	bit := posFor32(val)
	return i.set&bit != 0
}

//Returns true if bitmap contains all values
func (i *Index32) FindAll(values []int) bool {
	for _, val := range values {
		if !i.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (i *Index32) FindLeastOne(values []int) bool {
	for _, val := range values {
		if i.FindOne(val) {
			return true
		}
	}
	return false
}
