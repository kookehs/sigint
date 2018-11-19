package main

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

func ParseCharacter(payload []byte, character *Character) *Character {
	matches := CharacterIDRegExp.FindSubmatch(payload)

	if len(matches) > 1 {
		id := make([]byte, hex.DecodedLen(len(matches[1])))
		_, err := hex.Decode(id, matches[1])

		if err != nil {
			log.Println(err)
		} else {
			character.ID = binary.BigEndian.Uint32(id)
		}
	}

	matches = CharacterNameRegExp.FindSubmatch(payload)

	if len(matches) > 1 {
		name := make([]byte, hex.DecodedLen(len(matches[1])))
		_, err := hex.Decode(name, matches[1])

		if err != nil {
			log.Println(err)
		} else {
			character.Name = string(name)
		}
	}

	matches = CharacterGuildRegExp.FindSubmatch(payload)

	if len(matches) > 1 {
		guild := make([]byte, hex.DecodedLen(len(matches[1])))
		_, err := hex.Decode(guild, matches[1])

		if err != nil {
			log.Println(err)
		} else {
			character.Guild = string(guild)
		}
	}

	matches = CharacterAllianceRegExp.FindSubmatch(payload)

	if len(matches) > 1 {
		alliance := make([]byte, hex.DecodedLen(len(matches[1])))
		_, err := hex.Decode(alliance, matches[1])

		if err != nil {
			log.Println(err)
		} else {
			character.Alliance = string(alliance)
		}
	}

	matches = CharacterPositionRegExp.FindSubmatch(payload)

	if len(matches) > 2 {
		position := make([]byte, hex.DecodedLen(len(matches[1])))
		_, err := hex.Decode(position, matches[1])

		if err != nil {
			log.Println(err)
		} else {
			x := binary.BigEndian.Uint32(position)
			character.Position.X = math.Float32frombits(x)
		}

		position = make([]byte, hex.DecodedLen(len(matches[2])))
		_, err = hex.Decode(position, matches[2])

		if err != nil {
			log.Println(err)
		} else {
			y := binary.BigEndian.Uint32(position)
			character.Position.Y = math.Float32frombits(y)
		}
	}

	return character
}

func ParseCharacters(payload []byte, table map[uint32]Character) map[uint32]Character {
	characters := CharacterStructureRegExp.FindAllSubmatch(payload, -1)

	if characters == nil {
		return nil
	}

	for i := 0; i < len(characters); i++ {
		if len(characters[i]) <= 1 {
			continue
		}

		character := Character{}
		ParseCharacter(characters[i][1], &character)
		table[character.ID] = character
	}

	return table
}
