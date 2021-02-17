package sort

import "sort"

// genome:
// +int, +int8, +int16, +int32, +int64
// +float32, +float64
// rune, +string
// +uint, +uint8, +uint16, +uint32, +uint64

type Int8Slice []int8

func (p Int8Slice) Len() int           { return len(p) }
func (p Int8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int8Slice) Sort() { sort.Sort(p) }

// Int8s sorts a slice of int8s in increasing order.
func Int8s(a []int8) { sort.Sort(Int8Slice(a)) }

// Int8sAreSorted tests whether a slice of int8s is sorted in increasing order.
func Int8sAreSorted(a []int8) bool { return sort.IsSorted(Int8Slice(a)) }

type Int16Slice []int16

func (p Int16Slice) Len() int           { return len(p) }
func (p Int16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int16Slice) Sort() { sort.Sort(p) }

// Int16s sorts a slice of int16s in increasing order.
func Int16s(a []int16) { sort.Sort(Int16Slice(a)) }

// Int16sAreSorted tests whether a slice of int16s is sorted in increasing order.
func Int16sAreSorted(a []int16) bool { return sort.IsSorted(Int16Slice(a)) }

type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int32Slice) Sort() { sort.Sort(p) }

// Int32s sorts a slice of int32s in increasing order.
func Int32s(a []int32) { sort.Sort(Int32Slice(a)) }

// Int32sAreSorted tests whether a slice of int32s is sorted in increasing order.
func Int32sAreSorted(a []int32) bool { return sort.IsSorted(Int32Slice(a)) }

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Int64Slice) Sort() { sort.Sort(p) }

// Int64s sorts a slice of int64s in increasing order.
func Int64s(a []int64) { sort.Sort(Int64Slice(a)) }

// Int64sAreSorted tests whether a slice of int64s is sorted in increasing order.
func Int64sAreSorted(a []int64) bool { return sort.IsSorted(Int64Slice(a)) }

type Float32Slice []float32

func (p Float32Slice) Len() int           { return len(p) }
func (p Float32Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Float32Slice) Sort() { sort.Sort(p) }

// Float32s sorts a slice of float32s in increasing order
// (not-a-number values are treated as less than other values).
func Float32s(a []float32) { sort.Sort(Float32Slice(a)) }

// Float32sAreSorted tests whether a slice of float32s is sorted in increasing order
// (not-a-number values are treated as less than other values).
func Float32sAreSorted(a []float32) bool { return sort.IsSorted(Float32Slice(a)) }

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p RuneSlice) Sort() { sort.Sort(p) }

// Runes sorts a slice of runes in increasing order.
func Runes(a []rune) { sort.Sort(RuneSlice(a)) }

// RunesAreSorted tests whether a slice of runes is sorted in increasing order.
func RunesAreSorted(a []rune) bool { return sort.IsSorted(RuneSlice(a)) }

type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p UintSlice) Sort() { sort.Sort(p) }

// Uints sorts a slice of uints in increasing order.
func Uints(a []uint) { sort.Sort(UintSlice(a)) }

// UintsAreSorted tests whether a slice of uints is sorted in increasing order.
func UintsAreSorted(a []uint) bool { return sort.IsSorted(UintSlice(a)) }

type Uint8Slice []uint8

func (p Uint8Slice) Len() int           { return len(p) }
func (p Uint8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint8Slice) Sort() { sort.Sort(p) }

// Uint8s sorts a slice of uint8s in increasing order.
func Uint8s(a []uint8) { sort.Sort(Uint8Slice(a)) }

// Uint8sAreSorted tests whether a slice of uint8s is sorted in increasing order.
func Uint8sAreSorted(a []uint8) bool { return sort.IsSorted(Uint8Slice(a)) }

type Uint16Slice []uint16

func (p Uint16Slice) Len() int           { return len(p) }
func (p Uint16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint16Slice) Sort() { sort.Sort(p) }

// Uint16s sorts a slice of uint16s in increasing order.
func Uint16s(a []uint16) { sort.Sort(Uint16Slice(a)) }

// Uint16sAreSorted tests whether a slice of uint16s is sorted in increasing order.
func Uint16sAreSorted(a []uint16) bool { return sort.IsSorted(Uint16Slice(a)) }

type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint32Slice) Sort() { sort.Sort(p) }

// Uint32s sorts a slice of uint32s in increasing order.
func Uint32s(a []uint32) { sort.Sort(Uint32Slice(a)) }

// Uint32sAreSorted tests whether a slice of uint32s is sorted in increasing order.
func Uint32sAreSorted(a []uint32) bool { return sort.IsSorted(Uint32Slice(a)) }

type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p Uint64Slice) Sort() { sort.Sort(p) }

// Uint64s sorts a slice of uint64s in increasing order.
func Uint64s(a []uint64) { sort.Sort(Uint64Slice(a)) }

// Uint64sAreSorted tests whether a slice of uint64s is sorted in increasing order.
func Uint64sAreSorted(a []uint64) bool { return sort.IsSorted(Uint64Slice(a)) }

// TODO:verify

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float32) bool { return f != f }
