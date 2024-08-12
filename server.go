package main

import (
	"bytes"
	"github.com/DataDog/czlib"
	"goproxy/dataPacket"
	"goproxy/raklib"
	"goproxy/utils"
	"io"
	"time"
)

type ServerConnection struct {
	Connection *raklib.Conn
	IsShutdown chan bool
	IsOnline   bool
}

func NewConnectionToServer(address string) (*ServerConnection, error) {
	var errd error
	conn, err := raklib.DialTimeout(address, time.Second*10)
	if err == nil {
		if errd != nil {
			err = errd
		}
	}

	sc := &ServerConnection{Connection: conn, IsShutdown: make(chan bool), IsOnline: true}
	return sc, err
}

func (*ServerConnection) DecodeBatch(buffer []byte) [][]byte {
	var out bytes.Buffer
	if r, err := czlib.NewReader(bytes.NewReader(buffer)); err != nil {
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

func (scon *ServerConnection) SendPacketsRaw(pks [][]byte) {
	_, _ = scon.Connection.Write(scon.EncodeBatch(pks))
}

func (scon *ServerConnection) SendPacket(pk dataPacket.DataPacket) {
	buff := bytes.NewBuffer(nil)
	buff.WriteByte(pk.ID())
	pk.Marshal(dataPacket.NewWriter(buff, 2))
	_, _ = scon.Connection.Write(scon.EncodeBatch([][]byte{buff.Bytes()}))
}

func (scon *ServerConnection) EncodeBatch(pks [][]byte) []byte {
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
