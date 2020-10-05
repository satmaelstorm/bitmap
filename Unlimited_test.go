package bitmap

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func prepareTestSet(q int, max int) []int {
	r := make([]int, 0, q)
	for i := 1; i < q; i++ {
		r = append(r, rand.Intn(max))
	}
	return r
}

func TestUnlimited_findAllByBuildNewBitmap(t *testing.T) {
	is := assert.New(t)
	bitmap, _ := NewUnlimited([]int{1, 3, 4, 8, 16, 20, 21, 22, 23})
	is.True(bitmap.findAllByBuildNewBitmap([]int{1, 3, 4, 8, 16, 20, 21, 22, 23}))
	is.True(bitmap.findAllByBuildNewBitmap([]int{20, 21, 22, 23}))
	is.True(bitmap.findAllByBuildNewBitmap([]int{1, 3, 4, 8, 16}))
	is.True(bitmap.findAllByBuildNewBitmap([]int{1, 16, 23}))
	is.False(bitmap.findAllByBuildNewBitmap([]int{1, 16, 23, 32}))
	is.False(bitmap.findAllByBuildNewBitmap([]int{1, 16, 23, 19}))
	is.False(bitmap.findAllByBuildNewBitmap([]int{-1, 1, 16, 23, 19}))
}

func BenchmarkUnlimited_findAllByBuildNewBitmap1000(b *testing.B) {
	r := prepareTestSet(1000, 10000000)
	bitmap, _ := NewUnlimited(r)
	bitmap2, _ := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAllByBitmap(bitmap2)
	}
}

func BenchmarkBitmap_findAllByBuildNewBitmap100(b *testing.B) {
	r := prepareTestSet(100, 10000000)
	bitmap, _ := NewUnlimited(r)
	bitmap2, _ := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAllByBitmap(bitmap2)
	}
}

func BenchmarkBitmap_findAllByBuildNewBitmap10(b *testing.B) {
	r := prepareTestSet(10, 10000000)
	bitmap, _ := NewUnlimited(r)
	bitmap2, _ := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAllByBitmap(bitmap2)
	}
}

func BenchmarkBitmap_findAllBrute1000(b *testing.B) {
	r := prepareTestSet(1000, 10000000)
	bitmap, _ := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAll(r)
	}
}

func BenchmarkBitmap_findAllBrute100(b *testing.B) {
	r := prepareTestSet(100, 10000000)
	bitmap, _ := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAll(r)
	}
}

func BenchmarkBitmap_findAllBrute10(b *testing.B) {
	r := prepareTestSet(10, 10000000)
	bitmap, _ := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAll(r)
	}
}
