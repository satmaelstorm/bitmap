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

func TestUnlimited_Build(t *testing.T) {
	is := assert.New(t)
	b := Unlimited{
		set: []byte{0},
		len: 1,
	}
	b.build([]int{0})
	is.True(b.FindOne(0))
}

func TestUnlimited_FindOne(t *testing.T) {
	is := assert.New(t)
	bitmap := NewUnlimited([]int{1,3,4,8,16,20,2367423})
	is.True(bitmap.FindOne(1))
	is.True(bitmap.FindOne(3))
	is.True(bitmap.FindOne(4))
	is.True(bitmap.FindOne(8))
	is.True(bitmap.FindOne(16))
	is.True(bitmap.FindOne(20))
	is.True(bitmap.FindOne(2367423))
	is.False(bitmap.FindOne(0))
	is.False(bitmap.FindOne(5))
	is.False(bitmap.FindOne(213213))
	bitmap2 := NewUnlimited([]int{0})
	is.True(bitmap2.FindOne(0))
	is.False(bitmap2.FindOne(1))
	is.False(bitmap2.FindOne(3))
	is.False(bitmap2.FindOne(20))
	is.False(bitmap2.FindOne(2367423))
	is.False(bitmap2.FindOne(5))
	is.False(bitmap2.FindOne(213213))
}

func TestUnlimited_findAllByBuildNewBitmap(t *testing.T) {
	is := assert.New(t)
	bitmap := NewUnlimited([]int{1,3,4,8,16,20,21,22,23})
	is.True(bitmap.findAllByBuildNewBitmap([]int{1,3,4,8,16,20,21,22,23}))
	is.True(bitmap.findAllByBuildNewBitmap([]int{20,21,22,23}))
	is.True(bitmap.findAllByBuildNewBitmap([]int{1,3,4,8,16}))
	is.True(bitmap.findAllByBuildNewBitmap([]int{1,16,23}))
	is.False(bitmap.findAllByBuildNewBitmap([]int{1,16,23,32}))
	is.False(bitmap.findAllByBuildNewBitmap([]int{1,16,23,19}))
}

func TestUnlimited_FindAll(t *testing.T) {
	is := assert.New(t)
	bitmap := NewUnlimited([]int{1,3,4,8,16,20,21,22,23})
	is.True(bitmap.FindAll([]int{1,3,4,8,16,20,21,22,23}))
	is.True(bitmap.FindAll([]int{20,21,22,23}))
	is.True(bitmap.FindAll([]int{1,3,4,8,16}))
	is.True(bitmap.FindAll([]int{1,16,23}))
	is.False(bitmap.FindAll([]int{1,16,23,32}))
}

func TestUnlimited_FindLeastOne(t *testing.T) {
	is := assert.New(t)
	bitmap := NewUnlimited([]int{1,3,4,8,16,20,21,22,23})
	is.True(bitmap.FindLeastOne([]int{1,2,5,3,4}))
	is.True(bitmap.FindLeastOne([]int{2,5,20,0,100}))
	is.True(bitmap.FindLeastOne([]int{100,101,23}))
	is.True(bitmap.FindLeastOne([]int{2,5,6,7,8}))
	is.False(bitmap.FindLeastOne([]int{2,5,6,7,9}))
}

func BenchmarkUnlimited_findAllByBuildNewBitmap1000(b *testing.B) {
	r := prepareTestSet(1000, 10000000)
	bitmap := NewUnlimited(r)
	bitmap2 := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAllByBitmap(bitmap2)
	}
}

func BenchmarkBitmap_findAllByBuildNewBitmap100(b *testing.B) {
	r := prepareTestSet(100, 10000000)
	bitmap := NewUnlimited(r)
	bitmap2 := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAllByBitmap(bitmap2)
	}
}

func BenchmarkBitmap_findAllByBuildNewBitmap10(b *testing.B) {
	r := prepareTestSet(10, 10000000)
	bitmap := NewUnlimited(r)
	bitmap2 := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAllByBitmap(bitmap2)
	}
}

func BenchmarkBitmap_findAllBrute1000(b *testing.B) {
	r := prepareTestSet(1000, 10000000)
	bitmap := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAll(r)
	}
}

func BenchmarkBitmap_findAllBrute100(b *testing.B) {
	r := prepareTestSet(100, 10000000)
	bitmap := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAll(r)
	}
}

func BenchmarkBitmap_findAllBrute10(b *testing.B) {
	r := prepareTestSet(10, 10000000)
	bitmap := NewUnlimited(r)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmap.FindAll(r)
	}
}
