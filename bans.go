package steamapi

import (
	"context"
	"golang.org/x/time/rate"
	"net/url"
	"strings"
)

type playerBansJSON struct {
	Players []PlayerBan
}

// PlayerBan contains all ban status for community, VAC and economy
type PlayerBan struct {
	SteamID          uint64 `json:"SteamId,string"`
	CommunityBanned  bool
	VACBanned        bool
	EconomyBan       string
	NumberOfVACBans  uint
	DaysSinceLastBan uint
	NumberOfGameBans uint
}

// GetPlayerBans takes a list of steamIDs and returns PlayerBan slice
func GetPlayerBans(steamIDs []uint64, apiKey string, rl *rate.Limiter) ([]PlayerBan, error) {
	var allResp []PlayerBan
	var getPlayerSummaries = NewSteamMethod("ISteamUser", "GetPlayerBans", 1)

	// split into batches of 100 steamids, since endpoint is limited to 100
	strIds := steamIDs2SplitArray(steamIDs, 100)

	for _, strId := range strIds {
		vals := url.Values{}
		vals.Add("key", apiKey)
		vals.Add("steamids", strings.Join(strId, ","))

		if err := rl.Wait(context.Background()); err != nil {
			return nil, err
		}
		var resp playerBansJSON
		err := getPlayerSummaries.Request(vals, &resp)
		if err != nil {
			return nil, err
		}
		allResp = append(allResp, resp.Players...)
	}

	return allResp, nil
}
