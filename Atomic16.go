package bitmap

import (
	"errors"
	"sync"
)

//Bitmap, which can contain numbers from 0 to 15
//Thread-safe
type Atomic16 struct {
	Index16
	mu sync.RWMutex
}

//Create new bitmap index
func NewAtomic16(values []int) (*Atomic16, error) {
	r := new(Atomic16)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

//Return true if val contains in bitmap
func (a *Atomic16) FindOne(val int) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Index16.FindOne(val)
}

//Returns true if bitmap contains all values
func (a *Atomic16) FindAll(values []int) bool {
	for _, val := range values {
		if !a.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (a *Atomic16) FindLeastOne(values []int) bool {
	for _, val := range values {
		if a.FindOne(val) {
			return true
		}
	}
	return false
}

//Add val into bitmap
func (a *Atomic16) Add(val int) error {
	if val > 15 {
		return errors.New("type Atomic16 can't contain values more then 15")
	}
	bit := posFor16(val)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.set |= bit
	return nil
}

//Delete val from bitmap
func (a *Atomic16) Delete(val int) {
	if val > 15 {
		return
	}
	bit := posFor16(val)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.set = a.set &^ bit
}
