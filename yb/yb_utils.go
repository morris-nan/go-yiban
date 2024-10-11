package yb

import (
	"github.com/morris-nan/go-yiban/internel/utils"
	"github.com/morris-nan/go-yiban/ybm"
	"net/url"
)

func Format(endpoint string, params url.Values) string {
	return utils.Format(endpoint, params)
}

func ParseAuthInfo(info map[string]interface{}) ybm.VerifyInfo[ybm.OauthInfo] {
	return utils.ParseAuthInfo(info)
}
