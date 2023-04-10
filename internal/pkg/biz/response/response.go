package response

import (
	"RunnerGo-management/internal/pkg/biz/errno"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Em   string      `json:"em,omitempty"`
	Et   string      `json:"et,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type WbResponse struct {
	Code     int         `json:"code"`
	Em       string      `json:"em,omitempty"`
	Et       string      `json:"et,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	RouteUrl string      `json:"route_url"`
}

func display(c *gin.Context, code int, em string, et string, data interface{}) {
	respData := Response{
		Code: code,
		Em:   em,
		Et:   et,
		Data: data,
	}
	c.JSON(http.StatusOK, respData)
}

func wbDisplay(c context.Context, code int, em string, et string, data interface{}, routeUrl string) string {
	respData := WbResponse{
		Code:     code,
		Em:       em,
		Et:       et,
		Data:     data,
		RouteUrl: routeUrl,
	}
	resTemp, err := json.Marshal(respData)
	if err != nil {
		return ""
	}
	return string(resTemp)
}

// ErrorWithMsg 返回错误 附带更多信息
func ErrorWithMsg(c *gin.Context, code int, msg string) {
	if m, ok := errno.CodeMsgMap[code]; ok {
		msg = m + " " + msg
	}
	display(c, code, msg, errno.CodeAlertMap[code], struct{}{})
}

// SuccessWithData 返回成功并携带数据
func SuccessWithData(c *gin.Context, data interface{}) {
	display(c, errno.Ok, errno.CodeMsgMap[errno.Ok], errno.CodeAlertMap[errno.Ok], data)
}

func Success(c *gin.Context) {
	display(c, errno.Ok, errno.CodeMsgMap[errno.Ok], errno.CodeAlertMap[errno.Ok], struct{}{})
}

// websocket相关的响应方法

func WbErrorWithMsg(c context.Context, code int, msg string, routeUrl string) string {
	if m, ok := errno.CodeMsgMap[code]; ok {
		msg = m + " " + msg
	}
	return wbDisplay(c, code, msg, errno.CodeAlertMap[code], struct{}{}, routeUrl)
}

func WbSuccess(c context.Context, routeUrl string) string {
	return wbDisplay(c, errno.Ok, errno.CodeMsgMap[errno.Ok], errno.CodeAlertMap[errno.Ok], struct{}{}, routeUrl)
}

// WbSuccessWithData 返回成功并携带数据
func WbSuccessWithData(c context.Context, data interface{}, routeUrl string) string {
	return wbDisplay(c, errno.Ok, errno.CodeMsgMap[errno.Ok], errno.CodeAlertMap[errno.Ok], data, routeUrl)
}
