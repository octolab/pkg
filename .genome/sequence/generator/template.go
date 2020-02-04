// +build genome

package sequence

import "sync/atomic"

type T uint64

type RelationT struct {
	Add   interface{}
	Load  interface{}
	Store interface{}
	Swap  interface{}
}

var _ = map[string]RelationT{
	"uin32": {
		Add:   atomic.AddUint32,
		Load:  atomic.LoadUint32,
		Store: atomic.StoreUint32,
		Swap:  atomic.SwapUint32,
	},
	"uin64": {
		Add:   atomic.AddUint64,
		Load:  atomic.LoadUint64,
		Store: atomic.StoreUint64,
		Swap:  atomic.SwapUint64,
	},
}

// GeneratorT provides functionality to produce increasing sequence of numbers.
//
//  generator, sequence := new(sequence.GeneratorT).At(7), make([]entity.ID, 4)
//
//  for i := range sequence.Simple(len(sequence)) {
//  	sequence[i] = entity.ID(generator.Next())
//  }
//
type GeneratorT T

// At sets the GeneratorT to the new position.
func (generator *GeneratorT) At(position uint64) *GeneratorT {
	atomic.StoreUint64((*uint64)(generator), position)
	return generator
}

// Current returns a current value of the GeneratorT.
func (generator *GeneratorT) Current() uint64 {
	return atomic.LoadUint64((*uint64)(generator))
}

// Jump moves the GeneratorT forward at the specified distance.
func (generator *GeneratorT) Jump(distance uint64) uint64 {
	return atomic.AddUint64((*uint64)(generator), distance)
}

// Next moves the GeneratorT one step forward.
func (generator *GeneratorT) Next() uint64 {
	return atomic.AddUint64((*uint64)(generator), 1)
}

// Reset returns a current value of the GeneratorT and resets it.
func (generator *GeneratorT) Reset() uint64 {
	return atomic.SwapUint64((*uint64)(generator), 0)
}
