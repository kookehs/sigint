package core

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"math"
)

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func ParseByte(match [][]byte, out *byte) bool {
	if match == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(match[1])))
	_, err := hex.Decode(data, match[1])

	if err != nil {
		log.Println(err)
		return false
	}

	*out = data[0]
	return true
}

func ParsePoint(match [][]byte, out *Point) bool {
	if match == nil || out == nil {
		return false
	}

	position := make([]byte, hex.DecodedLen(len(match[1])))
	_, err := hex.Decode(position, match[1])

	if err != nil {
		log.Println(err)
		return false
	}

	x := binary.BigEndian.Uint32(position)
	(*out).X = math.Float32frombits(x)

	position = make([]byte, hex.DecodedLen(len(match[2])))
	_, err = hex.Decode(position, match[2])

	if err != nil {
		log.Println(err)
		return false
	}

	y := binary.BigEndian.Uint32(position)
	(*out).Y = math.Float32frombits(y)
	return true
}

func ParseString(match [][]byte, out *string) bool {
	if match == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(match[1])))
	_, err := hex.Decode(data, match[1])

	if err != nil {
		log.Println(err)
		return false
	}

	*out = string(data)
	return true
}

func ParseUint32(match [][]byte, out *uint32) bool {
	if match == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(match[1])))
	_, err := hex.Decode(data, match[1])

	if err != nil {
		log.Println(err)
		return false
	}

	*out = binary.BigEndian.Uint32(data)
	return true
}
