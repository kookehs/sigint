package core

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"math"
	"regexp"
)

const (
	CharacterAllianceFormat  = `7300[a-z0-9]{2}([a-z0-9]{4,10}?)2d78`
	CharacterIDFormat        = `69([a-z0-9]{8})0173`
	CharacterGuildFormat     = `7300[a-z0-9]{2}([a-z0-9]{6,60}?)0978`
	CharacterNameFormat      = `7300[a-z0-9]{2}([a-z0-9]{6,32}?)02(?:62|79)`
	CharacterPositionFormat  = `66([a-z0-9]{8})([a-z0-9]{8})0e79`
	CharacterStructureFormat = `(69.*?fc6b0017)`
)

var (
	CharacterAllianceRegExp  = regexp.MustCompile(CharacterAllianceFormat)
	CharacterIDRegExp        = regexp.MustCompile(CharacterIDFormat)
	CharacterGuildRegExp     = regexp.MustCompile(CharacterGuildFormat)
	CharacterNameRegExp      = regexp.MustCompile(CharacterNameFormat)
	CharacterPositionRegExp  = regexp.MustCompile(CharacterPositionFormat)
	CharacterStructureRegExp = regexp.MustCompile(CharacterStructureFormat)
)

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Character struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Guild    string `json:"guild"`
	Alliance string `json:"alliance"`
	Position Point  `json:"point"`
}

func ParseAlliance(payload []byte, out *string) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterAllianceRegExp.FindSubmatch(payload)

	if match == nil {
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

func ParseCharacter(payload []byte, out *Character) *Character {
	if payload == nil || out == nil {
		return nil
	}

	if !ParseID(payload, &out.ID) {
		return nil
	}

	if !ParseName(payload, &out.Name) {
		return nil
	}

	ParseGuild(payload, &out.Guild)
	ParseAlliance(payload, &out.Alliance)
	ParsePosition(payload, &out.Position)
	return out
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

func ParseID(payload []byte, out *uint32) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterIDRegExp.FindSubmatch(payload)

	if match == nil {
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

func ParseGuild(payload []byte, out *string) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterGuildRegExp.FindSubmatch(payload)

	if match == nil {
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

func ParseName(payload []byte, out *string) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterNameRegExp.FindSubmatch(payload)

	if match == nil {
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

func ParsePosition(payload []byte, out *Point) bool {
	if payload == nil || out == nil {
		return false
	}

	match := CharacterPositionRegExp.FindSubmatch(payload)

	if match == nil {
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
