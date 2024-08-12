package dataPacket



type RiderJumpPacket struct {
	Unknown int32
	PacketName string `json:"PacketName"`
}

// ID ...
func (*RiderJumpPacket) ID() byte {
	return IDRiderJumpPacket
}

// Marshal ...
func (pk *RiderJumpPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.Unknown)
}

// Unmarshal ...
func (pk *RiderJumpPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.Unknown)
}
