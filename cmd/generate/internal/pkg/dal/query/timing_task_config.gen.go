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

	"RunnerGo-management/cmd/generate/internal/pkg/dal/model"
)

func newTimingTaskConfig(db *gorm.DB, opts ...gen.DOOption) timingTaskConfig {
	_timingTaskConfig := timingTaskConfig{}

	_timingTaskConfig.timingTaskConfigDo.UseDB(db, opts...)
	_timingTaskConfig.timingTaskConfigDo.UseModel(&model.TimingTaskConfig{})

	tableName := _timingTaskConfig.timingTaskConfigDo.TableName()
	_timingTaskConfig.ALL = field.NewAsterisk(tableName)
	_timingTaskConfig.ID = field.NewInt32(tableName, "id")
	_timingTaskConfig.PlanID = field.NewInt32(tableName, "plan_id")
	_timingTaskConfig.SenceID = field.NewInt32(tableName, "sence_id")
	_timingTaskConfig.TeamID = field.NewInt32(tableName, "team_id")
	_timingTaskConfig.Frequency = field.NewBool(tableName, "frequency")
	_timingTaskConfig.TaskExecTime = field.NewInt64(tableName, "task_exec_time")
	_timingTaskConfig.TaskCloseTime = field.NewInt64(tableName, "task_close_time")
	_timingTaskConfig.Status = field.NewInt32(tableName, "status")
	_timingTaskConfig.CreatedAt = field.NewTime(tableName, "created_at")
	_timingTaskConfig.UpdatedAt = field.NewTime(tableName, "updated_at")
	_timingTaskConfig.DeletedAt = field.NewField(tableName, "deleted_at")

	_timingTaskConfig.fillFieldMap()

	return _timingTaskConfig
}

type timingTaskConfig struct {
	timingTaskConfigDo timingTaskConfigDo

	ALL           field.Asterisk
	ID            field.Int32 // 表id
	PlanID        field.Int32 // 计划id
	SenceID       field.Int32 // 场景id
	TeamID        field.Int32 // 团队id
	Frequency     field.Bool  // 任务执行频次
	TaskExecTime  field.Int64 // 任务执行时间
	TaskCloseTime field.Int64 // 任务结束时间
	Status        field.Int32 // 任务状态：0-未执行，1-执行中，2-已过期，3-已删除
	CreatedAt     field.Time  // 创建时间
	UpdatedAt     field.Time  // 更新时间
	DeletedAt     field.Field // 删除时间

	fieldMap map[string]field.Expr
}

func (t timingTaskConfig) Table(newTableName string) *timingTaskConfig {
	t.timingTaskConfigDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timingTaskConfig) As(alias string) *timingTaskConfig {
	t.timingTaskConfigDo.DO = *(t.timingTaskConfigDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timingTaskConfig) updateTableName(table string) *timingTaskConfig {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.PlanID = field.NewInt32(table, "plan_id")
	t.SenceID = field.NewInt32(table, "sence_id")
	t.TeamID = field.NewInt32(table, "team_id")
	t.Frequency = field.NewBool(table, "frequency")
	t.TaskExecTime = field.NewInt64(table, "task_exec_time")
	t.TaskCloseTime = field.NewInt64(table, "task_close_time")
	t.Status = field.NewInt32(table, "status")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *timingTaskConfig) WithContext(ctx context.Context) *timingTaskConfigDo {
	return t.timingTaskConfigDo.WithContext(ctx)
}

func (t timingTaskConfig) TableName() string { return t.timingTaskConfigDo.TableName() }

func (t timingTaskConfig) Alias() string { return t.timingTaskConfigDo.Alias() }

func (t *timingTaskConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timingTaskConfig) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.ID
	t.fieldMap["plan_id"] = t.PlanID
	t.fieldMap["sence_id"] = t.SenceID
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["frequency"] = t.Frequency
	t.fieldMap["task_exec_time"] = t.TaskExecTime
	t.fieldMap["task_close_time"] = t.TaskCloseTime
	t.fieldMap["status"] = t.Status
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t timingTaskConfig) clone(db *gorm.DB) timingTaskConfig {
	t.timingTaskConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timingTaskConfig) replaceDB(db *gorm.DB) timingTaskConfig {
	t.timingTaskConfigDo.ReplaceDB(db)
	return t
}

