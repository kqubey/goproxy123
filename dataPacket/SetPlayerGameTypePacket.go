package dataPacket

type SetPlayerGameTypePacket struct {
	Type int32
}

// ID ...
func (*SetPlayerGameTypePacket) ID() byte {
	return IDSetPlayerGameTypePacket
}

// Marshal ...
func (pk *SetPlayerGameTypePacket) Marshal(w *PacketWriter) {
	w.Varint32(&pk.Type)
}

// Unmarshal ...
func (pk *SetPlayerGameTypePacket) Unmarshal(r *PacketReader) {
	r.Varint32(&pk.Type)
}
