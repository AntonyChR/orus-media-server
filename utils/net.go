package utils

import (
	"errors"
	"net"
)

func GetLocalIP() (string, error) {
	var ip string
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for index, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if ipnet.IP.String()[0:3] == "192" {
					ip = ipnet.IP.String()
					break
				}
				if index == len(addresses)-1 {
					ip = ipnet.IP.String()
				}
			}
		}
	}

	if ip == "" {
		return "", errors.New("error getting local ip address")
	}

	return ip, nil
}
