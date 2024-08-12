package dataPacket

type TransferPacket struct {
	Address    string `json:"Address"`
	Port       uint16 `json:"Port"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*TransferPacket) ID() byte {
	return IDTransferPacket
}

// Marshal ...
func (pk *TransferPacket) Marshal(w *PacketWriter) {
	w.String(&pk.Address)
	w.Uint16(&pk.Port)
}

// Unmarshal ...
func (pk *TransferPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.String(&pk.Address)
	r.Uint16(&pk.Port)
}
