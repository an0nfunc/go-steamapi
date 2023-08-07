package steamapi

import (
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidId = errors.New("Invalid Steam ID")
)

type SteamId struct {
	X uint32
	Y uint32
	Z uint32
}

func NewIdFrom32bit(i uint32) (id SteamId) {
	id.Y = i % 2
	id.Z = i / 2
	return
}

func NewIdFrom64bit(i uint64) (id SteamId) {
	i -= 0x0110000100000000
	id = NewIdFrom32bit(uint32(i))
	return
}

func NewIdFromVanityUrl(vanityUrl, apiKey string, rl *rate.Limiter) (id SteamId, err error) {
	resp, err := ResolveVanityURL(vanityUrl, apiKey, rl)
	if err != nil {
		return
	}

	id = NewIdFrom64bit(resp.SteamID)
	return
}

func NewIdFromString(s string) (id SteamId, err error) {
	validId := regexp.MustCompile("STEAM_\\d:\\d:\\d+")

	if !validId.MatchString(s) {
		err = ErrInvalidId
		return
	}

	tmp := strings.Split(s, ":")
	tmpX, _ := strconv.ParseUint(strings.Split(tmp[0], "_")[1], 10, 32)
	tmpY, _ := strconv.ParseUint(tmp[1], 10, 32)
	tmpZ, _ := strconv.ParseUint(tmp[2], 10, 32)

	id.X = uint32(tmpX)
	id.Y = uint32(tmpY)
	id.Z = uint32(tmpZ)
	return
}

func (id SteamId) String() (s string) {
	return fmt.Sprintf("STEAM_%d:%d:%d", id.X, id.Y, id.Z)
}

func (id SteamId) As32Bit() (i uint32) {
	i = id.Z*2 + id.Y
	return
}

func (id SteamId) As64Bit() (i uint64) {
	i = uint64(id.As32Bit()) + 0x0110000100000000
	return
}
