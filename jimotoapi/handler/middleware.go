package handler

import (
	errs "errs"
	"github.com/asim/go-micro/v3/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"protos/account"
)

const (
	KEY_REQUEST_ID     = "request_id"
	KEY_ACCESS_TOKEN   = "Authorization"
	KEY_USER_ID        = "user_id"
	KEY_EMAIL          = "email"
	COOKIE_EXPIRE_TIME = 3600 * 24 * 30
)

func (c *Client) GenRequestId(ctx *gin.Context) {
	u, err := uuid.NewUUID()
	if err != nil {
		c.logger.Warning(c.context, "Gen request_id failed, err: ", err.Error())
		return
	}
	ctx.Set("request_id", u.String())
}

func GetRequestId(ctx *gin.Context) string {
	requestId, ok := ctx.Get("request_id")
	if !ok {
		requestId = ""
	}
	return requestId.(string)
}

func (c *Client) Cors(ctx *gin.Context) {
	method := ctx.Request.Method
	origin := ctx.Request.Header.Get("Origin")
	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id, Accept, Origin, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	ctx.Next()
}

func (c *Client) Authenticate(context *gin.Context) {
	c.logger.Info(c.context, "Authenticate for api request.")
	token := context.GetHeader(KEY_ACCESS_TOKEN)
	if token == "" {
		c.logger.Error(c.context, "Get access_token from header failed")
		context.JSON(http.StatusUnauthorized, gin.H{
			"code": errs.ERR_AUTH_FAILED,
			"msg":  errs.GetMsg(errs.ERR_AUTH_FAILED),
			"data": nil,
		})
		context.Abort()
		return
	}
	c.logger.Info(c.context, "access_token: ", token)

	req := &account.AuthRequest{
		Token:     token,
		RequestId: GetRequestId(context),
	}
	resp, err := c.accountClient.Authenticate(context, req)
	if err != nil {
		c.logger.Error(c.context, "Authenticate failed, err: ", err.Error())
		code := errors.Parse(err.Error()).Code
		msg := errors.Parse(err.Error()).Detail
		context.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  msg,
			"data": nil,
		})
		context.Abort()
		return
	}

	c.logger.Info(c.context, "Authenticate succeed, userId: ", resp.GetUserId())
	context.Set(KEY_USER_ID, resp.GetUserId())
	context.Set(KEY_EMAIL, resp.GetEmail())
	context.Next()
}
