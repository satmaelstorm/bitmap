package bitmap

import (
	"errors"
)

//Bitmap, which can contain numbers from 0 to 7
//Immutable. Thread-safe
type Index8 struct {
	set uint8
}

func posFor8(val int) uint8 {
	return 1 << (val & bits8mask)
}

//Create new bitmap index
func NewIndex8(values []int) (*Index8, error) {
	r := new(Index8)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

func (i *Index8) build(values []int) error {
	for _, val := range values {
		if val < 0 {
			return errors.New("types Index8 or Atomic8 can't contain values less then 0")
		}
		if val > 7 {
			return errors.New("types Index8 or Atomic8 can't contain values more then 7")
		}
		bit := posFor8(val)
		i.set |= bit
	}
	return nil
}

//Return true if val contains in bitmap
func (i *Index8) FindOne(val int) bool {
	if val < 0 {
		return false
	}
	if val > 7 {
		return false
	}
	bit := posFor8(val)
	return i.set&bit != 0
}

//Returns true if bitmap contains all values
func (i *Index8) FindAll(values []int) bool {
	for _, val := range values {
		if !i.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (i *Index8) FindLeastOne(values []int) bool {
	for _, val := range values {
		if i.FindOne(val) {
			return true
		}
	}
	return false
}

func (i *Index8) GobEncode() ([]byte, error) {
	return []byte{i.set}, nil
}

func (i *Index8) GobDecode(in []byte) error {
	if len(in) != 1 {
		return errors.New("bitmap Index8 must has 1 byte")
	}
	i.set = in[0]
	return nil
}
