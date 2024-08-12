package dataPacket


type ResourcePackChunkDataPacket struct {
	PacketName string `json:"PacketName"`
	PackID string
	ChunkIndex int32
	Progress int64
	Data string
}

// ID ...
func (*ResourcePackChunkDataPacket) ID() byte {
	return IDResourcePackChunkDataPacket
}

// Marshal ...
func (pk *ResourcePackChunkDataPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	//todo
}

// Unmarshal ...
func (pk *ResourcePackChunkDataPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.String(&pk.PackID)
	r.Int32(&pk.ChunkIndex)
	r.Int64(&pk.Progress)
	var len int32
	r.Int32(&len)
	bbuf := make([]byte, len)
	r.BytesLength(bbuf)
	pk.Data = string(bbuf)
}
