package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"jimotoapi/util"
	"jimotoapi/vo"
	"os"
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
