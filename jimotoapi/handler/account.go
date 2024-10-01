package handler

import (
	"encoding/json"
	errs "errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"protos/account"
)

func (c *Client) Login(context *gin.Context) {
	login := &AccountRequest{}
	if err := context.ShouldBind(login); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	r := &account.LoginRequest{
		Email:     login.Email,
		Password:  login.Password,
		RequestId: GetRequestId(context),
	}
	resp, err := c.accountClient.Login(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}

	c.logger.Info(c.context, "Handle login success.")
	context.SetCookie(KEY_ACCESS_TOKEN, resp.GetToken(), COOKIE_EXPIRE_TIME, "/", "", false, false)
	loginRespJson, err := json.Marshal(&LoginResp{UserId: resp.GetUserId()})
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}

	c.HandleSuccess(context, loginRespJson)
}

func (c *Client) Register(context *gin.Context) {
	reg := &AccountRequest{}
	if err := context.ShouldBind(reg); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	r := &account.RegisterRequest{
		Email:     reg.Email,
		Password:  reg.Password,
		RequestId: GetRequestId(context),
	}
	c.logger.Info(c.context, "Call register, request: ", r)
	_, err := c.accountClient.Register(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}

	c.logger.Info(c.context, "Handle register success.")
	c.HandleSuccess(context, nil)
}

func (c *Client) GetUserId(context *gin.Context) {
	value, exists := context.Get(KEY_USER_ID)
	if !exists || value == nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code": errs.ERR_AUTH_FAILED,
			"msg":  errs.GetMsg(errs.ERR_AUTH_FAILED),
			"data": nil,
		})
		context.Abort()
		return
	}
	marshal, err := json.Marshal(&User{UserId: value.(uint64)})
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.HandleSuccess(context, marshal)
}

// Logout currently just delete the token and no need to call service.
func (c *Client) Logout(context *gin.Context) {
	context.SetCookie(KEY_ACCESS_TOKEN, "", -1, "/", "", false, true)
	c.HandleSuccess(context, nil)
}

func (c *Client) getAuthedData(context *gin.Context, key string) any {
	v, ok := context.Get(key)
	if !ok {
		c.logger.Error(c.context, "Get data failed, key: ", key)
		context.JSON(http.StatusUnauthorized, gin.H{
			"code": errs.ERR_AUTH_FAILED,
			"msg":  errs.GetMsg(errs.ERR_AUTH_FAILED),
			"data": nil,
		})
		context.Abort()
		return nil
	}
	return v
}
