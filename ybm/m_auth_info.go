package ybm

type AuthInfo struct {
	VerifyRequest string `json:"verify_request" form:"verify_request"`
	YbUID         string `json:"yb_uid" form:"yb_uid"`
}
