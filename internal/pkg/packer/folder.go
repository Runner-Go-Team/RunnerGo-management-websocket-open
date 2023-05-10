package packer

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

func TransSaveFolderReqToMaoFolder(folder *rao.SaveFolderReq) *mao.Folder {
	//request, err := bson.Marshal(folder.Request)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("folder.request json marshal err %w", err))
	//}
	//
	//script, err := bson.Marshal(folder.Script)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("folder.script json marshal err %w", err))
	//}

	return &mao.Folder{
		TargetID: folder.TargetID,
		//Request:  request,
		//Script:   script,
	}
}

func TransTargetToRaoFolder(t *model.Target) *rao.Folder {
	return &rao.Folder{
		TargetID:    t.TargetID,
		TeamID:      t.TeamID,
		ParentID:    t.ParentID,
		Name:        t.Name,
		Method:      t.Method,
		Sort:        t.Sort,
		TypeSort:    t.TypeSort,
		Version:     t.Version,
		Description: t.Description,
		//Request:  &r,
		//Script:   &s,
	}
}
