package dataPacket

const (
	PlayStatusPacket_LOGIN_SUCCESS = 0
	PlayStatusPacket_LOGIN_FAILED_CLIENT         = 1
	PlayStatusPacket_LOGIN_FAILED_SERVER         = 2
	PlayStatusPacket_PLAYER_SPAWN                = 3
	PlayStatusPacket_LOGIN_FAILED_INVALID_TENANT = 4
	PlayStatusPacket_LOGIN_FAILED_VANILLA_EDU    = 5
	PlayStatusPacket_LOGIN_FAILED_EDU_VANILLA    = 6
)

type PlayStatusPacket struct {
	Status int32 `json:"Status"`
	PacketName string `json:"PacketName"`
}

// ID ...
func (*PlayStatusPacket) ID() byte {
	return IDPlayStatusPacket
}

// Marshal ...
func (pk *PlayStatusPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.BEInt32(&pk.Status)
}

// Unmarshal ...
func (pk *PlayStatusPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.BEInt32(&pk.Status)
}
