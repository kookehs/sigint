package core

import (
	"regexp"
)

const (
	MobIDFormat        = `69([a-z0-9]{8})fc6b006a`
	MobPositionFormat  = `66([a-z0-9]{8})([a-z0-9]{8})0879`
	MobStructureFormat = `(1ca.*?fc6b006a)`
	MobTierFormat      = `0262([a-z0-9]{2})0673`
)

var (
	MobIDRegExp        = regexp.MustCompile(MobIDFormat)
	MobPositionRegExp  = regexp.MustCompile(MobPositionFormat)
	MobStructureRegExp = regexp.MustCompile(MobStructureFormat)
	MobTierRegExp      = regexp.MustCompile(MobTierFormat)
)

type Mob struct {
	ID       uint32 `json:"id"`
	Position Point  `json:"point"`
	Tier     byte   `json:"tier"`
}

func ParseMob(payload []byte, out *Mob) *Mob {
	if payload == nil || out == nil {
		return nil
	}

	if !ParseMobID(payload, &out.ID) {
		return nil
	}

	ParseMobPosition(payload, &out.Position)
	ParseMobTier(payload, &out.Tier)
	return out
}

func ParseMobID(payload []byte, out *uint32) bool {
	if payload == nil || out == nil {
		return false
	}

	match := MobIDRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParseUint32(match, out)
}

func ParseMobPosition(payload []byte, out *Point) bool {
	if payload == nil || out == nil {
		return false
	}

	match := MobPositionRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParsePoint(match, out)
}

func ParseMobs(payload []byte, out map[uint32]Mob) map[uint32]Mob {
	if payload == nil || out == nil {
		return nil
	}

	matches := MobStructureRegExp.FindAllSubmatch(payload, -1)

	if matches == nil {
		return nil
	}

	for i := 0; i < len(matches); i++ {
		if len(matches[i]) <= 1 {
			continue
		}

		mob := Mob{}

		if ParseMob(matches[i][1], &mob) != nil {
			out[mob.ID] = mob
		}
	}

	return out
}

func ParseMobTier(payload []byte, out *byte) bool {
	if payload == nil || out == nil {
		return false
	}

	match := MobTierRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParseByte(match, out)
}
