package utils

import "net"

func ParseIP4(ips string) string {
	ipp := net.ParseIP(ips)
	ret := ips
	if ipp != nil {
		if ipp.To4().String() == ret {
			return ret
		}
	}
	ipss, _ := net.LookupIP(ips)
	for _, ip := range ipss {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}
