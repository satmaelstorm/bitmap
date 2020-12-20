package bitmap

type Error struct {
	Message string
	Err     error
}

func (be *Error) Error() string {
	return be.Message + ": " + be.Err.Error()
}
