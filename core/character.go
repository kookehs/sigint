package core

import (
	"regexp"
)

const (
	CharacterAllianceFormat  = `2c7300[a-z0-9]{2}([a-z0-9]{4,10}?)2d78`
	CharacterIDFormat        = `0069([a-z0-9]{8})0173`
	CharacterGuildFormat     = `087300[a-z0-9]{2}([a-z0-9]{6,60}?)0978`
	CharacterNameFormat      = `017300[a-z0-9]{2}([a-z0-9]{6,32}?)02(?:62|79)`
	CharacterPositionFormat  = `66([a-z0-9]{8})([a-z0-9]{8})0e79`
	CharacterStructureFormat = `(0069.*?fc6b0017)`
)

var (
	CharacterAllianceRegExp  = regexp.MustCompile(CharacterAllianceFormat)
	CharacterIDRegExp        = regexp.MustCompile(CharacterIDFormat)
	CharacterGuildRegExp     = regexp.MustCompile(CharacterGuildFormat)
	CharacterNameRegExp      = regexp.MustCompile(CharacterNameFormat)
	CharacterPositionRegExp  = regexp.MustCompile(CharacterPositionFormat)
	CharacterStructureRegExp = regexp.MustCompile(CharacterStructureFormat)
)

type Character struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Guild    string `json:"guild"`
	Alliance string `json:"alliance"`
	Position Point  `json:"point"`
}

func ParseCharacterAlliance(payload []byte, out *string) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterAllianceRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParseString(match, out)
}

func ParseCharacter(payload []byte, out *Character) *Character {
	if payload == nil || out == nil {
		return nil
	}

	if !ParseCharacterID(payload, &out.ID) {
		return nil
	}

	if !ParseCharacterName(payload, &out.Name) {
		return nil
	}

	ParseCharacterGuild(payload, &out.Guild)
	ParseCharacterAlliance(payload, &out.Alliance)
	ParseCharacterPosition(payload, &out.Position)
	return out
}

func ParseCharacterID(payload []byte, out *uint32) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterIDRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParseUint32(match, out)
}

func ParseCharacterGuild(payload []byte, out *string) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterGuildRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParseString(match, out)
}

func ParseCharacterName(payload []byte, out *string) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterNameRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParseString(match, out)
}

func ParseCharacterPosition(payload []byte, out *Point) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterPositionRegExp.FindSubmatch(payload)

	if match == nil {
		return false
	}

	return ParsePoint(match, out)
}

func ParseCharacters(payload []byte, out map[uint32]Character) map[uint32]Character {
	if payload == nil || out == nil {
		return nil
	}

	matches := CharacterStructureRegExp.FindAllSubmatch(payload, -1)

	if matches == nil {
		return nil
	}

	for i := 0; i < len(matches); i++ {
		if len(matches[i]) <= 1 {
			continue
		}

		character := Character{}

		if ParseCharacter(matches[i][1], &character) != nil {
			out[character.ID] = character
		}
	}

	return out
}
