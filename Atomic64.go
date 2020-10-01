package bitmap

import (
	"errors"
	"sync"
)

//Bitmap, which can contain numbers from 0 to 63
//Thread-safe
type Atomic64 struct {
	Index64
	mu sync.RWMutex
}

//Create new bitmap index
func NewAtomic64(values []int) (*Atomic64, error) {
	r := new(Atomic64)
	err := r.build(values)
	if nil != err {
		return nil, err
	}
	return r, nil
}

//Return true if val contains in bitmap
func (a *Atomic64) FindOne(val int) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Index64.FindOne(val)
}

//Returns true if bitmap contains all values
func (a *Atomic64) FindAll(values []int) bool {
	for _, val := range values {
		if !a.FindOne(val) {
			return false
		}
	}
	return true
}

//Returns true if bitmap contains exact one value from a list
func (a *Atomic64) FindLeastOne(values []int) bool {
	for _, val := range values {
		if a.FindOne(val) {
			return true
		}
	}
	return false
}

//Add val into bitmap
func (a *Atomic64) Add(val int) error {
	if val > 63 {
		return errors.New("type Atomic32 can't contain values more then 63")
	}
	bit := posFor64(val)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.set |= bit
	return nil
}

//Delete val from bitmap
func (a *Atomic64) Delete(val int) {
	if val > 63 {
		return
	}
	bit := posFor64(val)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.set = a.set &^ bit
}
