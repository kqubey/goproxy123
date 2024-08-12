package dataPacket

const UpdateBlockPacket_FLAG_NONE = 0b0000
const UpdateBlockPacket_FLAG_NEIGHBORS = 0b0001
const UpdateBlockPacket_FLAG_NETWORK = 0b0010
const UpdateBlockPacket_FLAG_NOGRAPHIC = 0b0100
const UpdateBlockPacket_FLAG_PRIORITY = 0b1000

type UpdateBlockPacket struct {
	PacketName string `json:"PacketName"`
	X          int32
	Y          uint32
	Z          int32
	BlockID    uint32
	Flags      uint32
}

/*
public $x;
	public $z;
	public $y;
	public $blockId;
	public $blockData;
	public $flags;
*/
// ID ...
func (*UpdateBlockPacket) ID() byte {
	return IDUpdateBlockPacket
}

// Marshal ...
func (pk *UpdateBlockPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.BlockCoords(&pk.X, &pk.Y, &pk.Z)
	w.Varuint32(&pk.BlockID)
	w.Varuint32(&pk.Flags)
}

// Unmarshal ...
func (pk *UpdateBlockPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.BlockCoords(&pk.X, &pk.Y, &pk.Z)
	r.Varuint32(&pk.BlockID)
	r.Varuint32(&pk.Flags)
}
