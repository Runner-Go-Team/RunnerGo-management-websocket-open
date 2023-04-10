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

	"RunnerGo-management/internal/pkg/dal/model"
)

func newStressPlanTaskConf(db *gorm.DB, opts ...gen.DOOption) stressPlanTaskConf {
	_stressPlanTaskConf := stressPlanTaskConf{}

	_stressPlanTaskConf.stressPlanTaskConfDo.UseDB(db, opts...)
	_stressPlanTaskConf.stressPlanTaskConfDo.UseModel(&model.StressPlanTaskConf{})

	tableName := _stressPlanTaskConf.stressPlanTaskConfDo.TableName()
	_stressPlanTaskConf.ALL = field.NewAsterisk(tableName)
	_stressPlanTaskConf.ID = field.NewInt32(tableName, "id")
	_stressPlanTaskConf.PlanID = field.NewString(tableName, "plan_id")
	_stressPlanTaskConf.TeamID = field.NewString(tableName, "team_id")
	_stressPlanTaskConf.SceneID = field.NewString(tableName, "scene_id")
	_stressPlanTaskConf.TaskType = field.NewInt32(tableName, "task_type")
	_stressPlanTaskConf.TaskMode = field.NewInt32(tableName, "task_mode")
	_stressPlanTaskConf.ControlMode = field.NewInt32(tableName, "control_mode")
	_stressPlanTaskConf.ModeConf = field.NewString(tableName, "mode_conf")
	_stressPlanTaskConf.RunUserID = field.NewString(tableName, "run_user_id")
	_stressPlanTaskConf.CreatedAt = field.NewTime(tableName, "created_at")
	_stressPlanTaskConf.UpdatedAt = field.NewTime(tableName, "updated_at")
	_stressPlanTaskConf.DeletedAt = field.NewField(tableName, "deleted_at")

	_stressPlanTaskConf.fillFieldMap()

	return _stressPlanTaskConf
}

type stressPlanTaskConf struct {
	stressPlanTaskConfDo stressPlanTaskConfDo

	ALL         field.Asterisk
	ID          field.Int32  // 配置ID
	PlanID      field.String // 计划ID
	TeamID      field.String // 团队ID
	SceneID     field.String // 场景ID
	TaskType    field.Int32  // 任务类型：1-普通模式，2-定时任务
	TaskMode    field.Int32  // 压测模式：1-并发模式，2-阶梯模式，3-错误率模式，4-响应时间模式，5-每秒请求数模式，6-每秒事务数模式
	ControlMode field.Int32  // 控制模式：0-集中模式，1-单独模式
	ModeConf    field.String // 压测模式配置详情
	RunUserID   field.String // 运行人用户ID
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 更新时间
	DeletedAt   field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s stressPlanTaskConf) Table(newTableName string) *stressPlanTaskConf {
	s.stressPlanTaskConfDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stressPlanTaskConf) As(alias string) *stressPlanTaskConf {
	s.stressPlanTaskConfDo.DO = *(s.stressPlanTaskConfDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stressPlanTaskConf) updateTableName(table string) *stressPlanTaskConf {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.PlanID = field.NewString(table, "plan_id")
	s.TeamID = field.NewString(table, "team_id")
	s.SceneID = field.NewString(table, "scene_id")
	s.TaskType = field.NewInt32(table, "task_type")
	s.TaskMode = field.NewInt32(table, "task_mode")
	s.ControlMode = field.NewInt32(table, "control_mode")
	s.ModeConf = field.NewString(table, "mode_conf")
	s.RunUserID = field.NewString(table, "run_user_id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *stressPlanTaskConf) WithContext(ctx context.Context) *stressPlanTaskConfDo {
	return s.stressPlanTaskConfDo.WithContext(ctx)
}

func (s stressPlanTaskConf) TableName() string { return s.stressPlanTaskConfDo.TableName() }

func (s stressPlanTaskConf) Alias() string { return s.stressPlanTaskConfDo.Alias() }

func (s *stressPlanTaskConf) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stressPlanTaskConf) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["plan_id"] = s.PlanID
	s.fieldMap["team_id"] = s.TeamID
	s.fieldMap["scene_id"] = s.SceneID
	s.fieldMap["task_type"] = s.TaskType
	s.fieldMap["task_mode"] = s.TaskMode
	s.fieldMap["control_mode"] = s.ControlMode
	s.fieldMap["mode_conf"] = s.ModeConf
	s.fieldMap["run_user_id"] = s.RunUserID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s stressPlanTaskConf) clone(db *gorm.DB) stressPlanTaskConf {
	s.stressPlanTaskConfDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stressPlanTaskConf) replaceDB(db *gorm.DB) stressPlanTaskConf {
	s.stressPlanTaskConfDo.ReplaceDB(db)
	return s
}

