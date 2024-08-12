package dataPacket

import "log"

// Text is sent by the client to the server to send chat messages, and by the server to the client to forward
// or send messages, which may be chat, popups, tips etc.
type LoginPacket struct {
	ClientProtocol int32
	// ConnectionRequest is a string containing information about the player and JWTs that may be used to
	// verify if the player is connected to XBOX Live. The connection request also contains the necessary
	// client public key to initiate encryption.
	ConnectionRequest []byte
}

// ID ...
func (*LoginPacket) ID() byte {
	return IDLoginPacket
}

// Marshal ...
func (pk *LoginPacket) Marshal(w *PacketWriter) {
	w.BEInt32(&pk.ClientProtocol)
	ged := uint8(0)
	w.Uint8(&ged)
	w.ByteSlice(&pk.ConnectionRequest)
}

// Unmarshal ...
func (pk *LoginPacket) Unmarshal(r *PacketReader) {
	r.BEInt32(&pk.ClientProtocol)
	ged := uint8(0)
	r.Uint8(&ged)
	log.Println("login pk: ged", ged, "proto:", pk.ClientProtocol)
	r.ByteSlice(&pk.ConnectionRequest)
}
