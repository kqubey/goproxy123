package dataPacket

import uuid2 "github.com/google/uuid"

const (
	PlayerListPacket_TYPE_ADD = 0
	PlayerListPacket_TYPE_REMOVE = 1
)
type PlayerListPacket struct {
	Entries map[string][]interface{} `json:"Entries"`
	Type byte `json:"Type"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*PlayerListPacket) ID() byte {
	return IDPlayerListPacket
}

// Marshal ...
func (pk *PlayerListPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)

}

// Unmarshal ...
func (pk *PlayerListPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	pk.Entries = make(map[string][]interface{})

	r.Uint8(&pk.Type)
	var count uint32
	r.Varuint32(&count)
	for i := 0; i <= int(count); i++ {
		if pk.Type == PlayerListPacket_TYPE_ADD {
			var uuid uuid2.UUID
			r.UUID(&uuid)
			var eid int32
			r.Varint32(&eid)
			var s1 string
			var s2 string
			var s3 string
			r.String(&s1)
			r.String(&s2)
			r.String(&s3)
			pk.Entries[uuid.String()] = append(pk.Entries[uuid.String()], eid)
			pk.Entries[uuid.String()] = append(pk.Entries[uuid.String()], s1)
			pk.Entries[uuid.String()] = append(pk.Entries[uuid.String()], s2)
			pk.Entries[uuid.String()] = append(pk.Entries[uuid.String()], s3)
		} else {
			var uuid uuid2.UUID
			r.UUID(&uuid)
			pk.Entries[uuid.String()] = nil
		}

	}
}
