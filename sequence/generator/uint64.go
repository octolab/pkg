package generator

import "sync/atomic"

// Uint64 provides functionality to produce increasing sequence of numbers.
//
//  uid, ids := new(generator.Uint64).At(7), make([]entity.ID, 4)
//
//  for i := range sequence.Simple(len(ids)) {
//  	ids[i] = entity.ID(uid.Next())
//  }
//
type Uint64 uint64

// At sets the Uint64 to the new position.
func (generator *Uint64) At(position uint64) *Uint64 {
	atomic.StoreUint64((*uint64)(generator), position)
	return generator
}

// Current returns a current value of the Uint64.
func (generator *Uint64) Current() uint64 {
	return atomic.LoadUint64((*uint64)(generator))
}

// Jump moves the Uint64 forward at the specified distance.
func (generator *Uint64) Jump(distance uint64) uint64 {
	return atomic.AddUint64((*uint64)(generator), distance)
}

// Next moves the Uint64 one step forward.
func (generator *Uint64) Next() uint64 {
	return atomic.AddUint64((*uint64)(generator), 1)
}

// Reset returns a current value of the Uint64 and resets it.
func (generator *Uint64) Reset() uint64 {
	return atomic.SwapUint64((*uint64)(generator), 0)
}
