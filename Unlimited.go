package bitmap

//Unlimited bitmap. Immutable. Thread-safe
type Unlimited struct {
	set []byte
	len int
}

func posForUnlimited(val int) (int, byte) {
	return val >> bits8, byte(1 << (val & bits8mask))
}

//Create bew Unlimited bitmap
func NewUnlimited(values []int) *Unlimited {
	bitmap := &Unlimited{set: []byte{0}, len: 1}
	if len(values) > 0 {
		bitmap.build(values)
	}
	return bitmap
}

func (u *Unlimited) build(values []int) {
	for _, val := range values {
		idx, bit := posForUnlimited(val)
		if idx >= u.len {
			u.extend(idx)
		}
		u.set[idx] |= bit
	}
}

func (u *Unlimited) extend(idx int) {
	newSet := make([]byte, idx+1)
	copy(newSet, u.set)
	u.len = len(newSet)
	u.set = newSet
}

//Return true if val contains in bitmap
func (u *Unlimited) FindOne(val int) bool {
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
	newBitmap := NewUnlimited(values)
	return u.FindAllByBitmap(newBitmap)
}
