// Package steamapi provides an interface to the
// Steam Web API methods.
package steamapi

import (
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-retry"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// BaseSteamAPIURLProduction is the steam url used to do requests in prod
const BaseSteamAPIURLProduction = "https://api.steampowered.com"

// BaseSteamAPIURL is the url used to do requests, defaulted to prod
var BaseSteamAPIURL = BaseSteamAPIURLProduction

// A SteamMethod represents a Steam Web API method.
type SteamMethod string

// NewSteamMethod creates a new SteamMethod.
func NewSteamMethod(interf, method string, version int) SteamMethod {
	m := fmt.Sprintf("%v/%v/%v/v%v/", BaseSteamAPIURL, interf, method, strconv.Itoa(version))
	return SteamMethod(m)
}

// Request makes a request to the Steam Web API with the given
// url values and stores the result in v.
//
// Returns an error if the return status code was not 200.
func (s SteamMethod) Request(data url.Values, v interface{}) error {
	url := string(s)
	if data != nil {
		url += "?" + data.Encode()
	}

	apiRetry := retry.WithCappedDuration(time.Minute, retry.NewExponential(2*time.Second))

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for resp.StatusCode == http.StatusTooManyRequests {
		if resp.Header.Get("Retry-After") != "" {
			rTime, err := time.Parse(time.RFC1123, resp.Header.Get("Retry-After"))
			if err == nil {
				time.Sleep(time.Until(rTime))
				goto httpRequestRetry
			}

			rSecs, err := strconv.Atoi(resp.Header.Get("Retry-After"))
			if err == nil {
				time.Sleep(time.Duration(rSecs) * time.Second)
				goto httpRequestRetry
			}

			return fmt.Errorf("retry-after header contains invalid value %s", resp.Header.Get("Retry-After"))
		}

		if sleep, ok := apiRetry.Next(); !ok {
			time.Sleep(sleep)
		} else {
			panic("retry should never stop")
		}

	httpRequestRetry:
		resp, err = http.Get(url)
		if err != nil {
			return err
		}
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("steamapi %s Status code %d", s, resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	return d.Decode(&v)
}
