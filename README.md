# go-yiban



一个在Gin框架中简单使用的例子

~~~go
func (y YiBan) YiBan(c *gin.Context) {

	var authInfo ybm.AuthInfo
	if err := c.ShouldBind(&authInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	client, err := yb.Client(
		ybc.WithAppID("xxxx"),
		ybc.WithAppSecret("xxxx"),
        ybc.WithRedirectURI("https://xxxx"),
		ybc.WithState("xxxx"),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	verify, err := client.Verify(authInfo.VerifyRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if verify == true {
		decode, err := client.Decode(authInfo.VerifyRequest)
		if err != nil {
			return
		}
		authInfo := yb.ParseAuthInfo(decode)
		//....
        //你的放行逻辑
		
		c.Redirect(http.StatusFound, "http://xxxx")
		return
	} else {

		c.Redirect(http.StatusFound, client.AuthURL())
		return
	}

}
~~~

