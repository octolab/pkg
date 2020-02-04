package sequence

import "sync/atomic"

// GeneratorUint64 provides functionality to produce increasing sequence of numbers.
//
//  generator, sequence := new(sequence.GeneratorUint64).At(7), make([]entity.ID, 4)
//
//  for i := range sequence.Simple(len(sequence)) {
//  	sequence[i] = entity.ID(generator.Next())
//  }
//
type GeneratorUint64 uint64

// At sets the GeneratorUint64 to the new position.
func (generator *GeneratorUint64) At(position uint64) *GeneratorUint64 {
	atomic.StoreUint64((*uint64)(generator), position)
	return generator
}

// Current returns a current value of the GeneratorUint64.
func (generator *GeneratorUint64) Current() uint64 {
	return atomic.LoadUint64((*uint64)(generator))
}

// Jump moves the GeneratorUint64 forward at the specified distance.
func (generator *GeneratorUint64) Jump(distance uint64) uint64 {
	return atomic.AddUint64((*uint64)(generator), distance)
}

// Next moves the GeneratorUint64 one step forward.
func (generator *GeneratorUint64) Next() uint64 {
	return atomic.AddUint64((*uint64)(generator), 1)
}

// Reset returns a current value of the GeneratorUint64 and resets it.
func (generator *GeneratorUint64) Reset() uint64 {
	return atomic.SwapUint64((*uint64)(generator), 0)
}
