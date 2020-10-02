package bitmap

type Immutable struct {
	index ImmutableBitmap
}

func NewImmutable(index ImmutableBitmap) *Immutable {
	return &Immutable{index: index}
}

func (i *Immutable) FindOne(val int) bool {
	return i.index.FindOne(val)
}

func (i *Immutable) FindAll(values []int) bool {
	return i.index.FindAll(values)
}

func (i *Immutable) FindLeastOne(values []int) bool {
	return i.index.FindLeastOne(values)
}
