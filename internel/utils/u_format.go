package utils

import (
	"github.com/morris-nan/go-yiban/ybconst"
	"net/url"
)

func Format(endpoint string, params url.Values) string {
	return ybconst.BaseURL + endpoint + "?" + params.Encode()
}
