package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jimotoapi/vo"
	"protos/account"
)

func (c *Client) Login(context *gin.Context) {
	login := &vo.AccountReq{}
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
	loginRespJson, err := json.Marshal(&vo.LoginResp{AccessToken: resp.GetToken()})
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}

	c.HandleSuccess(context, loginRespJson)
}

func (c *Client) Register(context *gin.Context) {
	reg := &vo.AccountReq{}
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
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	marshal, err := json.Marshal(&vo.User{UserId: authData.UserId})
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.HandleSuccess(context, marshal)
}
