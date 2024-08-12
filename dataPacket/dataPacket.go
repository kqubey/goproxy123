package dataPacket

import (
	"bytes"
	"fmt"
	"reflect"

	"io"
)

type PacketData struct {
	h       *Header
	full    []byte
	payload *bytes.Buffer
}

func getName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

// Packet represents a packet that may be sent over a Minecraft network connection. The packet needs to hold
// a method to encode itself to binary and decode itself from binary.
type DataPacket interface {
	// ID returns the ID of the packet. All of these identifiers of dataPacket may be found in id.go.
	ID() byte
	// Marshal encodes the packet to its binary representation into buf.
	Marshal(w *PacketWriter)
	// Unmarshal decodes a serialised packet in buf into the Packet instance. The serialised packet passed
	// into Unmarshal will not have a header in it.
	Unmarshal(r *PacketReader)
}

// Header is the header of a packet. It exists out of a single varuint32 which is composed of a packet ID and
// a sender and target sub client ID. These IDs are used for split screen functionality.
type Header struct {
	PacketID byte
}

// Write writes the header as a single varuint32 to buf.
func (header *Header) Write(w io.ByteWriter) error {
	return WriteByte(w, header.PacketID)
}

// Read reads a varuint32 from buf and sets the corresponding values to the Header.
func (header *Header) Read(r io.ByteReader) error {
	var value byte
	if err := ReadByte(r, &value); err != nil {
		return err
	}
	header.PacketID = value
	return nil
}

func ParseDataPacket(data []byte) *PacketData {
	buf := bytes.NewBuffer(data)
	hd, _ := buf.ReadByte()
	return &PacketData{h: &Header{PacketID: hd}, full: data, payload: buf}
}

func (p *PacketData) TryDecodePacket() (pk DataPacket, err error) {
	pkFunc, ok := RegisteredPackets[p.h.PacketID]
	r := NewReader(p.payload, 0) //todo shield
	if !ok {
		// No packet with the ID. This may be a custom packet of some sorts.
		pk = &UnknownPacket{PacketID: p.h.PacketID}
		pk.Unmarshal(r)
		if p.payload.Len() != 0 {
			fmt.Println("unread bytes trydecode")
			return pk, nil
		}
		return pk, nil
	}
	pk = pkFunc()

	defer func() {
		if recoveredErr := recover(); recoveredErr != nil {
			err = fmt.Errorf("%T: %w", pk, recoveredErr.(error))
		}
	}()
	pk.Unmarshal(r)
	if p.payload.Len() != 0 {
		fmt.Println("unread 2 trydecode")
		return pk, nil
	}
	return pk, nil
}
