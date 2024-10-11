package ybc

import "github.com/morris-nan/go-yiban/ybe"

func (c *client) Verify(verifyRequest string) (bool, error) {
	verifyInfo, err := c.Decode(verifyRequest)
	if err != nil {
		return false, err
	}
	oauth := verifyInfo["visit_oauth"]
	switch oauth.(type) {
	case bool:
		return false, nil
	case map[string]interface{}:
		return true, nil
	default:
		return false, ybe.ErrVerifyOauthInfo
	}
}
