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

func newTimedTaskConf(db *gorm.DB, opts ...gen.DOOption) timedTaskConf {
	_timedTaskConf := timedTaskConf{}

	_timedTaskConf.timedTaskConfDo.UseDB(db, opts...)
	_timedTaskConf.timedTaskConfDo.UseModel(&model.TimedTaskConf{})

	tableName := _timedTaskConf.timedTaskConfDo.TableName()
	_timedTaskConf.ALL = field.NewAsterisk(tableName)
	_timedTaskConf.ID = field.NewInt32(tableName, "id")
	_timedTaskConf.PlanID = field.NewInt64(tableName, "plan_id")
	_timedTaskConf.SenceID = field.NewInt64(tableName, "sence_id")
	_timedTaskConf.TeamID = field.NewInt64(tableName, "team_id")
	_timedTaskConf.UserID = field.NewInt64(tableName, "user_id")
	_timedTaskConf.Frequency = field.NewInt32(tableName, "frequency")
	_timedTaskConf.TaskExecTime = field.NewInt64(tableName, "task_exec_time")
	_timedTaskConf.TaskCloseTime = field.NewInt64(tableName, "task_close_time")
	_timedTaskConf.TaskType = field.NewInt32(tableName, "task_type")
	_timedTaskConf.TaskMode = field.NewInt32(tableName, "task_mode")
	_timedTaskConf.ModeConf = field.NewString(tableName, "mode_conf")
	_timedTaskConf.Status = field.NewInt32(tableName, "status")
	_timedTaskConf.CreatedAt = field.NewTime(tableName, "created_at")
	_timedTaskConf.UpdatedAt = field.NewTime(tableName, "updated_at")
	_timedTaskConf.DeletedAt = field.NewField(tableName, "deleted_at")

	_timedTaskConf.fillFieldMap()

	return _timedTaskConf
}

type timedTaskConf struct {
	timedTaskConfDo timedTaskConfDo

	ALL           field.Asterisk
	ID            field.Int32  // 表id
	PlanID        field.Int64  // 计划id
	SenceID       field.Int64  // 场景id
	TeamID        field.Int64  // 团队id
	UserID        field.Int64  // 用户ID
	Frequency     field.Int32  // 任务执行频次
	TaskExecTime  field.Int64  // 任务执行时间
	TaskCloseTime field.Int64  // 任务结束时间
	TaskType      field.Int32  // 任务类型：1-普通任务，2-定时任务
	TaskMode      field.Int32  // 压测模式：1-并发模式，2-阶梯模式，3-错误率模式，4-响应时间模式，5-每秒请求数模式，6 -每秒事务数模式
	ModeConf      field.String // 压测详细配置
	Status        field.Int32  // 任务状态：0-未启用，1-运行中，2-已过期
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (t timedTaskConf) Table(newTableName string) *timedTaskConf {
	t.timedTaskConfDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timedTaskConf) As(alias string) *timedTaskConf {
	t.timedTaskConfDo.DO = *(t.timedTaskConfDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timedTaskConf) updateTableName(table string) *timedTaskConf {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.PlanID = field.NewInt64(table, "plan_id")
	t.SenceID = field.NewInt64(table, "sence_id")
	t.TeamID = field.NewInt64(table, "team_id")
	t.UserID = field.NewInt64(table, "user_id")
	t.Frequency = field.NewInt32(table, "frequency")
	t.TaskExecTime = field.NewInt64(table, "task_exec_time")
	t.TaskCloseTime = field.NewInt64(table, "task_close_time")
	t.TaskType = field.NewInt32(table, "task_type")
	t.TaskMode = field.NewInt32(table, "task_mode")
	t.ModeConf = field.NewString(table, "mode_conf")
	t.Status = field.NewInt32(table, "status")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *timedTaskConf) WithContext(ctx context.Context) *timedTaskConfDo {
	return t.timedTaskConfDo.WithContext(ctx)
}

func (t timedTaskConf) TableName() string { return t.timedTaskConfDo.TableName() }

func (t timedTaskConf) Alias() string { return t.timedTaskConfDo.Alias() }

func (t *timedTaskConf) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timedTaskConf) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 15)
	t.fieldMap["id"] = t.ID
	t.fieldMap["plan_id"] = t.PlanID
	t.fieldMap["sence_id"] = t.SenceID
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["frequency"] = t.Frequency
	t.fieldMap["task_exec_time"] = t.TaskExecTime
	t.fieldMap["task_close_time"] = t.TaskCloseTime
	t.fieldMap["task_type"] = t.TaskType
	t.fieldMap["task_mode"] = t.TaskMode
	t.fieldMap["mode_conf"] = t.ModeConf
	t.fieldMap["status"] = t.Status
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t timedTaskConf) clone(db *gorm.DB) timedTaskConf {
	t.timedTaskConfDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timedTaskConf) replaceDB(db *gorm.DB) timedTaskConf {
	t.timedTaskConfDo.ReplaceDB(db)
	return t
}

