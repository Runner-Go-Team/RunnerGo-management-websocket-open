package variable

import (
	"encoding/json"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

func SaveGlobalParam(ctx *gin.Context, req *rao.SaveGlobalParamReq) error {
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectGlobalParam)
	var dataDetail string
	if req.ParamType == 1 { // cookie
		detailTemp, err := json.Marshal(req.Cookies)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else if req.ParamType == 2 { // header
		detailTemp, err := json.Marshal(req.Headers)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else if req.ParamType == 3 { // variables
		detailTemp, err := json.Marshal(req.Variables)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else if req.ParamType == 4 { // asserts
		detailTemp, err := json.Marshal(req.Asserts)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else {
		return fmt.Errorf("参数类型错误")
	}

	var globalParamData mao.GlobalParamData
	filter := bson.D{{"team_id", req.TeamID}, {"param_type", req.ParamType}}
	err := collection.FindOne(ctx, filter).Decode(&globalParamData)
	if err != nil { // 没查到，则新增
		insertData := &mao.GlobalParamData{
			TeamID:     req.TeamID,
			ParamType:  req.ParamType,
			DataDetail: dataDetail,
		}
		_, err = collection.InsertOne(ctx, insertData)
	} else { // 查到了，则更新
		updateData := &mao.GlobalParamData{
			TeamID:     req.TeamID,
			ParamType:  req.ParamType,
			DataDetail: dataDetail,
		}
		_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": updateData})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetGlobalParam(ctx *gin.Context, req *rao.GetGlobalParamReq) (*rao.GetGlobalParamResp, error) {
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectGlobalParam)
	cur, err := collection.Find(ctx, bson.D{{"team_id", req.TeamID}})
	if err != nil {
		return nil, fmt.Errorf("全局参数为空")
	}
	var globalParamDataArr []*mao.GlobalParamData
	if err := cur.All(ctx, &globalParamDataArr); err != nil {
		return nil, fmt.Errorf("全局参数数据获取失败")
	}

	cookieParam := make([]rao.CookieParam, 0, 100)
	headerParam := make([]rao.HeaderParam, 0, 100)
	variableParam := make([]rao.VariableParam, 0, 100)
	assertParam := make([]rao.AssertParam, 0, 100)
	for _, globalParamInfo := range globalParamDataArr {
		if globalParamInfo.ParamType == 1 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &cookieParam)
			if err != nil {
				return nil, err
			}
		}
		if globalParamInfo.ParamType == 2 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &headerParam)
			if err != nil {
				return nil, err
			}
		}
		if globalParamInfo.ParamType == 3 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &variableParam)
			if err != nil {
				return nil, err
			}
		}
		if globalParamInfo.ParamType == 4 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &assertParam)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &rao.GetGlobalParamResp{
		Cookies:   cookieParam,
		Headers:   headerParam,
		Variables: variableParam,
		Asserts:   assertParam,
	}
	return res, err
}

func SaveSceneParam(ctx *gin.Context, req *rao.SaveSceneParamReq) error {
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectSceneParam)
	var dataDetail string
	if req.ParamType == 1 { // cookie
		detailTemp, err := json.Marshal(req.Cookies)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else if req.ParamType == 2 { // header
		detailTemp, err := json.Marshal(req.Headers)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else if req.ParamType == 3 { // variables
		detailTemp, err := json.Marshal(req.Variables)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else if req.ParamType == 4 { // asserts
		detailTemp, err := json.Marshal(req.Asserts)
		if err != nil {
			return err
		}
		dataDetail = string(detailTemp)
	} else {
		return fmt.Errorf("参数类型错误")
	}

	var globalParamData mao.SceneParamData
	filter := bson.D{{"team_id", req.TeamID}, {"scene_id", req.SceneID}, {"param_type", req.ParamType}}
	err := collection.FindOne(ctx, filter).Decode(&globalParamData)
	if err != nil { // 没查到，则新增
		insertData := &mao.SceneParamData{
			TeamID:     req.TeamID,
			SceneID:    req.SceneID,
			ParamType:  req.ParamType,
			DataDetail: dataDetail,
		}
		_, err = collection.InsertOne(ctx, insertData)
	} else { // 查到了，则更新
		updateData := &mao.SceneParamData{
			TeamID:     req.TeamID,
			SceneID:    req.SceneID,
			ParamType:  req.ParamType,
			DataDetail: dataDetail,
		}
		_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": updateData})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetSceneParam(ctx *gin.Context, req *rao.GetSceneParamReq) (*rao.GetSceneParamResp, error) {
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectSceneParam)
	cur, err := collection.Find(ctx, bson.D{{"team_id", req.TeamID}, {"scene_id", req.SceneID}})
	if err != nil {
		return nil, fmt.Errorf("全局参数为空")
	}
	var globalParamDataArr []*mao.SceneParamData
	if err := cur.All(ctx, &globalParamDataArr); err != nil {
		return nil, fmt.Errorf("全局参数数据获取失败")
	}

	cookieParam := make([]rao.CookieParam, 0, 100)
	headerParam := make([]rao.HeaderParam, 0, 100)
	variableParam := make([]rao.VariableParam, 0, 100)
	assertParam := make([]rao.AssertParam, 0, 100)
	for _, globalParamInfo := range globalParamDataArr {
		if globalParamInfo.ParamType == 1 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &cookieParam)
			if err != nil {
				return nil, err
			}
		}
		if globalParamInfo.ParamType == 2 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &headerParam)
			if err != nil {
				return nil, err
			}
		}
		if globalParamInfo.ParamType == 3 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &variableParam)
			if err != nil {
				return nil, err
			}
		}
		if globalParamInfo.ParamType == 4 {
			err = json.Unmarshal([]byte(globalParamInfo.DataDetail), &assertParam)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &rao.GetSceneParamResp{
		Cookies:   cookieParam,
		Headers:   headerParam,
		Variables: variableParam,
		Asserts:   assertParam,
	}
	return res, err
}
