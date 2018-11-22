package core

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func ParseByte(b []byte, out *byte) bool {
	if b == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(data, b)

	if err != nil {
		log.Println(err)
		return false
	}

	*out = data[0]
	return true
}

func ParseFloat32(b []byte, out *float32) bool {
	if b == nil || out == nil {
		return false
	}

	position := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(position, b)

	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println(position)

	v := binary.BigEndian.Uint32(position)
	*out = math.Float32frombits(v)
	return true
}

func ParseString(b []byte, out *string) bool {
	if b == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(data, b)

	if err != nil {
		log.Println(err)
		return false
	}

	*out = string(data)
	return true
}

func ParseUint16(b []byte, out *uint16) bool {
	if b == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(data, b)

	if err != nil {
		log.Println(err)
		return false
	}

	*out = binary.BigEndian.Uint16(data)
	return true
}

func ParseUint32(b []byte, out *uint32) bool {
	if b == nil || out == nil {
		return false
	}

	data := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(data, b)

	if err != nil {
		log.Println(err)
		return false
	}

	*out = binary.BigEndian.Uint32(data)
	return true
}
