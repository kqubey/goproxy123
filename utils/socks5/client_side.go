package socks5

import (
	"errors"
	"io"
)

var (
	// ErrBadReply is the error when read reply
	ErrBadReply = errors.New("bad Reply")
	// ErrUserPassAuth is the error when got invalid username or password
	ErrUserPassAuth = errors.New("invalid Username or Password for Auth")
	// ErrVersion is version error
	ErrVersion = errors.New("invalid Version")
	// ErrUserPassVersion is username/password auth version error
	ErrUserPassVersion = errors.New("invalid Version of Username Password Auth")
	ErrBadRequest      = errors.New("bad Request")
)

// NewNegotiationRequest return negotiation request packet can be writed into server
func NewNegotiationRequest(methods []byte) *NegotiationRequest {
	return &NegotiationRequest{
		Ver:      Ver,
		NMethods: byte(len(methods)),
		Methods:  methods,
	}
}

// WriteTo write negotiation request packet into server
func (r *NegotiationRequest) WriteTo(w io.Writer) (int64, error) {
	var n int
	i, err := w.Write([]byte{r.Ver})
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write([]byte{r.NMethods})
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write(r.Methods)
	n = n + i
	if err != nil {
		return int64(n), err
	}

	return int64(n), nil
}

// NewNegotiationReplyFrom read negotiation reply packet from server
func NewNegotiationReplyFrom(r io.Reader) (*NegotiationReply, error) {
	bb := make([]byte, 2)
	if _, err := io.ReadFull(r, bb); err != nil {
		return nil, err
	}
	if bb[0] != Ver {
		return nil, ErrVersion
	}

	return &NegotiationReply{
		Ver:    bb[0],
		Method: bb[1],
	}, nil
}

// NewUserPassNegotiationRequest return user password negotiation request packet can be writed into server
func NewUserPassNegotiationRequest(username []byte, password []byte) *UserPassNegotiationRequest {
	return &UserPassNegotiationRequest{
		Ver:    UserPassVer,
		Ulen:   byte(len(username)),
		Uname:  username,
		Plen:   byte(len(password)),
		Passwd: password,
	}
}

// WriteTo write user password negotiation request packet into server
func (r *UserPassNegotiationRequest) WriteTo(w io.Writer) (int64, error) {
	var n int
	i, err := w.Write([]byte{r.Ver, r.Ulen})
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write(r.Uname)
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write([]byte{r.Plen})
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write(r.Passwd)
	n = n + i
	if err != nil {
		return int64(n), err
	}
	return int64(n), nil
}

// NewUserPassNegotiationReplyFrom read user password negotiation reply packet from server
func NewUserPassNegotiationReplyFrom(r io.Reader) (*UserPassNegotiationReply, error) {
	bb := make([]byte, 2)
	if _, err := io.ReadFull(r, bb); err != nil {
		return nil, err
	}
	if bb[0] != UserPassVer {
		return nil, ErrUserPassVersion
	}

	return &UserPassNegotiationReply{
		Ver:    bb[0],
		Status: bb[1],
	}, nil
}

// NewRequest return request packet can be writed into server, dstaddr should not have domain length
func NewRequest(cmd byte, atyp byte, dstaddr []byte, dstport []byte) *Request {
	if atyp == ATYPDomain {
		dstaddr = append([]byte{byte(len(dstaddr))}, dstaddr...)
	}
	return &Request{
		Ver:     Ver,
		Cmd:     cmd,
		Rsv:     0x00,
		Atyp:    atyp,
		DstAddr: dstaddr,
		DstPort: dstport,
	}
}

// WriteTo write request packet into server
func (r *Request) WriteTo(w io.Writer) (int64, error) {
	var n int
	i, err := w.Write([]byte{r.Ver, r.Cmd, r.Rsv, r.Atyp})
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write(r.DstAddr)
	n = n + i
	if err != nil {
		return int64(n), err
	}
	i, err = w.Write(r.DstPort)
	n = n + i
	if err != nil {
		return int64(n), err
	}

	return int64(n), nil
}

// NewReplyFrom read reply packet from server
func NewReplyFrom(r io.Reader) (*Reply, error) {
	bb := make([]byte, 4)
	if _, err := io.ReadFull(r, bb); err != nil {
		return nil, err
	}
	if bb[0] != Ver {
		return nil, ErrVersion
	}
	var addr []byte
	if bb[3] == ATYPIPv4 {
		addr = make([]byte, 4)
		if _, err := io.ReadFull(r, addr); err != nil {
			return nil, err
		}
	} else if bb[3] == ATYPIPv6 {
		addr = make([]byte, 16)
		if _, err := io.ReadFull(r, addr); err != nil {
			return nil, err
		}
	} else if bb[3] == ATYPDomain {
		dal := make([]byte, 1)
		if _, err := io.ReadFull(r, dal); err != nil {
			return nil, err
		}
		if dal[0] == 0 {
			return nil, ErrBadReply
		}
		addr = make([]byte, int(dal[0]))
		if _, err := io.ReadFull(r, addr); err != nil {
			return nil, err
		}
		addr = append(dal, addr...)
	} else {
		return nil, ErrBadReply
	}
	port := make([]byte, 2)
	if _, err := io.ReadFull(r, port); err != nil {
		return nil, err
	}

	return &Reply{
		Ver:     bb[0],
		Rep:     bb[1],
		Rsv:     bb[2],
		Atyp:    bb[3],
		BndAddr: addr,
		BndPort: port,
	}, nil
}
