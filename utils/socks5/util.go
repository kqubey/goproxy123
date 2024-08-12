package socks5

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
)

/*
func() {
		fbyts, _ := ioutil.ReadFile("socks5.txt")
		good := ""
		for _, pip := range strings.Split(string(fbyts), "\n") {
			log.Println("testing", pip)
			proxy, _ := socks5.NewClient(pip, "", "", 10, 10)
			proxyerr := make(chan error)
			proxyconn := make(chan net.Conn)
			go func() {
				udpConn, err := proxy.Dial("udp", utils.ParseIP4(strings.Split(config.Data.Address, ":")[0])+":"+strings.Split(config.Data.Address, ":")[1])
				if err != nil {
					proxyerr <- err
					return
				}
				proxyconn <- udpConn
			}()
			select {
			case <-proxyerr:
				break
			case raklib.ProxyConn = <-proxyconn:
				raklib.UseProxy = true
				_, err := raklib.DialTimeout(config.Data.Address, time.Second*5)
				if err == nil {
					log.Println("good", pip)
					good += pip + "\n"
				}
			case <-time.After(4 * time.Second):
				break
			}
		}



log.Println(good)
}()
os.Exit(0)
*/

func NewDatagram(atyp byte, dstaddr []byte, dstport []byte, data []byte) *Datagram {
	if atyp == ATYPDomain {
		dstaddr = append([]byte{byte(len(dstaddr))}, dstaddr...)
	}
	return &Datagram{
		Rsv:     []byte{0x00, 0x00},
		Frag:    0x00,
		Atyp:    atyp,
		DstAddr: dstaddr,
		DstPort: dstport,
		Data:    data,
	}
}

func (d *Datagram) Bytes() []byte {
	b := make([]byte, 0)
	b = append(b, d.Rsv...)
	b = append(b, d.Frag)
	b = append(b, d.Atyp)
	b = append(b, d.DstAddr...)
	b = append(b, d.DstPort...)
	b = append(b, d.Data...)
	return b
}

// ParseAddress format address x.x.x.x:xx to raw address.
// addr contains domain length
func ParseAddress(address string) (a byte, addr []byte, port []byte, err error) {
	var h, p string
	h, p, err = net.SplitHostPort(address)
	if err != nil {
		return
	}
	ip := net.ParseIP(h)
	if ip4 := ip.To4(); ip4 != nil {
		a = ATYPIPv4
		addr = ip4
	} else if ip6 := ip.To16(); ip6 != nil {
		a = ATYPIPv6
		addr = ip6
	} else {
		a = ATYPDomain
		addr = []byte{byte(len(h))}
		addr = append(addr, []byte(h)...)
	}
	i, _ := strconv.Atoi(p)
	port = make([]byte, 2)
	binary.BigEndian.PutUint16(port, uint16(i))
	return
}

func NewDatagramFromBytes(bb []byte) (*Datagram, error) {
	n := len(bb)
	minl := 4
	if n < minl {
		return nil, ErrBadRequest
	}
	var addr []byte
	if bb[3] == ATYPIPv4 {
		minl += 4
		if n < minl {
			return nil, ErrBadRequest
		}
		addr = bb[minl-4 : minl]
	} else if bb[3] == ATYPIPv6 {
		minl += 16
		if n < minl {
			return nil, ErrBadRequest
		}
		addr = bb[minl-16 : minl]
	} else if bb[3] == ATYPDomain {
		minl += 1
		if n < minl {
			return nil, ErrBadRequest
		}
		l := bb[4]
		if l == 0 {
			return nil, ErrBadRequest
		}
		minl += int(l)
		if n < minl {
			return nil, ErrBadRequest
		}
		addr = bb[minl-int(l) : minl]
		addr = append([]byte{l}, addr...)
	} else {
		return nil, ErrBadRequest
	}
	minl += 2
	if n <= minl {
		return nil, ErrBadRequest
	}
	port := bb[minl-2 : minl]
	data := bb[minl:]
	d := &Datagram{
		Rsv:     bb[0:2],
		Frag:    bb[2],
		Atyp:    bb[3],
		DstAddr: addr,
		DstPort: port,
		Data:    data,
	}

	return d, nil
}

// Address return request address like ip:xx
func (r *Request) Address() string {
	var s string
	if r.Atyp == ATYPDomain {
		s = bytes.NewBuffer(r.DstAddr[1:]).String()
	} else {
		s = net.IP(r.DstAddr).String()
	}
	p := strconv.Itoa(int(binary.BigEndian.Uint16(r.DstPort)))
	return net.JoinHostPort(s, p)
}

// Address return request address like ip:xx
func (r *Reply) Address() string {
	var s string
	if r.Atyp == ATYPDomain {
		s = bytes.NewBuffer(r.BndAddr[1:]).String()
	} else {
		s = net.IP(r.BndAddr).String()
	}
	p := strconv.Itoa(int(binary.BigEndian.Uint16(r.BndPort)))
	return net.JoinHostPort(s, p)
}

// Address return datagram address like ip:xx
func (d *Datagram) Address() string {
	var s string
	if d.Atyp == ATYPDomain {
		s = bytes.NewBuffer(d.DstAddr[1:]).String()
	} else {
		s = net.IP(d.DstAddr).String()
	}
	p := strconv.Itoa(int(binary.BigEndian.Uint16(d.DstPort)))
	return net.JoinHostPort(s, p)
}
