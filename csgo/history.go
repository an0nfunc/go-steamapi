package csgo

import (
	"github.com/an0nfunc/go-steamapi"
	"net/url"
	"strconv"
)

type shareCodeJson struct {
	Result struct {
		Code string `json:"nextcode"`
	} `json:"result"`
}

const APPID = 730

// GetNextMatchSharingCode queries for a newer ShareCode than knownCode. If there is no new sharecode, returns "n/a".
func GetNextMatchSharingCode(steamID uint64, authCode string, knownCode string, apiKey string) (shareCode string, err error) {
	getNextMatchSharingCode := steamapi.NewSteamMethod("ICSGOPlayers_"+strconv.Itoa(APPID), "GetNextMatchSharingCode", 1)

	queryParams := url.Values{}
	queryParams.Add("key", apiKey)
	queryParams.Add("steamid", strconv.FormatUint(steamID, 10))
	queryParams.Add("steamidkey", authCode)
	queryParams.Add("knowncode", knownCode)

	res := new(shareCodeJson)
	err = getNextMatchSharingCode.Request(queryParams, res)
	if err != nil {
		return
	}

	shareCode = res.Result.Code
	return
}