type stressPlanTaskConfDo struct{ gen.DO }

func (s stressPlanTaskConfDo) Debug() *stressPlanTaskConfDo {
	return s.withDO(s.DO.Debug())
}

func (s stressPlanTaskConfDo) WithContext(ctx context.Context) *stressPlanTaskConfDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stressPlanTaskConfDo) ReadDB() *stressPlanTaskConfDo {
	return s.Clauses(dbresolver.Read)
}

func (s stressPlanTaskConfDo) WriteDB() *stressPlanTaskConfDo {
	return s.Clauses(dbresolver.Write)
}

func (s stressPlanTaskConfDo) Session(config *gorm.Session) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Session(config))
}

func (s stressPlanTaskConfDo) Clauses(conds ...clause.Expression) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stressPlanTaskConfDo) Returning(value interface{}, columns ...string) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stressPlanTaskConfDo) Not(conds ...gen.Condition) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stressPlanTaskConfDo) Or(conds ...gen.Condition) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stressPlanTaskConfDo) Select(conds ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stressPlanTaskConfDo) Where(conds ...gen.Condition) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stressPlanTaskConfDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *stressPlanTaskConfDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s stressPlanTaskConfDo) Order(conds ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stressPlanTaskConfDo) Distinct(cols ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stressPlanTaskConfDo) Omit(cols ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stressPlanTaskConfDo) Join(table schema.Tabler, on ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stressPlanTaskConfDo) LeftJoin(table schema.Tabler, on ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stressPlanTaskConfDo) RightJoin(table schema.Tabler, on ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stressPlanTaskConfDo) Group(cols ...field.Expr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stressPlanTaskConfDo) Having(conds ...gen.Condition) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stressPlanTaskConfDo) Limit(limit int) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stressPlanTaskConfDo) Offset(offset int) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stressPlanTaskConfDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stressPlanTaskConfDo) Unscoped() *stressPlanTaskConfDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stressPlanTaskConfDo) Create(values ...*model.StressPlanTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stressPlanTaskConfDo) CreateInBatches(values []*model.StressPlanTaskConf, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stressPlanTaskConfDo) Save(values ...*model.StressPlanTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stressPlanTaskConfDo) First() (*model.StressPlanTaskConf, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTaskConf), nil
	}
}

func (s stressPlanTaskConfDo) Take() (*model.StressPlanTaskConf, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTaskConf), nil
	}
}

func (s stressPlanTaskConfDo) Last() (*model.StressPlanTaskConf, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTaskConf), nil
	}
}

func (s stressPlanTaskConfDo) Find() ([]*model.StressPlanTaskConf, error) {
	result, err := s.DO.Find()
	return result.([]*model.StressPlanTaskConf), err
}

func (s stressPlanTaskConfDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StressPlanTaskConf, err error) {
	buf := make([]*model.StressPlanTaskConf, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stressPlanTaskConfDo) FindInBatches(result *[]*model.StressPlanTaskConf, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stressPlanTaskConfDo) Attrs(attrs ...field.AssignExpr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stressPlanTaskConfDo) Assign(attrs ...field.AssignExpr) *stressPlanTaskConfDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stressPlanTaskConfDo) Joins(fields ...field.RelationField) *stressPlanTaskConfDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stressPlanTaskConfDo) Preload(fields ...field.RelationField) *stressPlanTaskConfDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stressPlanTaskConfDo) FirstOrInit() (*model.StressPlanTaskConf, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTaskConf), nil
	}
}

func (s stressPlanTaskConfDo) FirstOrCreate() (*model.StressPlanTaskConf, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTaskConf), nil
	}
}

func (s stressPlanTaskConfDo) FindByPage(offset int, limit int) (result []*model.StressPlanTaskConf, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s stressPlanTaskConfDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stressPlanTaskConfDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stressPlanTaskConfDo) Delete(models ...*model.StressPlanTaskConf) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stressPlanTaskConfDo) withDO(do gen.Dao) *stressPlanTaskConfDo {
	s.DO = *do.(*gen.DO)
	return s
}
