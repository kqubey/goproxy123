package dataPacket

import (
	"fmt"
)

type UnknownPacket struct {
	// PacketID is the packet ID of the packet.
	PacketID byte
	// Payload is the raw payload of the packet.
	Payload []byte
	PacketName string `json:"PacketName"'`
}

// ID ...
func (pk *UnknownPacket) ID() byte {
	return pk.PacketID
}

// Marshal ...
func (pk *UnknownPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Bytes(&pk.Payload)
}

// Unmarshal ...
func (pk *UnknownPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Bytes(&pk.Payload)
}

// String implements a hex representation of an unknown packet, so that it is easier to read and identify
// unknown incoming dataPacket.
func (pk *UnknownPacket) String() string {
	return fmt.Sprintf("{ID:0x%x Payload:0x%x}", pk.PacketID, pk.Payload)
}
