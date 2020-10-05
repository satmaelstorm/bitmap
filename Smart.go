package bitmap

type Smart struct {
	index *Immutable
	max   int
	min   int
}

func minMax(values []int) (min int, max int) {
	for _, e := range values {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	return min, max
}
func normalize(values []int, min int) []int {
	l := len(values)
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = values[i] - min
	}
	return result
}

func NewSmart(values []int) *Smart {
	min, max := minMax(values)
	diff := max - min
	normalized := normalize(values, min)
	result := new(Smart)
	result.max = max
	result.min = min
	switch  {
	case diff < 8:
		idx, _ := NewIndex8(normalized)
		result.index = NewImmutable(idx)
	case diff >=9 && diff < 16:
		idx, _ := NewIndex16(normalized)
		result.index = NewImmutable(idx)
	case diff >=16 && diff < 32:
		idx, _ := NewIndex32(normalized)
		result.index = NewImmutable(idx)
	case diff >=32 && diff < 64:
		idx, _ := NewIndex64(normalized)
		result.index = NewImmutable(idx)
	default:
		idx, _ := NewUnlimited(normalized)
		result.index = NewImmutable(idx)
	}
	return result
}

func (s *Smart) GetMax() int {
	return s.max
}

func (s *Smart) GetMin() int {
	return s.min
}

func (s *Smart) FindOne(val int) bool {
	return s.index.FindOne(val - s.min)
}

func (s *Smart) FindAll(values []int) bool {
	return s.index.FindAll(normalize(values, s.min))
}

func (s *Smart) FindLeastOne(values []int) bool {
	return s.index.FindLeastOne(normalize(values, s.min))
}
