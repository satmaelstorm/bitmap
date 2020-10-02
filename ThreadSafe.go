package bitmap

type ThreadSafe struct {
	index ThreadSafeBitmap
}

func NewThreadSafe(index ThreadSafeBitmap) *ThreadSafe {
	return &ThreadSafe{index: index}
}

func (t *ThreadSafe) Add(val int) error {
	return t.index.Add(val)
}

func (t *ThreadSafe) Delete(val int) {
	t.index.Delete(val)
}

func (t *ThreadSafe) FindOne(val int) bool {
	return t.index.FindOne(val)
}

func (t *ThreadSafe) FindAll(values []int) bool {
	return t.index.FindAll(values)
}

func (t *ThreadSafe) FindLeastOne(values []int) bool {
	return t.index.FindLeastOne(values)
}
