package dataPacket



type RespawnPacket struct {
	X float32 `json:"X"`
	Y float32 `json:"Y"`
	Z float32 `json:"Z"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*RespawnPacket) ID() byte {
	return IDRespawnPacket
}

// Marshal ...
func (pk *RespawnPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Float32(&pk.X)
	w.Float32(&pk.Y)
	w.Float32(&pk.Z)
}

// Unmarshal ...
func (pk *RespawnPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Float32(&pk.X)
	r.Float32(&pk.Y)
	r.Float32(&pk.Z)
}
