// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
)

func newUIScene(db *gorm.DB, opts ...gen.DOOption) uIScene {
	_uIScene := uIScene{}

	_uIScene.uISceneDo.UseDB(db, opts...)
	_uIScene.uISceneDo.UseModel(&model.UIScene{})

	tableName := _uIScene.uISceneDo.TableName()
	_uIScene.ALL = field.NewAsterisk(tableName)
	_uIScene.ID = field.NewInt64(tableName, "id")
	_uIScene.SceneID = field.NewString(tableName, "scene_id")
	_uIScene.SceneType = field.NewString(tableName, "scene_type")
	_uIScene.TeamID = field.NewString(tableName, "team_id")
	_uIScene.Name = field.NewString(tableName, "name")
	_uIScene.ParentID = field.NewString(tableName, "parent_id")
	_uIScene.Sort = field.NewInt32(tableName, "sort")
	_uIScene.Status = field.NewInt32(tableName, "status")
	_uIScene.Version = field.NewInt32(tableName, "version")
	_uIScene.Source = field.NewInt32(tableName, "source")
	_uIScene.PlanID = field.NewString(tableName, "plan_id")
	_uIScene.CreatedUserID = field.NewString(tableName, "created_user_id")
	_uIScene.RecentUserID = field.NewString(tableName, "recent_user_id")
	_uIScene.Description = field.NewString(tableName, "description")
	_uIScene.UIMachineKey = field.NewString(tableName, "ui_machine_key")
	_uIScene.SourceID = field.NewString(tableName, "source_id")
	_uIScene.Browsers = field.NewString(tableName, "browsers")
	_uIScene.CreatedAt = field.NewTime(tableName, "created_at")
	_uIScene.UpdatedAt = field.NewTime(tableName, "updated_at")
	_uIScene.DeletedAt = field.NewField(tableName, "deleted_at")

	_uIScene.fillFieldMap()

	return _uIScene
}

type uIScene struct {
	uISceneDo uISceneDo

	ALL           field.Asterisk
	ID            field.Int64  // id
	SceneID       field.String // 全局唯一ID
	SceneType     field.String // 类型：文件夹，场景
	TeamID        field.String // 团队id
	Name          field.String // 名称
	ParentID      field.String // 父级ID
	Sort          field.Int32  // 排序
	Status        field.Int32  // 回收站状态：1-正常，2-回收站
	Version       field.Int32  // 产品版本号
	Source        field.Int32  // 数据来源：1-场景管理，2-计划
	PlanID        field.String // 计划ID
	CreatedUserID field.String // 创建人ID
	RecentUserID  field.String // 最近修改人ID
	Description   field.String // 备注
	UIMachineKey  field.String // 指定执行的UI自动化机器key
	SourceID      field.String // 引用来源ID
	Browsers      field.String // 浏览器信息
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (u uIScene) Table(newTableName string) *uIScene {
	u.uISceneDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uIScene) As(alias string) *uIScene {
	u.uISceneDo.DO = *(u.uISceneDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uIScene) updateTableName(table string) *uIScene {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.SceneID = field.NewString(table, "scene_id")
	u.SceneType = field.NewString(table, "scene_type")
	u.TeamID = field.NewString(table, "team_id")
	u.Name = field.NewString(table, "name")
	u.ParentID = field.NewString(table, "parent_id")
	u.Sort = field.NewInt32(table, "sort")
	u.Status = field.NewInt32(table, "status")
	u.Version = field.NewInt32(table, "version")
	u.Source = field.NewInt32(table, "source")
	u.PlanID = field.NewString(table, "plan_id")
	u.CreatedUserID = field.NewString(table, "created_user_id")
	u.RecentUserID = field.NewString(table, "recent_user_id")
	u.Description = field.NewString(table, "description")
	u.UIMachineKey = field.NewString(table, "ui_machine_key")
	u.SourceID = field.NewString(table, "source_id")
	u.Browsers = field.NewString(table, "browsers")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *uIScene) WithContext(ctx context.Context) *uISceneDo { return u.uISceneDo.WithContext(ctx) }

func (u uIScene) TableName() string { return u.uISceneDo.TableName() }

func (u uIScene) Alias() string { return u.uISceneDo.Alias() }

func (u *uIScene) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uIScene) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 20)
	u.fieldMap["id"] = u.ID
	u.fieldMap["scene_id"] = u.SceneID
	u.fieldMap["scene_type"] = u.SceneType
	u.fieldMap["team_id"] = u.TeamID
	u.fieldMap["name"] = u.Name
	u.fieldMap["parent_id"] = u.ParentID
	u.fieldMap["sort"] = u.Sort
	u.fieldMap["status"] = u.Status
	u.fieldMap["version"] = u.Version
	u.fieldMap["source"] = u.Source
	u.fieldMap["plan_id"] = u.PlanID
	u.fieldMap["created_user_id"] = u.CreatedUserID
	u.fieldMap["recent_user_id"] = u.RecentUserID
	u.fieldMap["description"] = u.Description
	u.fieldMap["ui_machine_key"] = u.UIMachineKey
	u.fieldMap["source_id"] = u.SourceID
	u.fieldMap["browsers"] = u.Browsers
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u uIScene) clone(db *gorm.DB) uIScene {
	u.uISceneDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uIScene) replaceDB(db *gorm.DB) uIScene {
	u.uISceneDo.ReplaceDB(db)
	return u
}

