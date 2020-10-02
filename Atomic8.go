package bitmap

import (
	"errors"
	"sync"
)

//Bitmap, which can contain numbers from 0 to 7
//Thread-safe
type Atomic8 struct {
	Index8
	mu sync.RWMutex
}

//Create new bitmap index
func NewAtomic8(values []int) (*Atomic8, error) {
	r := new(Atomic8)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

//Return true if val contains in bitmap
func (a *Atomic8) FindOne(val int) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Index8.FindOne(val)
}

//Returns true if bitmap contains all values
func (a *Atomic8) FindAll(values []int) bool {
	for _, val := range values {
		if !a.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (a *Atomic8) FindLeastOne(values []int) bool {
	for _, val := range values {
		if a.FindOne(val) {
			return true
		}
	}
	return false
}

//Add val into bitmap
func (a *Atomic8) Add(val int) error {
	if val > 7 {
		return errors.New("type Atomic8 can't contain values more then 7")
	}
	bit := posFor8(val)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.set |= bit
	return nil
}

//Delete val from bitmap
func (a *Atomic8) Delete(val int) {
	if val > 7 {
		return
	}
	bit := posFor8(val)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.set = a.set &^ bit
}
