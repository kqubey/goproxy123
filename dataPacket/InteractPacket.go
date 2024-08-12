package dataPacket

const (
	InteractPacket_ACTION_RIGHT_CLICK   = 1
	InteractPacket_ACTION_LEFT_CLICK    = 2
	InteractPacket_ACTION_LEAVE_VEHICLE = 3
	InteractPacket_ACTION_MOUSEOVER     = 4
	//todo rename all consts in packets
)

type InteractPacket struct {
	Action     byte   `json:"Action"`
	Target     int32  `json:"Target"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*InteractPacket) ID() byte {
	return IDInteractPacket
}

// Marshal ...
func (pk *InteractPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Uint8(&pk.Action)
	w.Varint32(&pk.Target)
}

// Unmarshal ...
func (pk *InteractPacket) Unmarshal(r *PacketReader) {
	r.Uint8(&pk.Action)
	r.Varint32(&pk.Target)
}
