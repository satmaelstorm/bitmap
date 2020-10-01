package bitmap

type IBitmap interface {
	FindOne(val int) bool
	FindAll(values []int) bool
	FindLeastOne(values []int) bool
}
