package handler

import (
	"encoding/json"
	errs "errs"
	"fmt"
	"github.com/asim/go-micro/v3/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"jimotoapi/util"
	"jimotoapi/vo"
	"net/http"
	"os"
	"protos/account"
	"strings"
)

func (c *Client) UploadImage(context *gin.Context) {
	file, header, err := context.Request.FormFile("image")
	if err != nil {
		c.HandleRequestError(context, err)
		return
	}
	filename := header.Filename
	fn := strings.Split(filename, ".")
	suffix := ""
	if len(fn) > 1 {
		suffix = "." + fn[len(fn)-1]
	}
	ud := uuid.New().String()
	ufn := ud + suffix

	out, err := os.Create(fmt.Sprintf("%v/%v", c.config.ImageServer.LocalPath, ufn))
	if err != nil {
		c.HandleRequestError(context, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.HandleRequestError(context, err)
		return
	}

	rsp, err := json.Marshal(&vo.UploadResp{
		Url:      util.CompleteImageUrl(ufn, c.config),
		Filename: ufn,
	})
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.HandleSuccess(context, rsp)
}

func (c *Client) ParseAuthData(context *gin.Context) (*vo.AuthData, bool) {
	token := context.GetHeader(KEY_ACCESS_TOKEN)
	if token == "" {
		c.logger.Error(c.context, "Get access_token from header failed")
		context.JSON(http.StatusUnauthorized, gin.H{
			"code": errs.ERR_AUTH_FAILED,
			"msg":  errs.GetMsg(errs.ERR_AUTH_FAILED),
			"data": nil,
		})
		context.Abort()
		return nil, false
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
		return nil, false
	}
	return &vo.AuthData{
		UserId: resp.UserId,
		Email:  resp.Email,
	}, true
}
