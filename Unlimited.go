package bitmap

import (
	"encoding/json"
	"errors"
)

//Unlimited bitmap. Immutable. Thread-safe
type Unlimited struct {
	set []byte
	len int
}

func posForUnlimited(val int) (int, byte) {
	return val >> bits8, byte(1 << (val & bits8mask))
}

//Create bew Unlimited bitmap
func NewUnlimited(values []int) (*Unlimited, error) {
	bitmap := &Unlimited{set: []byte{0}, len: 1}
	if len(values) > 0 {
		err := bitmap.build(values)
		if err != nil {
			return nil, err
		}
	}
	return bitmap, nil
}

func (u *Unlimited) build(values []int) error {

	for _, val := range values {
		if val < 0 {
			return errors.New("type Unlimited can't contain values less then 0")
		}
		idx, bit := posForUnlimited(val)
		if idx >= u.len {
			u.extend(idx)
		}
		u.set[idx] |= bit
	}
	return nil
}

func (u *Unlimited) extend(idx int) {
	newSet := make([]byte, idx+1)
	copy(newSet, u.set)
	u.len = len(newSet)
	u.set = newSet
}

//Return true if val contains in bitmap
func (u *Unlimited) FindOne(val int) bool {
	if val < 0 {
		return false
	}
	idx, bit := posForUnlimited(val)
	if idx >= u.len {
		return false
	}
	return u.set[idx]&bit != 0
}

//Return true if all number from bitmap `source` present in bitmap
func (u *Unlimited) FindAllByBitmap(source *Unlimited) bool {
	if source.len > u.len {
		return false
	}
	for idx, bits := range source.set {
		if bits&u.set[idx] != bits {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains all values
func (u *Unlimited) FindAll(values []int) bool {
	if len(values) < 1 {
		return false
	}
	for _, val := range values {
		if !u.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (u *Unlimited) FindLeastOne(values []int) bool {
	for _, val := range values {
		if u.FindOne(val) {
			return true
		}
	}
	return false
}

func (u *Unlimited) findAllByBuildNewBitmap(values []int) bool {
	newBitmap, err := NewUnlimited(values)
	if err != nil {
		return false
	}
	return u.FindAllByBitmap(newBitmap)
}

func (u *Unlimited) GobEncode() ([]byte, error) {
	return u.set, nil
}

func (u *Unlimited) GobDecode(in []byte) error {
	u.len = len(in)
	u.set = make([]byte, u.len)
	for i, v := range in {
		u.set[i] = v
	}
	return nil
}

func (u *Unlimited) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.set)
}

func (u *Unlimited) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &u.set)
	if err != nil {
		return &Error{Message: "Unlimited Bitmap UnmarshalJSON", Err: err}
	}
	u.len = len(u.set)
	return nil
}
