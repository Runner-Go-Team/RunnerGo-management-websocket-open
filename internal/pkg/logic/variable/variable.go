package variable

import (
	"RunnerGo-management/internal/pkg/dal/mao"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"RunnerGo-management/internal/pkg/biz/consts"
	"RunnerGo-management/internal/pkg/dal"
	"RunnerGo-management/internal/pkg/dal/model"
	"RunnerGo-management/internal/pkg/dal/query"
	"RunnerGo-management/internal/pkg/dal/rao"
	"RunnerGo-management/internal/pkg/packer"
)

func SaveVariable(ctx context.Context, req *rao.SaveVariableReq) error {
	tx := query.Use(dal.DB()).Variable

	_, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.VarID)).Assign(
		tx.TeamID.Value(req.TeamID),
		tx.Var.Value(req.Var),
		tx.Val.Value(req.Val),
		tx.Status.Value(req.Status),
		tx.Description.Value(req.Description),
	).FirstOrCreate()

	return err
}

func DeleteVariable(ctx context.Context, teamID string, varID int64) error {
	tx := query.Use(dal.DB()).Variable

	_, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.ID.Eq(varID)).Delete()
	return err
}

func ListGlobalVariables(ctx context.Context, teamID string, limit, offset int) ([]*rao.Variable, int64, error) {
	tx := query.Use(dal.DB()).Variable

	v, cnt, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.Type.Eq(consts.VariableTypeGlobal)).FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return packer.TransModelVariablesToRaoVariables(v), cnt, nil
}

func SyncGlobalVariables(ctx context.Context, teamID string, variables []*rao.Variable) error {
	vs := packer.TransRaoVariablesToModelVariables(teamID, variables)

	return query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		if _, err := tx.Variable.WithContext(ctx).Where(tx.Variable.TeamID.Eq(teamID), tx.Variable.Type.Eq(consts.VariableTypeGlobal)).Delete(); err != nil {
			return err
		}

		return tx.Variable.WithContext(ctx).CreateInBatches(vs, 10)
	})
}

func ListSceneVariables(ctx context.Context, teamID string, sceneID string, limit, offset int) ([]*rao.Variable, int64, error) {
	tx := dal.GetQuery().Variable

	v, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.SceneID.Eq(sceneID), tx.Type.Eq(consts.VariableTypeScene)).Limit(limit).Offset(offset).Find()

	//v, cnt, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.SceneID.Eq(sceneID), tx.Type.Eq(consts.VariableTypeScene)).FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	cnt, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.SceneID.Eq(sceneID), tx.Type.Eq(consts.VariableTypeScene)).Count()
	if err != nil {
		return nil, 0, err
	}

	return packer.TransModelVariablesToRaoVariables(v), cnt, nil
}

func SyncSceneVariables(ctx context.Context, teamID string, sceneID string, variables []*rao.Variable) error {
	vs := packer.TransSceneRaoVariablesToModelVariables(teamID, sceneID, variables)

	return query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		if _, err := tx.Variable.WithContext(ctx).Where(tx.Variable.TeamID.Eq(teamID), tx.Variable.Type.Eq(consts.VariableTypeScene)).Unscoped().Delete(); err != nil {
			return err
		}

		return tx.Variable.WithContext(ctx).CreateInBatches(vs, 10)
	})
}

func ImportSceneVariables(ctx context.Context, req *rao.ImportVariablesReq, userID string) error {

	tx := dal.GetQuery().VariableImport
	return tx.WithContext(ctx).Create(&model.VariableImport{
		TeamID:     req.TeamID,
		SceneID:    req.SceneID,
		Name:       req.Name,
		URL:        req.URL,
		UploaderID: userID,
	})
}

func DeleteImportSceneVariables(ctx context.Context, req *rao.DeleteImportSceneVariablesReq) error {
	tx := dal.GetQuery().VariableImport
	_, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(req.TeamID), tx.SceneID.Eq(req.SceneID), tx.Name.Eq(req.Name)).Delete()
	return err
}

func ListImportSceneVariables(ctx context.Context, teamID string, sceneID string) ([]*rao.Import, error) {
	tx := dal.GetQuery().VariableImport
	vi, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.SceneID.Eq(sceneID)).Limit(5).Find()
	if err != nil {
		return nil, err
	}

	return packer.TransImportVariablesToRaoImportVariables(vi), nil
}

func UpdateImportSceneVariables(ctx *gin.Context, req *rao.UpdateImportSceneVariablesReq) error {
	tx := dal.GetQuery().VariableImport
	updateData := make(map[string]interface{}, 1)
	updateData["status"] = req.Status
	_, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.ID)).Updates(updateData)
	if err != nil {
		return err
	}
	return nil
}

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
