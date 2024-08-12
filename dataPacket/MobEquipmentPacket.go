package dataPacket

import "goproxy/item"

type MobEquipmentPacket struct {
	PacketName   string `json:"PacketName"`
	Eid          int32
	Item         *item.Item
	Slot         byte
	SelectedSlot byte
	WindowID     byte
}

// ID ...
func (*MobEquipmentPacket) ID() byte {
	return IDMobEquipmentPacket
}

// Marshal ...
func (pk *MobEquipmentPacket) Marshal(w *PacketWriter) {
	w.Varint32(&pk.Eid)
	w.Item(pk.Item)
	w.Uint8(&pk.Slot)
	w.Uint8(&pk.SelectedSlot)
	w.Uint8(&pk.WindowID)
}

// Unmarshal ...
func (pk *MobEquipmentPacket) Unmarshal(r *PacketReader) {
	r.Varint32(&pk.Eid)
	pk.Item = r.Item()
	r.Uint8(&pk.Slot)
	r.Uint8(&pk.SelectedSlot)
	r.Uint8(&pk.WindowID)
}
