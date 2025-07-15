package helper

import "net"

func GetClientIP(addr net.Addr) string {
	addr_str := addr.String()
	host, _, err := net.SplitHostPort(addr_str)
	if err != nil {
		return addr_str
	}
	return host
}