type timingTaskConfigDo struct{ gen.DO }

func (t timingTaskConfigDo) Debug() *timingTaskConfigDo {
	return t.withDO(t.DO.Debug())
}

func (t timingTaskConfigDo) WithContext(ctx context.Context) *timingTaskConfigDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timingTaskConfigDo) ReadDB() *timingTaskConfigDo {
	return t.Clauses(dbresolver.Read)
}

func (t timingTaskConfigDo) WriteDB() *timingTaskConfigDo {
	return t.Clauses(dbresolver.Write)
}

func (t timingTaskConfigDo) Session(config *gorm.Session) *timingTaskConfigDo {
	return t.withDO(t.DO.Session(config))
}

func (t timingTaskConfigDo) Clauses(conds ...clause.Expression) *timingTaskConfigDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timingTaskConfigDo) Returning(value interface{}, columns ...string) *timingTaskConfigDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timingTaskConfigDo) Not(conds ...gen.Condition) *timingTaskConfigDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timingTaskConfigDo) Or(conds ...gen.Condition) *timingTaskConfigDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timingTaskConfigDo) Select(conds ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timingTaskConfigDo) Where(conds ...gen.Condition) *timingTaskConfigDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timingTaskConfigDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *timingTaskConfigDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t timingTaskConfigDo) Order(conds ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timingTaskConfigDo) Distinct(cols ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timingTaskConfigDo) Omit(cols ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timingTaskConfigDo) Join(table schema.Tabler, on ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timingTaskConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timingTaskConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timingTaskConfigDo) Group(cols ...field.Expr) *timingTaskConfigDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timingTaskConfigDo) Having(conds ...gen.Condition) *timingTaskConfigDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timingTaskConfigDo) Limit(limit int) *timingTaskConfigDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timingTaskConfigDo) Offset(offset int) *timingTaskConfigDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timingTaskConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timingTaskConfigDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timingTaskConfigDo) Unscoped() *timingTaskConfigDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timingTaskConfigDo) Create(values ...*model.TimingTaskConfig) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timingTaskConfigDo) CreateInBatches(values []*model.TimingTaskConfig, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timingTaskConfigDo) Save(values ...*model.TimingTaskConfig) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timingTaskConfigDo) First() (*model.TimingTaskConfig, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimingTaskConfig), nil
	}
}

func (t timingTaskConfigDo) Take() (*model.TimingTaskConfig, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimingTaskConfig), nil
	}
}

func (t timingTaskConfigDo) Last() (*model.TimingTaskConfig, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimingTaskConfig), nil
	}
}

func (t timingTaskConfigDo) Find() ([]*model.TimingTaskConfig, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimingTaskConfig), err
}

func (t timingTaskConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimingTaskConfig, err error) {
	buf := make([]*model.TimingTaskConfig, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timingTaskConfigDo) FindInBatches(result *[]*model.TimingTaskConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timingTaskConfigDo) Attrs(attrs ...field.AssignExpr) *timingTaskConfigDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timingTaskConfigDo) Assign(attrs ...field.AssignExpr) *timingTaskConfigDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timingTaskConfigDo) Joins(fields ...field.RelationField) *timingTaskConfigDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timingTaskConfigDo) Preload(fields ...field.RelationField) *timingTaskConfigDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timingTaskConfigDo) FirstOrInit() (*model.TimingTaskConfig, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimingTaskConfig), nil
	}
}

func (t timingTaskConfigDo) FirstOrCreate() (*model.TimingTaskConfig, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimingTaskConfig), nil
	}
}

func (t timingTaskConfigDo) FindByPage(offset int, limit int) (result []*model.TimingTaskConfig, count int64, err error) {
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

func (t timingTaskConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timingTaskConfigDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timingTaskConfigDo) Delete(models ...*model.TimingTaskConfig) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timingTaskConfigDo) withDO(do gen.Dao) *timingTaskConfigDo {
	t.DO = *do.(*gen.DO)
	return t
}
