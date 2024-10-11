package ybc

type Options func(c *client)

func WithAppID(appID string) Options {
	return func(c *client) {
		c.appID = appID
	}
}

func WithAppSecret(appSecret string) Options {
	return func(c *client) {
		c.appSecret = appSecret
	}
}

func WithRedirectURI(redirectURI string) Options {
	return func(c *client) {
		c.redirectURI = redirectURI
	}
}

func WithState(state string) Options {
	return func(c *client) {
		c.state = state
	}
}
