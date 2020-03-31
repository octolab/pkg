package generator

import (
	"encoding/hex"

	"github.com/google/uuid"
)

// UUID provides functionality to produce a new random
// Universal Unique Identifier as defined in RFC 4122.
//
//  uuid, ids := new(generator.UUID), make([]entity.ID, 4)
//  for i := range sequence.Simple(len(ids)) {
//  	ids[i] = entity.ID(uid.Next())
//  }
//
type UUID [16]byte

// Next returns a new random UUID which may or may not be valid.
func (generator UUID) Next() UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		return generator
	}
	return UUID(id)
}

// String returns the string form of the UUID,
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, or "" if the UUID is invalid.
func (generator UUID) String() string {
	if !generator.Valid() {
		return ""
	}
	var buf [36]byte
	dst := buf[:]
	hex.Encode(dst, generator[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], generator[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], generator[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], generator[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], generator[10:])
	return string(buf[:])
}

// Valid returns true if the UUID is valid.
func (generator UUID) Valid() bool {
	return generator != [16]byte{}
}
