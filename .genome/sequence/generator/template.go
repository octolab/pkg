// +build genome

package generator

import "sync/atomic"

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

// T provides functionality to produce increasing sequence of numbers.
//
//  uid, ids := new(generator.T).At(7), make([]entity.ID, 4)
//
//  for i := range sequence.Simple(len(ids)) {
//  	ids[i] = entity.ID(uid.Next())
//  }
//
type T uint64

// At sets the T to the new position.
func (generator *T) At(position uint64) *T {
	atomic.StoreUint64((*uint64)(generator), position)
	return generator
}

// Current returns a current value of the T.
func (generator *T) Current() uint64 {
	return atomic.LoadUint64((*uint64)(generator))
}

// Jump moves the T forward at the specified distance.
func (generator *T) Jump(distance uint64) uint64 {
	return atomic.AddUint64((*uint64)(generator), distance)
}

// Next moves the T one step forward.
func (generator *T) Next() uint64 {
	return atomic.AddUint64((*uint64)(generator), 1)
}

// Reset returns a current value of the T and resets it.
func (generator *T) Reset() uint64 {
	return atomic.SwapUint64((*uint64)(generator), 0)
}
