package event

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"regexp"

	"github.com/kookehs/sigint/core"
)

const (
	LeaveFormat = `0069([a-f0-9]{8})fc6b0001`
)

var (
	LeaveRegExp = regexp.MustCompile(LeaveFormat)
)

func ParseLeave(payload []byte, out map[uint32]core.Character) *core.Character {
	match := LeaveRegExp.FindSubmatch(payload)

	if match == nil {
		return nil
	}

	data := make([]byte, hex.DecodedLen(len(match[1])))
	_, err := hex.Decode(data, match[1])

	if err != nil {
		log.Println(err)
		return nil
	}

	id := binary.BigEndian.Uint32(data)
	character, ok := out[id]

	if !ok {
		return nil
	}

	delete(out, id)
	return &character
}
