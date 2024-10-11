package ybm

type VerifyInfo[T OauthInfo | bool] struct {
	VisitTime  int64    `json:"visit_time"`
	VisitUser  UserInfo `json:"visit_user"`
	VisitOauth T        `json:"visit_oauth"`
}

type UserInfo struct {
	UserID   string `json:"userid"`
	UserName string `json:"username,omitempty"`
	UserNick string `json:"usernick,omitempty"`
	UserSex  string `json:"usersex,omitempty"`
}

type OauthInfo struct {
	AccessToken  string `json:"access_token" mapstructure:"access_token"`
	TokenExpires string `json:"token_expires" mapstructure:"token_expires"`
}
