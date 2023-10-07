package uiScene

import (
	"context"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/query"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/manageService"
	"github.com/go-omnibus/proof"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateReportDataResultMulti(ctx context.Context, resultData *rao.UIEngineResultDataMsg) ([]*rao.UIEngineResultDataMsg, error) {
	// 查询数据
	filter := bson.D{{"operator_id", resultData.OperatorID}, {"report_id", resultData.TopicID}}

	collectionName := consts.CollectUISendSceneOperator
	if !resultData.IsReport {
		collectionName = consts.CollectUISendSceneOperatorDebug
	}
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(collectionName)
	ret := &mao.UISendSceneOperator{}
	if err := collection.FindOne(ctx, filter).Decode(&ret); err != nil {
		return nil, err
	}

	// 修改 Mongo 数据 查询接口详情数据
	ret.Status = resultData.Status
	if ret.Status == "success" {
		ret.RunStatus = consts.UISceneRunStatusSuccess
	} else {
		ret.RunStatus = consts.UISceneRunStatusFail
	}
	ret.ExecTime = resultData.ExecTime
	ret.RunEndTimes = time.Now().Unix()
	ret.Msg = resultData.Msg
	ret.Screenshot = resultData.Screenshot
	ret.End = resultData.End
	ret.IsMulti = resultData.IsMulti

	asserts, err := bson.Marshal(mao.AssertResults{
		Asserts: resultData.Assertions,
	})
	if err != nil {
		log.Logger.Info("TransSaveReqToUISceneOperatorMao.asserts bson marshal err", proof.WithError(err))
	}
	ret.AssertResults = asserts

	withdraws, err := bson.Marshal(mao.WithdrawResults{
		Withdraws: resultData.Withdraws,
	})
	if err != nil {
		log.Logger.Info("TransSaveReqToUISceneOperatorMao.withdraws bson marshal err", proof.WithError(err))
	}
	ret.WithdrawResults = withdraws

	// 正常步骤转成图片
	if resultData.IsReport && !resultData.IsMulti && len(resultData.Screenshot) > 0 {
		filePath, err := manageService.FileUploadBase64Req(ctx, &rao.FileUploadBase64Req{
			FileString: resultData.Screenshot,
			PathDir:    resultData.TopicID,
			FileName:   resultData.OperatorID,
		})
		if err != nil {
			log.Logger.Error("manageService.FileUploadBase64Req err", err)
		}
		// 返回文件的本地路径
		ret.ScreenshotUrl = filePath
		if len(filePath) > 0 {
			ret.Screenshot = ""
		}
	}

	// 循环步骤处理并且转成图片
	uiEngineResultDataMsgs := make([]*rao.UIEngineResultDataMsg, 0)
	if resultData.IsMulti {
		var multiResults *mao.MultiResults
		if err := bson.Unmarshal(ret.MultiResults, &multiResults); err != nil {
			log.Logger.Errorf("ret.MultiResults bson unmarshal err %w", err)
		}

		if resultData.IsReport && len(resultData.Screenshot) > 0 {
			var multiResultsCount = len(multiResults.MultiResults)
			filePath, err := manageService.FileUploadBase64Req(ctx, &rao.FileUploadBase64Req{
				FileString: resultData.Screenshot,
				PathDir:    resultData.TopicID,
				FileName:   fmt.Sprintf("%s_%d", resultData.OperatorID, multiResultsCount),
			})
			if err != nil {
				log.Logger.Error("manageService.FileUploadBase64Req err", err)
			}
			// 返回文件的本地路径
			ret.ScreenshotUrl = filePath
			if len(filePath) > 0 {
				ret.Screenshot = ""
			}
		}

		uiEngineResultDataMsg := &rao.UIEngineResultDataMsg{
			TopicID:    resultData.TopicID,
			UserID:     resultData.UserID,
			OperatorID: resultData.OperatorID,
			SceneID:    resultData.SceneID,
			Sort:       resultData.Sort,
			ExecTime:   resultData.ExecTime,
			Status:     resultData.Status,
			Msg:        resultData.Msg,
			End:        resultData.End,
			Assertions: resultData.Assertions,
		}
		if len(ret.ScreenshotUrl) > 0 {
			uiEngineResultDataMsg.Screenshot = ""
			uiEngineResultDataMsg.ScreenshotUrl = ret.ScreenshotUrl
		} else {
			uiEngineResultDataMsg.Screenshot = resultData.Screenshot
		}
		if ret.Status == "success" {
			uiEngineResultDataMsg.RunStatus = consts.UISceneRunStatusSuccess
		} else {
			uiEngineResultDataMsg.RunStatus = consts.UISceneRunStatusFail
		}

		multiResults.MultiResults = append(multiResults.MultiResults, uiEngineResultDataMsg)
		for _, result := range multiResults.MultiResults {
			uiEngineResultDataMsgs = append(uiEngineResultDataMsgs, result)
		}

		maoMultiResults, err := bson.Marshal(mao.MultiResults{
			MultiResults: multiResults.MultiResults,
		})
		if err != nil {
			log.Logger.Info("mao.MultiResults bson marshal err", proof.WithError(err))
		}
		ret.MultiResults = maoMultiResults
	} else {
		maoMultiResults, _ := bson.Marshal(mao.MultiResults{
			MultiResults: nil,
		})
		ret.MultiResults = maoMultiResults
	}

	if err = query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		if _, err = collection.UpdateOne(ctx, filter, bson.M{"$set": ret}); err != nil {
			log.Logger.Error("ConsumerUIEngineResult-- Mongo UpdateOne，err:", ret.SceneID, "err:", err)
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return uiEngineResultDataMsgs, nil
}

func UpdateSimpleParent(ctx context.Context, resultData *rao.UIEngineResultDataMsg, parentOperatorID string) error {
	filter := bson.D{{"operator_id", parentOperatorID}, {"report_id", resultData.TopicID}}
	collectionName := consts.CollectUISendSceneOperator
	if !resultData.IsReport {
		collectionName = consts.CollectUISendSceneOperatorDebug
	}
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(collectionName)

	ret := &mao.UISendSceneOperator{}
	if err := collection.FindOne(ctx, filter).Decode(&ret); err != nil {
		return err
	}

	ret.Status = "success"
	ret.RunStatus = consts.UISceneRunStatusSuccess

	if err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		if _, err := collection.UpdateOne(ctx, filter, bson.M{"$set": ret}); err != nil {
			log.Logger.Error("UpdateSimpleParent-- Mongo UpdateOne，err:", ret.SceneID, "err:", err)
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
