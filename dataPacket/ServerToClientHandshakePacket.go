package dataPacket

type ServerToClientHandshakePacket struct {
	PublicKey   string
	ServerToken string
}

// ID ...
func (*ServerToClientHandshakePacket) ID() byte {
	return IDServerToClientHandshakePacket
}

// Marshal ...
func (pk *ServerToClientHandshakePacket) Marshal(w *PacketWriter) {

}

// Unmarshal ...
func (pk *ServerToClientHandshakePacket) Unmarshal(r *PacketReader) {
	r.String(&pk.PublicKey)
	r.String(&pk.ServerToken)
}
