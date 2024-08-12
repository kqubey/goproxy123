package dataPacket

type FullChunkPacket struct {
	ChunkX     int32  `json:"ChunkX"`
	ChunkZ     int32  `json:"ChunkZ"`
	Data       string `json:"Data"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*FullChunkPacket) ID() byte {
	return IDFullChunkPacket
}

// Marshal ...
func (pk *FullChunkPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.ChunkX)
	w.Varint32(&pk.ChunkZ)
	w.String(&pk.Data)
}

// Unmarshal ...
func (pk *FullChunkPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.ChunkX)
	r.Varint32(&pk.ChunkZ)
	r.String(&pk.Data)
}
