package generator

import "sync/atomic"

// Uint32 provides functionality to produce increasing sequence of numbers.
//
//	uid, ids := new(generator.Uint32).At(7), make([]entity.ID, 4)
//
//	for i := range sequence.Simple(len(ids)) {
//		ids[i] = entity.ID(uid.Next())
//	}
type Uint32 uint32

// At sets the Uint32 to the new position.
func (generator *Uint32) At(position uint32) *Uint32 {
	atomic.StoreUint32((*uint32)(generator), position)
	return generator
}

// Current returns a current value of the Uint32.
func (generator *Uint32) Current() uint32 {
	return atomic.LoadUint32((*uint32)(generator))
}

// Jump moves the Uint32 forward at the specified distance.
func (generator *Uint32) Jump(distance uint32) uint32 {
	return atomic.AddUint32((*uint32)(generator), distance)
}

// Next moves the Uint32 one step forward.
func (generator *Uint32) Next() uint32 {
	return atomic.AddUint32((*uint32)(generator), 1)
}

// Reset returns a current value of the Uint32 and resets it.
func (generator *Uint32) Reset() uint32 {
	return atomic.SwapUint32((*uint32)(generator), 0)
}
