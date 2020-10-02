package bitmap

import (
	"errors"
)

//Bitmap, which can contain numbers from 0 to 15
//Immutable. Thread-safe
type Index16 struct {
	set uint16
}

func posFor16(val int) uint16 {
	return 1 << (val & bits16mask)
}

//Create new bitmap index
func NewIndex16(values []int) (*Index16, error) {
	r := new(Index16)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

func (i *Index16) build(values []int) error {
	for _, val := range values {
		if val > 15 {
			return errors.New("types Index16 or Atomic16 can't contain values more then 15")
		}
		bit := posFor16(val)
		i.set |= bit
	}
	return nil
}

//Return true if val contains in bitmap
func (i *Index16) FindOne(val int) bool {
	if val > 15 {
		return false
	}
	bit := posFor16(val)
	return i.set&bit != 0
}

//Returns true if bitmap contains all values
func (i *Index16) FindAll(values []int) bool {
	for _, val := range values {
		if !i.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (i *Index16) FindLeastOne(values []int) bool {
	for _, val := range values {
		if i.FindOne(val) {
			return true
		}
	}
	return false
}
