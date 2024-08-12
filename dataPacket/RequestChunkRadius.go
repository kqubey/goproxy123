package dataPacket



type RequestChunkRadiusPacket struct {
	Radius int32 `json:"Radius"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*RequestChunkRadiusPacket) ID() byte {
	return IDRequestChunkRadiusPacket
}

// Marshal ...
func (pk *RequestChunkRadiusPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.Radius)
}

// Unmarshal ...
func (pk *RequestChunkRadiusPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.Radius)
}
