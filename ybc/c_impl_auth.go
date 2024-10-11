package ybc

import (
	"github.com/morris-nan/go-yiban/internel/utils"
	"github.com/morris-nan/go-yiban/ybconst"
	"net/url"
)

func (c *client) AuthURL() string {
	params := url.Values{}
	params.Set("client_id", c.appID)
	params.Set("redirect_uri", c.redirectURI)
	params.Set("state", c.state)
	return utils.Format(ybconst.ApiOauthAuthorize, params)
}
