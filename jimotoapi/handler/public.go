package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
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
	rsp, err := json.Marshal(&UploadResp{Filename: ufn})
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}

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
	c.HandleSuccess(context, rsp)
}

func (c *Client) CompleteImageUrl(filename string) string {
	if c.config.ImageServer.Local {
		if filename == "" {
			return ""
		}
		return fmt.Sprintf("%v://%v/%v/%v", c.config.Server.Scheme, c.config.Server.Addr, c.config.ImageServer.LocalPath, filename)
	}
	return ""
}

func (c *Client) CompleteImageUrls(filenames []string) []string {
	rst := make([]string, len(filenames))
	for i, filename := range filenames {
		rst[i] = c.CompleteImageUrl(filename)
	}
	return rst
}
