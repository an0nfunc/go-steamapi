package steamapi

import (
	"context"
	"encoding/json"
	"golang.org/x/time/rate"
	"net/url"
	"strconv"
)

// classInfoJSON are details of the specific class_id
type classInfoJSON struct {
	Result map[string]json.RawMessage `json:"result"`
}

// Info is the details about the class info
type Info struct {
	ClassID        string `json:"classid"`
	IconURL        string `json:"icon_url"`
	MarketHashName string `json:"market_hash_name"`
	Tradable       string
	Marketable     string
}

// GetAssetClassInfo returns asset details
func GetAssetClassInfo(appID, classID uint64, language, apiKey string, rl *rate.Limiter) (*Info, error) {
	var getAssetClassInfo = NewSteamMethod("ISteamEconomy", "GetAssetClassInfo", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("appid", strconv.FormatUint(appID, 10))
	vals.Add("language", language)
	vals.Add("class_count", "1")
	vals.Add("classid0", strconv.FormatUint(classID, 10))

	if err := rl.Wait(context.Background()); err != nil {
		return nil, err
	}
	var resp classInfoJSON
	err := getAssetClassInfo.Request(vals, &resp)
	if err != nil {
		return nil, err
	}

	var info Info
	for _, object := range resp.Result {
		err := json.Unmarshal(object, &info)
		if err != nil {
			continue
		}
	}

	return &info, nil
}
