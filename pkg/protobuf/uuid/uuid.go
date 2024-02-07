package uuid

import (
	"encoding/binary"

	"github.com/google/uuid"
)

func ToUUIDProto(u uuid.UUID) *Uuid {
	highBits := binary.BigEndian.Uint64(u[0:8])
	lowBits := binary.BigEndian.Uint64(u[8:16])

	return &Uuid{
		HighBits: highBits,
		LowBits:  lowBits,
	}
}

func FromUUIDProto(u *Uuid) uuid.UUID {
	var id uuid.UUID
	binary.BigEndian.PutUint64(id[0:8], u.HighBits)
	binary.BigEndian.PutUint64(id[8:16], u.LowBits)

	return id
}
