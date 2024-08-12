package dataPacket

import (
	"goproxy/nbt"
)

type BlockEntityDataPacket struct {
	X   int32
	Y   uint32
	Z   int32
	NBT map[string]interface{}
}

// ID ...
func (*BlockEntityDataPacket) ID() byte {
	return IDBlockEntityDataPacket
}

// Marshal ...
func (pk *BlockEntityDataPacket) Marshal(w *PacketWriter) {
	w.BlockCoords(&pk.X, &pk.Y, &pk.Z)
	w.NBT(&pk.NBT, nbt.NetworkLittleEndian)
}

// Unmarshal ...
func (pk *BlockEntityDataPacket) Unmarshal(r *PacketReader) {
	pk.NBT = make(map[string]interface{})
	r.BlockCoords(&pk.X, &pk.Y, &pk.Z)
	//var bts []byte
	//r.Bytes(&bts)
	//log.Println(nbt.Dump(bts, nbt.LittleEndian))
	r.NBT(&pk.NBT, nbt.NetworkLittleEndian)
}
