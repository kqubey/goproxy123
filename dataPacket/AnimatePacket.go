package dataPacket

const (
	 AnimatePacket_ACTION_SWING_ARM = 1

	 AnimatePacket_ACTION_STOP_SLEEP = 3
	 AnimatePacket_ACTION_CRITICAL_HIT = 4
	 AnimatePacket_ACTION_ROW_RIGHT = 128
	 AnimatePacket_ACTION_ROW_LEFT = 129

)

type AnimatePacket struct {
	PacketName string `json:"PacketName"`
 Action int32 `json:"Action"`
 Eid int32 `json:"Eid"`
 Float float32 `json:"Float"`
}

// ID ...
func (*AnimatePacket) ID() byte {
	return IDAnimatePacket
}

// Marshal ...
func (pk *AnimatePacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
 w.Varint32(&pk.Action)
 w.Varint32(&pk.Eid)
 if ((int64(pk.Float)) & 0x80) > 0 {
   w.Float32(&pk.Float)
 }

}

// Unmarshal ...
func (pk *AnimatePacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.Action)
 r.Varint32(&pk.Eid)
 pk.Float = 0.0
}
