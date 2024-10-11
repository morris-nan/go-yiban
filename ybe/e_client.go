package ybe

import "errors"

var (
	ErrAppIDEmpty       = errors.New("app id is empty")
	ErrAppSecretEmpty   = errors.New("app secret is empty")
	ErrRedirectURIEmpty = errors.New("redirect uri is empty")
)
