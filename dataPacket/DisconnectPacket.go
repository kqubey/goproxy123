package dataPacket

type DisconnectPacket struct {
	HideDisconnectionScreen bool `json:"HideDisconnectionScreen"`
	Message string `json:"Message"`
	PacketName string `json:"PacketName"`
}
// ID ...
func (*DisconnectPacket) ID() byte {
	return IDDisconnectPacket
}

// Marshal ...
func (pk *DisconnectPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Bool(&pk.HideDisconnectionScreen)
	w.String(&pk.Message)
}

// Unmarshal ...
func (pk *DisconnectPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Bool(&pk.HideDisconnectionScreen)
	r.String(&pk.Message)
}
