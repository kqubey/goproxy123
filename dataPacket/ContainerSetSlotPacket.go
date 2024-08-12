package dataPacket

import "goproxy/item"

type ContainerSetSlotPacket struct {
	WindowID   byte       `json:"WindowID"`
	Slot       int32      `json:"Slot"`
	HotbarSlot int32      `json:"HotbarSlot"`
	Item       *item.Item `json:"Item"`
	PacketName string     `json:"PacketName"`
}

// ID ...
func (*ContainerSetSlotPacket) ID() byte {
	return IDContainerSetSlotPacket
}

// Marshal ...
func (pk *ContainerSetSlotPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Uint8(&pk.WindowID)
	w.Varint32(&pk.Slot)
	w.Varint32(&pk.HotbarSlot)
	w.Item(pk.Item)
	var unk byte = 0x00
	w.Uint8(&unk)
}

// Unmarshal ...
func (pk *ContainerSetSlotPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Uint8(&pk.WindowID)
	r.Varint32(&pk.Slot)
	r.Varint32(&pk.HotbarSlot)
	pk.Item = r.Item()
	var unk byte
	r.Uint8(&unk)
}
