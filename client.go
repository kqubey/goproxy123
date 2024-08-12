package main

import (
	"bytes"
	"github.com/DataDog/czlib"
	"goproxy/dataPacket"
	"goproxy/raklib"
	"goproxy/utils"
	"io"
	"log"
)

type ClientConnection struct {
	Listener   *raklib.Listener
	Connection *raklib.Conn
	IsShutdown chan bool
	IsOnline   bool
}

func NewConnectionToClient() *ClientConnection {
	listener, _ := raklib.Listen("0.0.0.0:19132")
	listener.PongData([]byte("MCPE;§bAlacrium §aConnector;1;1.1;0;10;"))
	Listener = listener
	nconn, _ := listener.Accept()
	return &ClientConnection{Listener: listener, Connection: nconn.(*raklib.Conn), IsShutdown: make(chan bool), IsOnline: true}
}

func (*ClientConnection) DecodeBatch(buffer []byte) [][]byte {
	var out bytes.Buffer
	if r, err := czlib.NewReader(bytes.NewReader(buffer)); err != nil {
		log.Println(err)
		return nil
	} else {
		_, _ = io.Copy(&out, r)
		var packets [][]byte
		for out.Len() != 0 {
			packets = append(packets, []byte(utils.ReadString(&out)))
		}
		return packets
	}
}

func (ccon *ClientConnection) SendPacketsRaw(pks [][]byte) {
	_, _ = ccon.Connection.Write(ccon.EncodeBatch(pks))
}

func (ccon *ClientConnection) SendPacket(pk dataPacket.DataPacket) {
	buff := bytes.NewBuffer(nil)
	buff.WriteByte(pk.ID())
	pk.Marshal(dataPacket.NewWriter(buff, 2))

	_, _ = ccon.Connection.Write(ccon.EncodeBatch([][]byte{buff.Bytes()}))
}

func (ccon *ClientConnection) EncodeBatch(pks [][]byte) []byte {
	var b bytes.Buffer
	zw, _ := czlib.NewWriterLevel(&b, 7)
	l := make([]byte, 5)

	for _, pk := range pks {
		_ = utils.WriteVarUInt32(zw, uint32(len(pk)), l)
		_, _ = zw.Write(pk)
	}
	_ = zw.Close()

	var bb bytes.Buffer
	bb.WriteByte(0xfe)
	bb.Write(b.Bytes())
	return bb.Bytes()
}
