package summoner

import (
	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var buildUri = v4.BuildUriFunc(`summoner`)

func get(c riot.ApiClient, endpoint string) (summoner Summoner, err error) {
	req := c.NewRequest(buildUri(endpoint))
	req.Get(&summoner)
	return summoner, err
}

func ByName(c riot.ApiClient, name string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/by-name/"+name)
	return summoner, err
}

func ByAccountID(c riot.ApiClient, id string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/by-account/"+id)
	return summoner, err
}

func ByPUUID(c riot.ApiClient, puuid string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/by-puuid/"+puuid)
	return summoner, err
}

func ByID(c riot.ApiClient, id string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/"+id)
	return summoner, err
}
