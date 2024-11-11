package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jimotoapi/util"
	"jimotoapi/vo"
	"protos/account"
	"strconv"
)

func (c *Client) GetProfile(context *gin.Context) {
	var uId uint64
	// 1. if query param exists, use query param.
	uIdQ := context.Query(KEY_USER_ID)
	if uIdQ != "" {
		uIdI, _ := strconv.Atoi(uIdQ)
		uId = uint64(uIdI)
	} else {
		// 2. else use auth data
		authData, ok := c.ParseAuthData(context)
		if !ok {
			return
		}
		uId = authData.UserId
	}
	r := &account.GetProfileRequest{
		UserId:    uId,
		RequestId: GetRequestId(context),
	}
	resp, err := c.accountClient.GetProfile(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	pr := resp.GetProfile()
	pr.Avatar = util.CompleteImageUrl(pr.Avatar, c.config)
	prj, err := json.Marshal(pr)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle get profile success, profile: ", string(prj))
	c.HandleSuccess(context, prj)
}

func (c *Client) CreateProfile(context *gin.Context) {
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	p := &vo.CreateProfileReq{}
	if err := context.ShouldBind(p); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	r := &account.CreateProfileRequest{
		Profile: &account.Profile{
			UserId:   authData.UserId,
			Username: p.Username,
			Email:    p.Email,
		},
		RequestId: GetRequestId(context),
	}
	_, err := c.accountClient.CreateProfile(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	c.HandleSuccess(context, nil)
}

func (c *Client) UpdateProfile(context *gin.Context) {
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	p := &vo.UpdateProfileReq{}
	if err := context.ShouldBind(p); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	r := &account.UpdateProfileRequest{
		Profile: &account.Profile{
			UserId:   authData.UserId,
			Username: p.Username,
			Avatar:   p.AvatarUrl,
		},
		RequestId: GetRequestId(context),
	}
	_, err := c.accountClient.UpdateProfile(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	c.HandleSuccess(context, nil)
}

func (c *Client) DeleteProfile(context *gin.Context) {
	//uId := c.getAuthedData(context, KEY_USER_ID)
}
