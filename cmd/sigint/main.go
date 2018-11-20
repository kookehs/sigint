package main

import (
	"encoding/binary"
	"encoding/hex"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/kookehs/sigint/core"
	"github.com/kookehs/sigint/event"
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
	characters := make(map[uint32]core.Character)
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

				events := event.CodeRegExp.FindAllSubmatch(payload, -1)

				if events == nil {
					continue
				}

				for i := 0; i < len(events); i++ {
					if len(events[i]) <= 1 {
						continue
					}

					code := make([]byte, hex.DecodedLen(len(events[i][1])))
					_, err = hex.Decode(code, events[i][1])

					if err != nil {
						continue
					}

					eventCode := binary.BigEndian.Uint16(code)

					switch eventCode {
					case event.Leave:
						event.ParseLeave(payload, characters)
					case event.NewMob:
					case event.CastSpell:
					case event.NewCharacter:
						core.ParseCharacters(payload, characters)
					case event.NewSimpleHarvestableObjectList:
					case event.PlayerCounts:
					}
				}
			}

			continue
		}
	}
}