package dataPacket

import (
	"goproxy/resourcePack"
)

type ResourcePacksInfoPacket struct {
	MustAccept bool   `json:"MustAccept"`
	PacketName string `json:"PacketName"`

	BehaviorPackStack []*resourcePack.ResourcePackInfoEntry
	ResourcePackStack []*resourcePack.ResourcePackInfoEntry
}

// ID ...
func (*ResourcePacksInfoPacket) ID() byte {
	return IDResourcePacksInfoPacket
}

// Marshal ...
func (pk *ResourcePacksInfoPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
	w.Bool(&pk.MustAccept)

	bhpcount := int16(len(pk.BehaviorPackStack))
	w.Int16(&bhpcount)
	for _, bhp := range pk.BehaviorPackStack {
		if bhpcount <= 0 {
			break
		}
		bhpcount--
		w.String(&bhp.PackID)
		w.String(&bhp.Version)
		w.Int64(&bhp.PackSize)
	}

	rpcount := int16(len(pk.ResourcePackStack))
	w.Int16(&rpcount)
	for _, rhp := range pk.ResourcePackStack {
		if rpcount <= 0 {
			break
		}
		rpcount--
		w.String(&rhp.PackID)
		w.String(&rhp.Version)
		w.Int64(&rhp.PackSize)
	}
}

// Unmarshal ...
func (pk *ResourcePacksInfoPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	r.Bool(&pk.MustAccept)

	var bhpcount int16
	r.Int16(&bhpcount)
	for {

		if bhpcount <= 0 {
			break
		}
		bhpcount--
		var packid string
		r.String(&packid)
		var version string
		r.String(&version)
		var psize int64
		r.Int64(&psize)
		pk.BehaviorPackStack = append(pk.BehaviorPackStack, resourcePack.NewResourcePackInfoEntry(packid, version, psize))
	}

	var rpcount int16
	r.Int16(&rpcount)
	for {

		if rpcount <= 0 {
			break
		}
		rpcount--
		var packid string
		r.String(&packid)

		var version string
		r.String(&version)
		var psize int64
		r.Int64(&psize)
		pk.ResourcePackStack = append(pk.ResourcePackStack, resourcePack.NewResourcePackInfoEntry(packid, version, psize))
	}
	//todo
}