type timedTaskConfDo struct{ gen.DO }

func (t timedTaskConfDo) Debug() *timedTaskConfDo {
	return t.withDO(t.DO.Debug())
}

func (t timedTaskConfDo) WithContext(ctx context.Context) *timedTaskConfDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timedTaskConfDo) ReadDB() *timedTaskConfDo {
	return t.Clauses(dbresolver.Read)
}

func (t timedTaskConfDo) WriteDB() *timedTaskConfDo {
	return t.Clauses(dbresolver.Write)
}

func (t timedTaskConfDo) Session(config *gorm.Session) *timedTaskConfDo {
	return t.withDO(t.DO.Session(config))
}

func (t timedTaskConfDo) Clauses(conds ...clause.Expression) *timedTaskConfDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timedTaskConfDo) Returning(value interface{}, columns ...string) *timedTaskConfDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timedTaskConfDo) Not(conds ...gen.Condition) *timedTaskConfDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timedTaskConfDo) Or(conds ...gen.Condition) *timedTaskConfDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timedTaskConfDo) Select(conds ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timedTaskConfDo) Where(conds ...gen.Condition) *timedTaskConfDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timedTaskConfDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *timedTaskConfDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t timedTaskConfDo) Order(conds ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timedTaskConfDo) Distinct(cols ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timedTaskConfDo) Omit(cols ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timedTaskConfDo) Join(table schema.Tabler, on ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timedTaskConfDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timedTaskConfDo) RightJoin(table schema.Tabler, on ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timedTaskConfDo) Group(cols ...field.Expr) *timedTaskConfDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timedTaskConfDo) Having(conds ...gen.Condition) *timedTaskConfDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timedTaskConfDo) Limit(limit int) *timedTaskConfDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timedTaskConfDo) Offset(offset int) *timedTaskConfDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timedTaskConfDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timedTaskConfDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timedTaskConfDo) Unscoped() *timedTaskConfDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timedTaskConfDo) Create(values ...*model.TimedTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timedTaskConfDo) CreateInBatches(values []*model.TimedTaskConf, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timedTaskConfDo) Save(values ...*model.TimedTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timedTaskConfDo) First() (*model.TimedTaskConf, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimedTaskConf), nil
	}
}

func (t timedTaskConfDo) Take() (*model.TimedTaskConf, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimedTaskConf), nil
	}
}

func (t timedTaskConfDo) Last() (*model.TimedTaskConf, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimedTaskConf), nil
	}
}

func (t timedTaskConfDo) Find() ([]*model.TimedTaskConf, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimedTaskConf), err
}

func (t timedTaskConfDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimedTaskConf, err error) {
	buf := make([]*model.TimedTaskConf, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timedTaskConfDo) FindInBatches(result *[]*model.TimedTaskConf, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timedTaskConfDo) Attrs(attrs ...field.AssignExpr) *timedTaskConfDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timedTaskConfDo) Assign(attrs ...field.AssignExpr) *timedTaskConfDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timedTaskConfDo) Joins(fields ...field.RelationField) *timedTaskConfDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timedTaskConfDo) Preload(fields ...field.RelationField) *timedTaskConfDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timedTaskConfDo) FirstOrInit() (*model.TimedTaskConf, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimedTaskConf), nil
	}
}

func (t timedTaskConfDo) FirstOrCreate() (*model.TimedTaskConf, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimedTaskConf), nil
	}
}

func (t timedTaskConfDo) FindByPage(offset int, limit int) (result []*model.TimedTaskConf, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t timedTaskConfDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timedTaskConfDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timedTaskConfDo) Delete(models ...*model.TimedTaskConf) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timedTaskConfDo) withDO(do gen.Dao) *timedTaskConfDo {
	t.DO = *do.(*gen.DO)
	return t
}
