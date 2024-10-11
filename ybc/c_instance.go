package ybc

import "github.com/morris-nan/go-yiban/ybe"

func New(options ...Options) (Client, error) {
	c := client{
		appID:       "",
		appSecret:   "",
		redirectURI: "",
		state:       "",
	}
	for _, option := range options {
		option(&c)
	}
	switch {
	case c.appID == "":
		return nil, ybe.ErrAppIDEmpty
	case c.appSecret == "":
		return nil, ybe.ErrAppSecretEmpty
	case c.redirectURI == "":
		return nil, ybe.ErrRedirectURIEmpty
	default:
		return &c, nil
	}
}