type uISceneDo struct{ gen.DO }

func (u uISceneDo) Debug() *uISceneDo {
	return u.withDO(u.DO.Debug())
}

func (u uISceneDo) WithContext(ctx context.Context) *uISceneDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uISceneDo) ReadDB() *uISceneDo {
	return u.Clauses(dbresolver.Read)
}

func (u uISceneDo) WriteDB() *uISceneDo {
	return u.Clauses(dbresolver.Write)
}

func (u uISceneDo) Session(config *gorm.Session) *uISceneDo {
	return u.withDO(u.DO.Session(config))
}

func (u uISceneDo) Clauses(conds ...clause.Expression) *uISceneDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uISceneDo) Returning(value interface{}, columns ...string) *uISceneDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uISceneDo) Not(conds ...gen.Condition) *uISceneDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uISceneDo) Or(conds ...gen.Condition) *uISceneDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uISceneDo) Select(conds ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uISceneDo) Where(conds ...gen.Condition) *uISceneDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uISceneDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *uISceneDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u uISceneDo) Order(conds ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uISceneDo) Distinct(cols ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uISceneDo) Omit(cols ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uISceneDo) Join(table schema.Tabler, on ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uISceneDo) LeftJoin(table schema.Tabler, on ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uISceneDo) RightJoin(table schema.Tabler, on ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uISceneDo) Group(cols ...field.Expr) *uISceneDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uISceneDo) Having(conds ...gen.Condition) *uISceneDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uISceneDo) Limit(limit int) *uISceneDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uISceneDo) Offset(offset int) *uISceneDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uISceneDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *uISceneDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uISceneDo) Unscoped() *uISceneDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uISceneDo) Create(values ...*model.UIScene) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uISceneDo) CreateInBatches(values []*model.UIScene, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uISceneDo) Save(values ...*model.UIScene) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uISceneDo) First() (*model.UIScene, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIScene), nil
	}
}

func (u uISceneDo) Take() (*model.UIScene, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIScene), nil
	}
}

func (u uISceneDo) Last() (*model.UIScene, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIScene), nil
	}
}

func (u uISceneDo) Find() ([]*model.UIScene, error) {
	result, err := u.DO.Find()
	return result.([]*model.UIScene), err
}

func (u uISceneDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UIScene, err error) {
	buf := make([]*model.UIScene, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uISceneDo) FindInBatches(result *[]*model.UIScene, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uISceneDo) Attrs(attrs ...field.AssignExpr) *uISceneDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uISceneDo) Assign(attrs ...field.AssignExpr) *uISceneDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uISceneDo) Joins(fields ...field.RelationField) *uISceneDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uISceneDo) Preload(fields ...field.RelationField) *uISceneDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uISceneDo) FirstOrInit() (*model.UIScene, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIScene), nil
	}
}

func (u uISceneDo) FirstOrCreate() (*model.UIScene, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIScene), nil
	}
}

func (u uISceneDo) FindByPage(offset int, limit int) (result []*model.UIScene, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u uISceneDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uISceneDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uISceneDo) Delete(models ...*model.UIScene) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uISceneDo) withDO(do gen.Dao) *uISceneDo {
	u.DO = *do.(*gen.DO)
	return u
}