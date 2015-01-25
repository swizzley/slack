package slack

import (
	"errors"
	"net/url"
)

type oAuthResponseFull struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	SlackResponse
}

func GetOAuthToken(clientId, clientSecret, code, redirectUri string, debug bool) (accessToken string, scope string, err error) {
	values := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectUri},
	}
	response := &oAuthResponseFull{}
	err = parseResponse("oauth.access", values, response, debug)
	if err != nil {
		return "", "", err
	}
	if !response.Ok {
		return "", "", errors.New(response.Error)
	}
	return response.AccessToken, response.Scope, nil
}
