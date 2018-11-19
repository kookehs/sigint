package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	MaxPacketSize = 1600
	Port          = 5056
)

func main() {
	networkInterface := ""
	handle, err := pcap.OpenLive(networkInterface, MaxPacketSize, true, pcap.BlockForever)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer handle.Close()
	characters := make(map[uint32]Character)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			continue
		}

		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, ok := udpLayer.(*layers.UDP)

			if !ok {
				continue
			}

			if udp.SrcPort == Port {
				payload := make([]byte, hex.EncodedLen(len(udp.Payload)))
				hex.Encode(payload, udp.Payload)

				events := EventCodeRegExp.FindAllSubmatch(payload, -1)

				if events == nil {
					continue
				}

				for i := 0; i < len(events); i++ {
					if len(events[i]) <= 1 {
						continue
					}

					event := make([]byte, hex.DecodedLen(len(events[i][1])))
					_, err = hex.Decode(event, events[i][1])

					if err != nil {
						continue
					}

					eventCode := binary.BigEndian.Uint16(event)

					switch eventCode {
					case 23:
						ParseCharacters(payload, characters)
						fmt.Println(characters)
					}
				}
			}

			continue
		}
	}
}
