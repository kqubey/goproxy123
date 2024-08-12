package dataPacket

type SetEntityDataPacket struct {
	Eid        int32                  `json:"Eid"`
	Metadata   map[uint32]interface{} `json:"Metadata"`
	PacketName string                 `json:"PacketName"`
}

// ID ...
func (*SetEntityDataPacket) ID() byte {
	return IDSetEntityDataPacket
}

// Marshal ...
func (pk *SetEntityDataPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.Eid)
	w.EntityMetadata(&pk.Metadata)
}

// Unmarshal ...
func (pk *SetEntityDataPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.Eid)
	r.EntityMetadata(&pk.Metadata)
}
