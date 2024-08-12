package dataPacket


type ResourcePackChunkRequestPacket struct {
	PacketName string `json:"PacketName"`
	PackID string
	ChunkIndex int32
}

// ID ...
func (*ResourcePackChunkRequestPacket) ID() byte {
	return IDResourcePackChunkRequestPacket
}

// Marshal ...
func (pk *ResourcePackChunkRequestPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.String(&pk.PackID)
	w.Int32(&pk.ChunkIndex)
}

// Unmarshal ...
func (pk *ResourcePackChunkRequestPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)

}
