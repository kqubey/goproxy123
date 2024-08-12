package dataPacket

import "goproxy/item"

type MobArmorEquipmentPacket struct {
	Eid   int32
	Slots [4]*item.Item
}

// ID ...
func (*MobArmorEquipmentPacket) ID() byte {
	return IDMobArmorEquipmentPacket
}

// Marshal ...
func (pk *MobArmorEquipmentPacket) Marshal(w *PacketWriter) {
	w.Varint32(&pk.Eid)
	for _, itemd := range pk.Slots {
		w.Item(itemd)
	}
}

// Unmarshal ...
func (pk *MobArmorEquipmentPacket) Unmarshal(r *PacketReader) {
	r.Varint32(&pk.Eid)
	for i, _ := range pk.Slots {
		pk.Slots[i] = r.Item()
	}
}
