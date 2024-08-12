package dataPacket

import (
	"github.com/google/uuid"
	"goproxy/item"
	"log"
)

type CraftingEventPacket struct {
	WindowID byte
	Type     int32
	UUID     uuid.UUID
	Input    []item.Item
	Output   []item.Item
}

// ID ...
func (*CraftingEventPacket) ID() byte {
	return IDCraftingEventPacket
}

// Marshal ...
func (pk *CraftingEventPacket) Marshal(w *PacketWriter) {
	w.Uint8(&pk.WindowID)
	w.Varint32(&pk.Type)
	w.UUID(&pk.UUID)

	var size = uint32(len(pk.Input))
	log.Println("size:", size)
	w.Varuint32(&size)
	for i := uint32(0); i < size; i++ {
		w.Item(&pk.Input[i])
	}
	size = uint32(len(pk.Output))
	log.Println("size:", size)
	w.Varuint32(&size)
	for i := uint32(0); i < size; i++ {
		w.Item(&pk.Output[i])
	}
}

// Unmarshal ...
func (pk *CraftingEventPacket) Unmarshal(r *PacketReader) {
	r.Uint8(&pk.WindowID)
	r.Varint32(&pk.Type)
	r.UUID(&pk.UUID)

	var size uint32
	r.Varuint32(&size)
	log.Println("readed size", size)
	for i := uint32(0); i < size; i++ {
		pk.Input = append(pk.Input, *r.Item())
	}
	r.Varuint32(&size)
	for i := uint32(0); i < size; i++ {
		pk.Output = append(pk.Output, *r.Item())
	}
}
