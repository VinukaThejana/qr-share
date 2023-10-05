// Package utils contains utility functions
package utils

import (
	"net"
)

// GetOutBoundIP is a function that is used to get the IP address of the device that can
// connect to the internet
func GetOutBoundIP() (ip net.IP, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP, nil
}
