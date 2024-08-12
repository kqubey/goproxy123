package dataPacket

type EntityMovePacket struct {
	PacketName string `json:"PacketName"`
	Eid        int32
	X          float32
	Y          float32
	Z          float32
	Yaw        float32
	HeadYaw    float32
	Pitch      float32
	OnGround   bool
	Teleported bool
}

// ID ...
func (*EntityMovePacket) ID() byte {
	return IDEntityMovePacket
}

// Marshal ...
func (pk *EntityMovePacket) Marshal(w *PacketWriter) {
	w.Varint32(&pk.Eid)
	w.Float32(&pk.X)
	w.Float32(&pk.Y)
	w.Float32(&pk.Z)
	w.ByteRotation(&pk.Pitch)
	w.ByteRotation(&pk.Yaw)
	w.ByteRotation(&pk.HeadYaw)
	w.Bool(&pk.OnGround)
	w.Bool(&pk.Teleported)
}

// Unmarshal ...
func (pk *EntityMovePacket) Unmarshal(r *PacketReader) {
	r.Varint32(&pk.Eid)
	r.Float32(&pk.X)
	r.Float32(&pk.Y)
	r.Float32(&pk.Z)
	r.ByteRotation(&pk.Pitch)
	r.ByteRotation(&pk.Yaw)
	r.ByteRotation(&pk.HeadYaw)
	r.Bool(&pk.OnGround)
	r.Bool(&pk.Teleported)
}
