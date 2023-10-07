package manageService

import (
	"context"
	"encoding/json"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/go-resty/resty/v2"
)

const (
	FileUploadBase64ReqUri = "/management/api/v1/file/upload/base64"
)

func FileUploadBase64Req(ctx context.Context, req *rao.FileUploadBase64Req) (string, error) {
	log.Logger.Info("权限接口--获取我的角色信息--参数", req)

	bodyByte, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(bodyByte).
		Post(conf.Conf.Clients.Manager.Domain + FileUploadBase64ReqUri)
	if err != nil {
		return "", err
	}

	resp := rao.SendApiResp{}
	err = json.Unmarshal(response.Body(), &resp)
	if err != nil {
		log.Logger.Info("权限接口--添加团队成员--返回值解析失败：err：", err)
		return "", err
	}

	respDateMarshal, _ := json.Marshal(resp.Data)
	respData := rao.FileUploadBase64Resp{}
	err = json.Unmarshal(respDateMarshal, &respData)
	return respData.Path, nil
}
