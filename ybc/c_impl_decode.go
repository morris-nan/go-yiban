package ybc

import (
	"encoding/json"
	"github.com/morris-nan/go-yiban/internel/utils"
)

func (c *client) Decode(verifyRequest string) (map[string]any, error) {
	var verifyInfo map[string]interface{}
	decode, err := utils.Decode(verifyRequest, c.appSecret, c.appID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(decode), &verifyInfo); err != nil {
		return nil, err
	}
	return verifyInfo, nil
}
