package dataPacket

const (

)

type ResourcePackDataInfoPacket struct {
	PackID string `json:"PackID"`
	MaxChunkSize int32 `json:"MaxChunkSize"`
	ChunkCount int32 `json:"ChunkCount"`
	CompressedPackSize int64 `json:"CompressedPackSize"`
	Sha512 string `json:"Sha512"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*ResourcePackDataInfoPacket) ID() byte {
	return IDResourcePackDataInfoPacket
}

// Marshal ...
func (pk *ResourcePackDataInfoPacket) Marshal(w *PacketWriter) {
	//todo
	pk.PacketName = getName(pk)
}

// Unmarshal ...
func (pk *ResourcePackDataInfoPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.String(&pk.PackID)
	r.Int32(&pk.MaxChunkSize)
	r.Int32(&pk.ChunkCount)
	r.Int64(&pk.CompressedPackSize)
	r.String(&pk.Sha512)
}
