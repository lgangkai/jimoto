package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"protos/account"
)

func (c *Client) GetProfile(context *gin.Context) {
	uId := c.getAuthedData(context, KEY_USER_ID)
	if uId == nil {
		return
	}
	r := &account.GetProfileRequest{
		UserId:    uId.(uint64),
		RequestId: GetRequestId(context),
	}
	resp, err := c.accountClient.GetProfile(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	pr := resp.GetProfile()
	pr.Avatar = c.CompleteImageUrl(pr.Avatar)
	prj, err := json.Marshal(pr)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle get profile success, profile: ", string(prj))
	c.HandleSuccess(context, prj)
}

func (c *Client) CreateProfile(context *gin.Context) {
	uId := c.getAuthedData(context, KEY_USER_ID)
	p := &CreateProfileRequest{}
	if err := context.ShouldBind(p); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	r := &account.CreateProfileRequest{
		Profile: &account.Profile{
			UserId:   uId.(uint64),
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
	//uId := c.getAuthedData(context, KEY_USER_ID)
}

func (c *Client) DeleteProfile(context *gin.Context) {
	//uId := c.getAuthedData(context, KEY_USER_ID)
}
