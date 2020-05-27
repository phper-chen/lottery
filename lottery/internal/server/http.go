package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	"app/lottery/internal/model"
	"app/lottery/utils"
)

func prizeList(c *gin.Context) {
	code := http.StatusOK
	data, err := svc.GetPrizeList(c)
	if err != nil {
		code = http.StatusNotFound
	}
	// TODO 错误响应规范化处理
	c.Render(code, render.JSON{Data: data})
}

func drawPrize(c *gin.Context) {
	code := http.StatusOK
	drawRet, err := svc.DrawPrize(c, 13720600497)
	if err != nil {
		code = http.StatusBadRequest
	}
	c.Render(code, render.JSON{Data: drawRet})

}

func drawRecords(c *gin.Context) {
	data := svc.ExportDrawRecords(c)
	c.Writer.Header().Set("Content-Type", "application/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", data.Title))
	c.Render(http.StatusOK, utils.CSV{Title: data.Title, Content: data.Content})
}

func participate(c *gin.Context) {
	code := http.StatusOK
	req := new(model.UserInfoReq)
	if err := c.ShouldBind(&req); err != nil {
		c.Render(http.StatusBadRequest, render.JSON{Data: err})
		return
	}
	data, err := svc.Participate(c, req)
	if err != nil {
		c.Render(http.StatusBadRequest, render.JSON{Data: err})
		return
	}
	// TODO 错误响应规范化处理
	c.Render(code, render.JSON{Data: data})
}

func allUsers(c *gin.Context) {
	code := http.StatusOK
	data, err := svc.GetAllUsers(c)
	if err != nil {
		c.Render(http.StatusBadRequest, render.JSON{Data: err})
		return
	}
	// TODO 错误响应规范化处理
	c.Render(code, render.JSON{Data: data})
}