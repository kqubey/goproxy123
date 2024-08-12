package dataPacket

type SetEntityMotionPacket struct {
	PacketName string `json:"PacketName"`
	Eid        int32
	MotionX    float32
	MotionY    float32
	MotionZ    float32
}

// ID ...
func (*SetEntityMotionPacket) ID() byte {
	return IDSetEntityMotionPacket
}

// Marshal ...
func (pk *SetEntityMotionPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.Eid)
	w.Float32(&pk.MotionX)
	w.Float32(&pk.MotionY)
	w.Float32(&pk.MotionZ)
}

// Unmarshal ...
func (pk *SetEntityMotionPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.Eid)
	r.Float32(&pk.MotionX)
	r.Float32(&pk.MotionY)
	r.Float32(&pk.MotionZ)
}
