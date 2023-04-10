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

func newPlan(db *gorm.DB, opts ...gen.DOOption) plan {
	_plan := plan{}

	_plan.planDo.UseDB(db, opts...)
	_plan.planDo.UseModel(&model.Plan{})

	tableName := _plan.planDo.TableName()
	_plan.ALL = field.NewAsterisk(tableName)
	_plan.ID = field.NewInt64(tableName, "id")
	_plan.TeamID = field.NewInt64(tableName, "team_id")
	_plan.Rank = field.NewInt64(tableName, "rank")
	_plan.Name = field.NewString(tableName, "name")
	_plan.TaskType = field.NewInt32(tableName, "task_type")
	_plan.Mode = field.NewInt32(tableName, "mode")
	_plan.Status = field.NewInt32(tableName, "status")
	_plan.CreateUserIdentify = field.NewString(tableName, "create_user_identify")
	_plan.RunUserIdentify = field.NewString(tableName, "run_user_identify")
	_plan.CreateUserID = field.NewInt64(tableName, "create_user_id")
	_plan.RunUserID = field.NewInt64(tableName, "run_user_id")
	_plan.Remark = field.NewString(tableName, "remark")
	_plan.CronExpr = field.NewString(tableName, "cron_expr")
	_plan.CreatedAt = field.NewTime(tableName, "created_at")
	_plan.UpdatedAt = field.NewTime(tableName, "updated_at")
	_plan.DeletedAt = field.NewField(tableName, "deleted_at")

	_plan.fillFieldMap()

	return _plan
}

type plan struct {
	planDo planDo

	ALL                field.Asterisk
	ID                 field.Int64  // 主键
	TeamID             field.Int64  // 团队ID
	Rank               field.Int64  // 团队内份数
	Name               field.String // 计划名称
	TaskType           field.Int32  // 计划类型{1:普通任务,2:定时任务}
	Mode               field.Int32  // 压测类型 1 // 并发模式，  2 // 阶梯模式，  3 // 错误率模式，  4 // 响应时间模式，  5 //每秒请求数模式，  6 //每秒事务数模式，
	Status             field.Int32  // 计划状态1:未开始,2:进行中
	CreateUserIdentify field.String
	RunUserIdentify    field.String
	CreateUserID       field.Int64  // 创建人id
	RunUserID          field.Int64  // 运行人id
	Remark             field.String // 备注
	CronExpr           field.String // 定时任务表达式
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field

	fieldMap map[string]field.Expr
}

func (p plan) Table(newTableName string) *plan {
	p.planDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p plan) As(alias string) *plan {
	p.planDo.DO = *(p.planDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *plan) updateTableName(table string) *plan {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.TeamID = field.NewInt64(table, "team_id")
	p.Rank = field.NewInt64(table, "rank")
	p.Name = field.NewString(table, "name")
	p.TaskType = field.NewInt32(table, "task_type")
	p.Mode = field.NewInt32(table, "mode")
	p.Status = field.NewInt32(table, "status")
	p.CreateUserIdentify = field.NewString(table, "create_user_identify")
	p.RunUserIdentify = field.NewString(table, "run_user_identify")
	p.CreateUserID = field.NewInt64(table, "create_user_id")
	p.RunUserID = field.NewInt64(table, "run_user_id")
	p.Remark = field.NewString(table, "remark")
	p.CronExpr = field.NewString(table, "cron_expr")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *plan) WithContext(ctx context.Context) *planDo { return p.planDo.WithContext(ctx) }

func (p plan) TableName() string { return p.planDo.TableName() }

func (p plan) Alias() string { return p.planDo.Alias() }

func (p *plan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *plan) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 16)
	p.fieldMap["id"] = p.ID
	p.fieldMap["team_id"] = p.TeamID
	p.fieldMap["rank"] = p.Rank
	p.fieldMap["name"] = p.Name
	p.fieldMap["task_type"] = p.TaskType
	p.fieldMap["mode"] = p.Mode
	p.fieldMap["status"] = p.Status
	p.fieldMap["create_user_identify"] = p.CreateUserIdentify
	p.fieldMap["run_user_identify"] = p.RunUserIdentify
	p.fieldMap["create_user_id"] = p.CreateUserID
	p.fieldMap["run_user_id"] = p.RunUserID
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["cron_expr"] = p.CronExpr
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p plan) clone(db *gorm.DB) plan {
	p.planDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p plan) replaceDB(db *gorm.DB) plan {
	p.planDo.ReplaceDB(db)
	return p
}

type planDo struct{ gen.DO }

func (p planDo) Debug() *planDo {
	return p.withDO(p.DO.Debug())
}

func (p planDo) WithContext(ctx context.Context) *planDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p planDo) ReadDB() *planDo {
	return p.Clauses(dbresolver.Read)
}

func (p planDo) WriteDB() *planDo {
	return p.Clauses(dbresolver.Write)
}

func (p planDo) Session(config *gorm.Session) *planDo {
	return p.withDO(p.DO.Session(config))
}

func (p planDo) Clauses(conds ...clause.Expression) *planDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p planDo) Returning(value interface{}, columns ...string) *planDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p planDo) Not(conds ...gen.Condition) *planDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p planDo) Or(conds ...gen.Condition) *planDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p planDo) Select(conds ...field.Expr) *planDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p planDo) Where(conds ...gen.Condition) *planDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p planDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *planDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p planDo) Order(conds ...field.Expr) *planDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p planDo) Distinct(cols ...field.Expr) *planDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p planDo) Omit(cols ...field.Expr) *planDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p planDo) Join(table schema.Tabler, on ...field.Expr) *planDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p planDo) LeftJoin(table schema.Tabler, on ...field.Expr) *planDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p planDo) RightJoin(table schema.Tabler, on ...field.Expr) *planDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p planDo) Group(cols ...field.Expr) *planDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p planDo) Having(conds ...gen.Condition) *planDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p planDo) Limit(limit int) *planDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p planDo) Offset(offset int) *planDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p planDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *planDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p planDo) Unscoped() *planDo {
	return p.withDO(p.DO.Unscoped())
}

func (p planDo) Create(values ...*model.Plan) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p planDo) CreateInBatches(values []*model.Plan, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p planDo) Save(values ...*model.Plan) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p planDo) First() (*model.Plan, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plan), nil
	}
}

func (p planDo) Take() (*model.Plan, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plan), nil
	}
}

func (p planDo) Last() (*model.Plan, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plan), nil
	}
}

func (p planDo) Find() ([]*model.Plan, error) {
	result, err := p.DO.Find()
	return result.([]*model.Plan), err
}

func (p planDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Plan, err error) {
	buf := make([]*model.Plan, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p planDo) FindInBatches(result *[]*model.Plan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p planDo) Attrs(attrs ...field.AssignExpr) *planDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p planDo) Assign(attrs ...field.AssignExpr) *planDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p planDo) Joins(fields ...field.RelationField) *planDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p planDo) Preload(fields ...field.RelationField) *planDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p planDo) FirstOrInit() (*model.Plan, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plan), nil
	}
}

func (p planDo) FirstOrCreate() (*model.Plan, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Plan), nil
	}
}

func (p planDo) FindByPage(offset int, limit int) (result []*model.Plan, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p planDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p planDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p planDo) Delete(models ...*model.Plan) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *planDo) withDO(do gen.Dao) *planDo {
	p.DO = *do.(*gen.DO)
	return p
}