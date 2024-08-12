package dataPacket

const PlayerActionPacket_ACTION_START_BREAK = 0
const PlayerActionPacket_ACTION_ABORT_BREAK = 1
const PlayerActionPacket_ACTION_STOP_BREAK = 2

const PlayerActionPacket_ACTION_RELEASE_ITEM = 5
const PlayerActionPacket_ACTION_STOP_SLEEPING = 6
const PlayerActionPacket_ACTION_SPAWN_SAME_DIMENSION = 7
const PlayerActionPacket_ACTION_JUMP = 8
const PlayerActionPacket_ACTION_START_SPRINT = 9
const PlayerActionPacket_ACTION_STOP_SPRINT = 10
const PlayerActionPacket_ACTION_START_SNEAK = 11
const PlayerActionPacket_ACTION_STOP_SNEAK = 12
const PlayerActionPacket_ACTION_SPAWN_OVERWORLD = 13
const PlayerActionPacket_ACTION_SPAWN_NETHER = 14
const PlayerActionPacket_ACTION_START_GLIDE = 15
const PlayerActionPacket_ACTION_STOP_GLIDE = 16

const PlayerActionPacket_ACTION_BUILD_DENIED = 17

const PlayerActionPacket_ACTION_CONTINUE_BREAK = 18

type PlayerActionPacket struct {
	/*
	    public $eid;
	   	public $action;
	   	public $x;
	   	public $y;
	   	public $z;
	   	public $face;
	*/
	Eid        int32
	Action     int32
	X          int32
	Y          uint32
	Z          int32
	Face       int32
	PacketName string `json:"PacketName"`
}

// ID ...
func (*PlayerActionPacket) ID() byte {
	return IDPlayerActionPacket
}

// Marshal ...
func (pk *PlayerActionPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Varint32(&pk.Eid)
	w.Varint32(&pk.Action)
	w.BlockCoords(&pk.X, &pk.Y, &pk.Z)
	w.Varint32(&pk.Face)
}

// Unmarshal ...
func (pk *PlayerActionPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Varint32(&pk.Eid)
	r.Varint32(&pk.Action)
	r.BlockCoords(&pk.X, &pk.Y, &pk.Z)
	r.Varint32(&pk.Face)
}
