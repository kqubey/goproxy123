package dataPacket


type CommandStepPacket struct {
	PacketName string `json:"PacketName"`
 Command string
 Overload string
 Uvarint1 uint32
 CurrentStep uint32
 Done bool
 ClientID uint32
 InputJson string
 OutputJson string
}

// ID ...
func (*CommandStepPacket) ID() byte {
	return IDCommandStepPacket
}

// Marshal ...
func (pk *CommandStepPacket) Marshal(w *PacketWriter) {
	pk.PacketName = getName(pk)
 w.String(&pk.Command)
 w.String(&pk.Overload)
 w.Varuint32(&pk.Uvarint1)
 w.Varuint32(&pk.CurrentStep)
 w.Bool(&pk.Done)
 w.Varuint32(&pk.ClientID)
 w.String(&pk.InputJson)
 w.String(&pk.OutputJson)
 //todo remaining
}

// Unmarshal ...
func (pk *CommandStepPacket) Unmarshal(r *PacketReader) {
	pk.PacketName = getName(pk)
	
}
