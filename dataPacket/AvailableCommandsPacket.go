package dataPacket



type AvailableCommandsPacket struct {
	Commands   string `json:"Commands"`
	Unknown    string `json:"Unknown"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*AvailableCommandsPacket) ID() byte {
	return IDAvailableCommandsPacket
}

// Marshal ...
func (pk *AvailableCommandsPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
}

// Unmarshal ...
func (pk *AvailableCommandsPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.String(&pk.Commands)
	r.String(&pk.Unknown)
}
