package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	c.logger.Info(c.context, "Authenticate succeed, authData: ", authData)
	context.Next()
}
