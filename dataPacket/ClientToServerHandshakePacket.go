package dataPacket

type ClientToServerHandshakePacket struct {
}

// ID ...
func (*ClientToServerHandshakePacket) ID() byte {
	return IDClientToServerHandshakePacket
}

// Marshal ...
func (pk *ClientToServerHandshakePacket) Marshal(w *PacketWriter) {
}

// Unmarshal ...
func (pk *ClientToServerHandshakePacket) Unmarshal(r *PacketReader) {
}
