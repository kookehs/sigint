package core

import (
	"fmt"
	"regexp"
)

type Action struct {
	Position Point  `json:"position"`
	Self     uint16 `json:"self"`
	Target   uint16 `json:"target"`
}

const (
	CastSpellIDFormat        = `6b([a-f0-9]{4})(?:016b([a-f0-9]{4}))?`
	CastSpellStructureFormat = `6b.*?fc680011`
	CastSpellPositionFormat  = `66([a-f0-9]{8})([a-f0-9]{8})`
)

var (
	CastSpellIDRegExp        = regexp.MustCompile(CastSpellIDFormat)
	CastSpellStructureRegExp = regexp.MustCompile(CastSpellStructureFormat)
	CastSpellPositionRegExp  = regexp.MustCompile(CastSpellPositionFormat)
)

func ParseCastSpell(payload []byte, out *Action) *Action {
	if payload == nil || out == nil {
		return nil
	}

	if !ParseCastSpellID(payload, &out.Self, &out.Target) {
		return nil
	}

	ParseCastSpellPosition(payload, &out.Position)
	return out
}

func ParseCastSpellID(payload []byte, self *uint16, target *uint16) bool {
	if payload == nil || self == nil || target == nil {
		return false
	}

	match := CastSpellIDRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	ret := ParseUint16(match[1], self)

	if len(match) > 1 {
		ret = ret && ParseUint16(match[2], target)
	}

	return ret
}

func ParseCastSpellPosition(payload []byte, out *Point) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CastSpellPositionRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	fmt.Println(match)
	return ParseFloat32(match[1], &out.X) && ParseFloat32(match[2], &out.Y)
}
