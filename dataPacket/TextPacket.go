package dataPacket

const (
	TextPacket_TextTypeRaw = iota
	TextPacket_TextTypeChat
	TextPacket_TextTypeTranslation
	TextPacket_TextTypePopup
	TextPacket_TextTypeJukeboxPopup
	TextPacket_TextTypeTip
	TextPacket_TextTypeSystem
	TextPacket_TextTypeWhisper
	TextPacket_TextTypeAnnouncement
	TextPacket_TextTypeObject
	TextPacket_TextTypeObjectWhisper
)

// Text is sent by the client to the server to send chat messages, and by the server to the client to forward
// or send messages, which may be chat, popups, tips etc.
type TextPacket struct {
	PacketName string `json:"PacketName"`
	// TextType is the type of the text sent. When a client sends this to the server, it should always be
	// TextTypeChat. If the server sends it, it may be one of the other text types above.
	TextType byte `json:"TextType"`
	// SourceName is the name of the source of the messages. This source is displayed in text types such as
	// the TextTypeChat and TextTypeWhisper, where typically the username is shown.
	SourceName string `json:"SourceName"`
	// Message is the message of the packet. This field is set for each TextType and is the main component of
	// the packet.
	Message string `json:"Message"`
	// Parameters is a list of parameters that should be filled into the message. These parameters are only
	// written if the type of the packet is TextTypeTip, TextTypePopup or TextTypeJukeboxPopup.
	Parameters []string
}

// ID ...
func (*TextPacket) ID() byte {
	return IDTextPacket
}

// Marshal ...
func (pk *TextPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)

	w.Uint8(&pk.TextType)
	switch pk.TextType {
	case TextPacket_TextTypeChat, TextPacket_TextTypeWhisper, TextPacket_TextTypeAnnouncement:
		w.String(&pk.SourceName)
		w.String(&pk.Message)
	case TextPacket_TextTypeRaw, TextPacket_TextTypeTip, TextPacket_TextTypeSystem, TextPacket_TextTypeObject, TextPacket_TextTypeObjectWhisper:
		w.String(&pk.SourceName)
		w.String(&pk.Message)
	case TextPacket_TextTypeTranslation, TextPacket_TextTypePopup, TextPacket_TextTypeJukeboxPopup:
		var length uint32

		w.String(&pk.Message)
		w.Varuint32(&length)
		pk.Parameters = make([]string, length)
		for i := uint32(0); i < length; i++ {
			w.String(&pk.Parameters[i])
		}
	}
}

// Unmarshal ...
func (pk *TextPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)

	r.Uint8(&pk.TextType)
	switch pk.TextType {
 case TextPacket_TextTypeChat, TextPacket_TextTypeWhisper, TextPacket_TextTypeAnnouncement:
		r.String(&pk.SourceName)
		r.String(&pk.Message)
	case TextPacket_TextTypeRaw, TextPacket_TextTypeTip, TextPacket_TextTypeSystem, TextPacket_TextTypeObject, TextPacket_TextTypeObjectWhisper:
		r.String(&pk.SourceName)
		r.String(&pk.Message)
	case TextPacket_TextTypeTranslation, TextPacket_TextTypePopup, TextPacket_TextTypeJukeboxPopup:
		var length uint32

		r.String(&pk.Message)
		r.Varuint32(&length)
		pk.Parameters = make([]string, length)
		for i := uint32(0); i < length; i++ {
			r.String(&pk.Parameters[i])
		}
	}
}
