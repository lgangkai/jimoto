package handler

import (
	errs "errs"
	"github.com/asim/go-micro/v3/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Client) HandleRpcError(context *gin.Context, err error) {
	c.logger.Error(c.context, "Call rpc server failed, error: ", err)
	code := errors.Parse(err.Error()).Code
	msg := errors.Parse(err.Error()).Detail
	context.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
	context.Abort()
}

func (c *Client) HandleJsonError(context *gin.Context, err error) {
	c.logger.Error(c.context, "Marshal data to json failed, err: ", err.Error())
	context.JSON(http.StatusInternalServerError, gin.H{
		"code": errs.ERR_PARAMETER_PARSE_FAILED,
		"msg":  errs.GetMsg(errs.ERR_PARAMETER_PARSE_FAILED),
		"data": nil,
	})
	context.Abort()
}

func (c *Client) HandleRequestError(context *gin.Context, err error) {
	c.logger.Error(c.context, "Request read error, err: ", err.Error())
	context.JSON(http.StatusBadRequest, gin.H{
		"code": errs.ERR_BAD_REQUEST,
		"msg":  errs.GetMsg(errs.ERR_BAD_REQUEST),
		"data": nil,
	})
	context.Abort()
}

func (c *Client) HandleSuccess(context *gin.Context, msg []byte) {
	if msg == nil {
		msg = []byte{}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": errs.SUCCESS,
		"msg":  errs.GetMsg(errs.SUCCESS),
		"data": string(msg),
	})
}
