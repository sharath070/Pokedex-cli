package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	val, ok := c.cache.Get(url)
	if ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationRespose := RespShallowLocations{}
	err = json.Unmarshal(data, &locationRespose)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)
	return locationRespose, nil
}
