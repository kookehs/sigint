package main

import (
	"log"

	"github.com/google/gopacket/pcap"
)

func GetNetworkDevices() []pcap.Interface {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Println(err)
		return nil
	}

	return devices
}
