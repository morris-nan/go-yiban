package utils

import "github.com/morris-nan/go-yiban/ybm"

func ParseAuthInfo(info map[string]interface{}) ybm.VerifyInfo[ybm.OauthInfo] {
	var verifyInfo ybm.VerifyInfo[ybm.OauthInfo]
	if visitTime, ok := info["visit_time"].(float64); ok {
		verifyInfo.VisitTime = int64(visitTime)
	}
	if visitUser, ok := info["visit_user"].(map[string]interface{}); ok {
		verifyInfo.VisitUser.UserID = visitUser["userid"].(string)
		verifyInfo.VisitUser.UserName = visitUser["username"].(string)
		verifyInfo.VisitUser.UserNick = visitUser["usernick"].(string)
		verifyInfo.VisitUser.UserSex = visitUser["usersex"].(string)
	}
	if visitOauth, ok := info["visit_oauth"].(map[string]interface{}); ok {
		verifyInfo.VisitOauth.AccessToken = visitOauth["access_token"].(string)
		verifyInfo.VisitOauth.TokenExpires = visitOauth["token_expires"].(string)
	}
	return verifyInfo
}

func ParseNotAuthInfo(info map[string]interface{}) ybm.VerifyInfo[bool] {
	var verifyInfo ybm.VerifyInfo[bool]
	if visitTime, ok := info["visit_time"].(float64); ok {
		verifyInfo.VisitTime = int64(visitTime)
	}
	if visitUser, ok := info["visit_user"].(map[string]interface{}); ok {
		verifyInfo.VisitUser.UserID = visitUser["userid"].(string)
	}
	verifyInfo.VisitOauth = false
	return verifyInfo
}
